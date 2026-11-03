#!/usr/bin/env bash
# Copyright (c) 2021 Michael Canady
# SPDX-License-Identifier: MIT
#
# Compares the live branch rulesets with the policy file, or applies the
# policy file to them. The policy file is the source of truth for branch
# protection; never change the live rulesets by hand.
#
# Usage:
#   scripts/rulesets.sh check   Report drift; exit 1 if any ruleset differs.
#   scripts/rulesets.sh plan    Report drift and print the payload that apply would send.
#   scripts/rulesets.sh apply   Update every ruleset that differs (needs admin).
#
# Each branchProtectionRules entry maps to the live ruleset whose only
# included ref is refs/heads/<branchNamePattern>. The script sets every
# setting that the policy file declares. A live setting that the file
# doesn't declare, such as allowed_merge_methods, is kept as is.
#
# Environment:
#   POLICY_FILE  Policy file. Default: .github/policies/servicenow-sdk-go-branch-protection.yml
#   GH_REPO      Repository. Default: the repository of the current checkout.
set -euo pipefail

policy_file="${POLICY_FILE:-.github/policies/servicenow-sdk-go-branch-protection.yml}"
repo="${GH_REPO:-$(gh repo view --json nameWithOwner --jq .nameWithOwner)}"

# The jq program that turns one branchProtectionRules entry into the ruleset
# fields it declares.
read -r -d '' desired_jq <<'EOF' || true
def unsupported($key):
  if .[$key] == true then error("\(.branchNamePattern): \($key): true isn't supported") else . end;

unsupported("restrictsPushes")
| unsupported("restrictsReviewDismissals")
| {
    pattern: .branchNamePattern,
    include: ["refs/heads/\(.branchNamePattern)"],
    # isAdminEnforced: false means repository admins can bypass the ruleset.
    bypass_actors: (
      if .isAdminEnforced then []
      else [{actor_id: 5, actor_type: "RepositoryRole", bypass_mode: "always"}]
      end
    ),
    rules: (
      [
        (if .allowsDeletions | not then {type: "deletion"} else empty end),
        (if .allowsForcePushes | not then {type: "non_fast_forward"} else empty end),
        (if .requiresLinearHistory then {type: "required_linear_history"} else empty end),
        (if .requiresCommitSignatures then {type: "required_signatures"} else empty end),
        (if .requiresPullRequestBeforeMerging then {
          type: "pull_request",
          parameters: {
            required_approving_review_count: .requiredApprovingReviewsCount,
            dismiss_stale_reviews_on_push: .dismissStaleReviews,
            require_code_owner_review: .requireCodeOwnersReview,
            require_last_push_approval: (.requireLastPushApproval // false),
            required_review_thread_resolution: .requiresConversationResolution
          }
        } else empty end),
        (if (.requiredStatusChecks // []) | length > 0 then {
          type: "required_status_checks",
          parameters: {
            strict_required_status_checks_policy: .requiresStrictStatusChecks,
            required_status_checks: [.requiredStatusChecks[] | {context: .}]
          }
        } else empty end)
      ] | sort_by(.type)
    )
  }
EOF

# The jq program that projects a live ruleset onto the fields that $desired
# declares, so the two can be compared with ==.
read -r -d '' project_jq <<'EOF' || true
{
  pattern: $desired.pattern,
  include: .conditions.ref_name.include,
  bypass_actors: [.bypass_actors[] | {actor_id, actor_type, bypass_mode}],
  rules: (
    [.rules[]
      | . as $live
      | ($desired.rules[] | select(.type == $live.type)) as $want
      | if $want.parameters == null then {type}
        else {
          type,
          parameters: (
            .parameters
            | with_entries(select(.key as $k | $want.parameters | has($k)))
            | if has("required_status_checks")
              then .required_status_checks |= [.[] | {context}]
              else . end
          )
        }
        end
    ]
    # Keep live rules that the file doesn't declare, so they show as drift.
    + [.rules[] | select(.type as $t | [$desired.rules[].type] | index($t) | not) | {type}]
    | sort_by(.type)
  )
}
EOF

# The jq program that builds the PUT payload: the live ruleset with the
# declared fields replaced. Rule types that the file doesn't declare are
# dropped; parameters that the file doesn't declare are kept.
read -r -d '' payload_jq <<'EOF' || true
{
  name,
  target,
  enforcement,
  conditions,
  bypass_actors: $desired.bypass_actors,
  rules: [
    $desired.rules[]
    | . as $want
    | ([$live.rules[] | select(.type == $want.type)] | first) as $have
    | if $want.parameters == null then {type}
      else {type, parameters: (($have.parameters // {}) * $want.parameters)}
      end
  ]
}
EOF

mode="${1:-}"
case "$mode" in
  check | plan | apply) ;;
  *)
    echo "Usage: $0 check|plan|apply" >&2
    exit 2
    ;;
esac

# Capture the output first: a failure inside a process substitution wouldn't
# stop the script.
desired_json="$(yq -o=json '.configuration.branchProtectionRules[]' "$policy_file" | jq -c "$desired_jq")"
mapfile -t desired_rules <<<"$desired_json"
if [[ -z "$desired_json" ]]; then
  echo "$policy_file declares no branchProtectionRules." >&2
  exit 1
fi
live_ids="$(gh api "repos/$repo/rulesets" --jq '.[] | select(.target == "branch") | .id')"
live_rulesets=()
for id in $live_ids; do
  live_rulesets+=("$(gh api "repos/$repo/rulesets/$id")")
done

drift=0
for desired in "${desired_rules[@]}"; do
  pattern="$(jq -r .pattern <<<"$desired")"
  include="$(jq -c .include <<<"$desired")"

  matches=()
  for ruleset in "${live_rulesets[@]}"; do
    if [[ "$(jq -c .conditions.ref_name.include <<<"$ruleset")" == "$include" ]]; then
      matches+=("$ruleset")
    fi
  done
  if [[ ${#matches[@]} -ne 1 ]]; then
    echo "$pattern: expected one live ruleset for $include, found ${#matches[@]}." >&2
    exit 1
  fi
  live="${matches[0]}"
  id="$(jq -r .id <<<"$live")"
  name="$(jq -r .name <<<"$live")"
  want="$(jq -S . <<<"$desired")"
  have="$(jq -S --argjson desired "$desired" "$project_jq" <<<"$live")"

  if [[ "$want" == "$have" ]]; then
    echo "$name ($pattern): matches the policy file."
    continue
  fi

  drift=1
  echo "$name ($pattern): differs from the policy file."
  diff -u --label "live $name" --label "policy $pattern" <(echo "$have") <(echo "$want") || true

  payload="$(jq --argjson desired "$desired" --argjson live "$live" "$payload_jq" <<<"$live")"
  case "$mode" in
    plan)
      echo "$name ($pattern): apply would send PUT repos/$repo/rulesets/$id with:"
      jq . <<<"$payload"
      ;;
    apply)
      gh api --method PUT "repos/$repo/rulesets/$id" --input - <<<"$payload" >/dev/null
      echo "$name ($pattern): applied."
      ;;
  esac
done

if [[ "$mode" == "check" && "$drift" -ne 0 ]]; then
  echo "Run 'mise run rulesets:apply' as a repository admin to apply the policy file." >&2
  exit 1
fi

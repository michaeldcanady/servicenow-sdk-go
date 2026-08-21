#!/usr/bin/env bash
# Apply the maintenance-line branch ruleset (ADR 011 rule 3 / #658).
#
# Codifies docs/release-branch-protection.md so the admin-side settings are
# reviewable and repeatable instead of click-through UI folklore:
#
#   scripts/apply-maintenance-ruleset.sh [--dry-run] [status-check-name...]
#
# Idempotent: if a ruleset named "maintenance-lines" exists it is updated in
# place; otherwise it is created. Requires gh authenticated as a repository
# administrator (rulesets are an admin API). Status checks are optional args:
# GitHub only offers checks that have already reported on the branch, so run
# once with none, open the bootstrap PR from the runbook, then re-apply with:
#
#   scripts/apply-maintenance-ruleset.sh \
#     "Check Branch Name" "Check Linked Issue" "Validate PR Title" \
#     "test (stable)" "lint"
set -euo pipefail

REPO="michaeldcanady/servicenow-sdk-go"
RULESET_NAME="maintenance-lines"

dry_run=false
checks=()
for arg in "$@"; do
  if [[ "$arg" == "--dry-run" ]]; then
    dry_run=true
  else
    checks+=("$arg")
  fi
done

# Build required_status_checks parameters. An empty list means "no checks
# enforced yet" (the bootstrap state); re-apply with names to enforce them.
check_rules='[]'
if (( ${#checks[@]} )); then
  check_rules=$(printf '{"context":"%s"}' "${checks[0]}")
  shift_check="${checks[0]}"
  for c in "${checks[@]:1}"; do
    check_rules+=",$(printf '{"context":"%s"}' "$c")"
  done
  check_rules="[${check_rules}]"
fi

# bypass_actors actor_id 5 with actor_type RepositoryRole = repository
# administrator, so emergencies stay possible without weakening the gate.
payload=$(cat <<EOF
{
  "name": "${RULESET_NAME}",
  "target": "branch",
  "enforcement": "active",
  "conditions": {
    "ref_name": {
      "include": ["refs/heads/release/v*"],
      "exclude": []
    }
  },
  "bypass_actors": [
    {"actor_id": 5, "actor_type": "RepositoryRole", "bypass_mode": "always"}
  ],
  "rules": [
    {"type": "deletion"},
    {"type": "non_fast_forward"},
    {"type": "pull_request", "parameters": {
      "required_approving_review_count": 1,
      "dismiss_stale_reviews_on_push": true,
      "require_code_owner_review": false,
      "require_last_push_approval": false,
      "required_review_thread_resolution": false,
      "allowed_merge_methods": ["merge", "squash", "rebase"]
    }},
    {"type": "required_status_checks", "parameters": {
      "strict_required_status_checks_policy": true,
      "required_status_checks": ${check_rules}
    }}
  ]
}
EOF
)

if $dry_run; then
  echo "$payload" | python3 -m json.tool
  exit 0
fi

existing_id=$(gh api "repos/${REPO}/rulesets" \
  --jq ".[] | select(.name == \"${RULESET_NAME}\") | .id" || true)

if [[ -n "${existing_id}" ]]; then
  echo "Updating existing ruleset '${RULESET_NAME}' (id ${existing_id})…"
  gh api -X PUT "repos/${REPO}/rulesets/${existing_id}" --input - <<< "$payload"
else
  echo "Creating ruleset '${RULESET_NAME}'…"
  gh api -X POST "repos/${REPO}/rulesets" --input - <<< "$payload"
fi

echo "Done. Verify under Settings → Rules → Rulesets."

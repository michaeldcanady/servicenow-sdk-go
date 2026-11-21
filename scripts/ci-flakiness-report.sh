#!/bin/bash
# Copyright (c) 2021 Michael Canady
# SPDX-License-Identifier: MIT

#
# Prints a Markdown report of per-job re-run and cancel rates for the past
# week, from the GitHub Actions API. The weekly ci-flakiness-report workflow
# appends the output to the job summary.
#
# Usage:
#   scripts/ci-flakiness-report.sh
#
# Environment:
#   GH_TOKEN      Token for `gh api`. Needs `actions: read`. In Actions this
#                 is the job's GITHUB_TOKEN; locally, `gh auth login` works.
#   REPO          owner/name. Defaults to $GITHUB_REPOSITORY, then the
#                 current repo as seen by `gh repo view`.
#   DAYS          Look-back window in days. Default 7.
#   RUN_IDS       Optional space-separated workflow run IDs. When set, the
#                 script reports on exactly those runs and skips the
#                 time-window query. Use it to check one run.
#   MIN_EXECUTIONS, RATE_THRESHOLD
#                 macOS follow-up threshold (see below). Defaults 20 and 10.
#
# Definitions. A job "execution" is one attempt of one job that reached a
# conclusion other than skipped. A "re-run" is an execution whose
# run_attempt is greater than 1. A "cancel" is an execution that concluded
# as cancelled. Rates are percentages of executions. Jobs are grouped by
# workflow name and job name, so a matrix leg is its own row.
#
# Two scopes: "main" (non-PR runs on the main branch) and "pull request"
# (pull_request runs). Release branches and tags are not counted.
#
# Caveat: ci.yml sets cancel-in-progress, so a newer push cancels an older
# PR run. Those cancels count here. Compare the pull request cancel rate
# with the main one before you blame a runner.
#
# macOS follow-up threshold: move the macOS jobs (the follow-up to
# SNSFG1-27) when a macOS job has at least MIN_EXECUTIONS executions and a
# combined re-run plus cancel rate of RATE_THRESHOLD percent or more in
# two consecutive weekly reports. The report flags rows that meet this.

set -euo pipefail

DAYS="${DAYS:-7}"
MIN_EXECUTIONS="${MIN_EXECUTIONS:-20}"
RATE_THRESHOLD="${RATE_THRESHOLD:-10}"
REPO="${REPO:-${GITHUB_REPOSITORY:-}}"
if [[ -z "$REPO" ]]; then
  REPO="$(gh repo view --json nameWithOwner --jq .nameWithOwner)"
fi

workdir="$(mktemp -d)"
trap 'rm -rf "$workdir"' EXIT

runs_file="$workdir/runs.jsonl"
jobs_file="$workdir/jobs.jsonl"
: >"$jobs_file"

if [[ -n "${RUN_IDS:-}" ]]; then
  for id in $RUN_IDS; do
    gh api "repos/$REPO/actions/runs/$id" \
      --jq '{id, workflow: .name, event, branch: .head_branch}' >>"$runs_file"
  done
else
  since="$(date -u -d "$DAYS days ago" +%Y-%m-%d)"
  gh api --paginate "repos/$REPO/actions/runs?created=>=$since&status=completed&per_page=100" \
    --jq '.workflow_runs[] | {id, workflow: .name, event, branch: .head_branch}' >"$runs_file"
fi

# Keep only the runs in scope, and tag each with its scope.
scoped_file="$workdir/scoped.jsonl"
jq -c '
  if .event == "pull_request" then . + {scope: "pull request"}
  elif .branch == "main" then . + {scope: "main"}
  else empty end' "$runs_file" >"$scoped_file"

run_count="$(wc -l <"$scoped_file")"
echo "Fetching jobs for $run_count runs from $REPO." >&2

while IFS= read -r run; do
  id="$(jq -r .id <<<"$run")"
  # filter=all includes jobs from earlier attempts, which is how re-runs show up.
  gh api --paginate "repos/$REPO/actions/runs/$id/jobs?filter=all&per_page=100" \
    --jq '.jobs[] | {name, conclusion, attempt: .run_attempt}' |
    jq -c --argjson run "$run" '. + {workflow: $run.workflow, scope: $run.scope}' >>"$jobs_file"
done <"$scoped_file"

echo "# CI flakiness report"
echo
echo "Repository: \`$REPO\`. Window: the last $DAYS days. Runs counted: $run_count."
echo "Generated: $(date -u +%Y-%m-%dT%H:%M:%SZ)."
echo
echo "A re-run is an execution with run_attempt above 1. A cancel is an execution that concluded as cancelled. Skipped jobs are excluded. Rates are percentages of executions."
echo "Rows marked **FLAG** are macOS jobs with at least $MIN_EXECUTIONS executions and a combined re-run plus cancel rate of $RATE_THRESHOLD% or more. Two consecutive flagged reports trigger the macOS follow-up."

for scope in "main" "pull request"; do
  echo
  echo "## Scope: $scope"
  echo
  jq -s -r --arg scope "$scope" --argjson min "$MIN_EXECUTIONS" --argjson thr "$RATE_THRESHOLD" '
    def pct(n; d): if d == 0 then 0 else (n * 1000 / d | round) / 10 end;
    [ .[] | select(.scope == $scope and .conclusion != null and .conclusion != "skipped") ]
    | group_by(.workflow + " / " + .name)
    | map({
        job: (.[0].workflow + " / " + .[0].name),
        n: length,
        reruns: (map(select(.attempt > 1)) | length),
        cancels: (map(select(.conclusion == "cancelled")) | length)
      })
    | map(. + {rerun_pct: pct(.reruns; .n), cancel_pct: pct(.cancels; .n)})
    | map(. + {flag: ((.job | test("macos"; "i")) and .n >= $min
                      and (pct(.reruns + .cancels; .n) >= $thr))})
    | sort_by(-(.rerun_pct + .cancel_pct), .job)
    | if length == 0 then "No completed jobs in this scope."
      else
        ["| Job | Executions | Re-runs | Re-run % | Cancels | Cancel % | Flag |",
         "| --- | ---: | ---: | ---: | ---: | ---: | --- |"]
        + map("| \(.job | gsub("\\|"; "\\|")) | \(.n) | \(.reruns) | \(.rerun_pct) | \(.cancels) | \(.cancel_pct) | \(if .flag then "**FLAG**" else "" end) |")
        | join("\n")
      end' "$jobs_file"
done

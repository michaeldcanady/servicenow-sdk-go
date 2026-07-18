# CI test failures never fail the build

- **Priority:** P0 — release blocker
- **Raised by:** Senior DevOps
- **Area:** CI/CD

## Problem

`.github/scripts/report-tests.sh` always exits 0, so the `test-go` job in `ci.yml` is
green even when tests fail. Every merge since this script was introduced has had an
unenforced test gate.

```bash
TEST_EXIT_CODE=0
go test -coverprofile=coverage.out -json -v ./... > test-output.json
echo "exit: $?"          # consumes $? but never stores it
...
exit $TEST_EXIT_CODE     # always 0
```

`TEST_EXIT_CODE` is initialized to 0 and never reassigned. `set -o pipefail` does not
help because the `go test` line has no pipe, and there is no `set -e`. The
`if: failure()` steps that post the sticky failure comment can therefore never trigger
either.

Tests currently pass locally (`go test ./...` is clean on `release/2.0`), so nothing
broken has slipped through *yet* — but the gate must be real before a 2.0 launch.

## Recommendation

```bash
go test -coverprofile=coverage.out -json -v ./... > test-output.json
TEST_EXIT_CODE=$?
```

Then keep `exit $TEST_EXIT_CODE` at the end. Add a deliberate failing-test dry run on a
branch to confirm the job goes red and the sticky comment posts.

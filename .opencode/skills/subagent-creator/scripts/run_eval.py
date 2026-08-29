#!/usr/bin/env python3
"""Run trigger evaluation for a subagent description.

Tests whether a subagent's description causes the orchestrator to invoke it
(via the `task` tool with a matching `subagent_type`) for a set of queries.
Outputs results as JSON.

This mirrors skill-creator's run_eval.py, but subagents trigger differently
from skills: there is no `Skill` tool or available-skill listing. Instead,
the orchestrator sees a list of agent types (name + description) and decides
whether to delegate via the `Task` tool with `subagent_type` set to one of
them. We detect that decision the same way skill-creator detects skill
triggering — by planting a uniquely-named candidate and watching the NDJSON
event stream for that name — just pointed at `.opencode/agents/<unique-name>.md`
and watching for a `task` tool call instead of a `Skill` tool call. Each query
runs headless via `opencode run --format json`.
"""

import argparse
import json
import os
import select
import subprocess
import sys
import time
import uuid
from concurrent.futures import ProcessPoolExecutor, as_completed
from pathlib import Path

from scripts.utils import parse_agent_md


def find_project_root() -> Path:
    """Find the project root by walking up from cwd looking for .opencode/.

    Mimics how opencode discovers its project root, so the temporary
    agent file we create ends up where `opencode run` will look for it.
    """
    current = Path.cwd()
    for parent in [current, *current.parents]:
        if (parent / ".opencode").is_dir():
            return parent
    return current


def run_single_query(
    query: str,
    agent_name: str,
    agent_description: str,
    timeout: int,
    project_root: str,
    model: str | None = None,
) -> bool:
    """Run a single query and return whether the candidate subagent was invoked.

    Creates a temporary subagent definition in `.opencode/agents/` so it
    appears in the orchestrator's available agent-types list, then runs
    `opencode run` with the raw query. Uses `--format json` to detect the
    `task` tool call as early as possible from the NDJSON event stream rather
    than waiting for the full final message.
    """
    unique_id = uuid.uuid4().hex[:8]
    clean_name = f"{agent_name}-test-{unique_id}"
    project_agents_dir = Path(project_root) / ".opencode" / "agents"
    agent_file = project_agents_dir / f"{clean_name}.md"

    try:
        project_agents_dir.mkdir(parents=True, exist_ok=True)
        agent_content = (
            f"---\n"
            f"name: {clean_name}\n"
            f"description: {agent_description}\n"
            f"tools: Read, Grep, Glob\n"
            f"---\n\n"
            f"You are a test subagent used only to measure whether your description "
            f"causes you to be invoked. Report back, in one sentence, what task you "
            f"were asked to do, then stop.\n"
        )
        agent_file.write_text(agent_content)

        cmd = [
            "opencode",
            "run",
            query,
            "--format", "json",
        ]
        if model:
            cmd.extend(["--model", model])

        env = dict(os.environ)

        process = subprocess.Popen(
            cmd,
            stdout=subprocess.PIPE,
            stderr=subprocess.DEVNULL,
            cwd=project_root,
            env=env,
        )

        triggered = False
        completed = False
        start_time = time.time()
        buffer = ""

        try:
            while time.time() - start_time < timeout:
                if process.poll() is not None:
                    remaining = process.stdout.read()
                    if remaining:
                        buffer += remaining.decode("utf-8", errors="replace")
                    break

                ready, _, _ = select.select([process.stdout], [], [], 1.0)
                if not ready:
                    continue

                chunk = os.read(process.stdout.fileno(), 8192)
                if not chunk:
                    break
                buffer += chunk.decode("utf-8", errors="replace")

                while "\n" in buffer and not completed:
                    line, buffer = buffer.split("\n", 1)
                    line = line.strip()
                    if not line:
                        continue

                    try:
                        event = json.loads(line)
                    except json.JSONDecodeError:
                        continue

                    # opencode `--format json` emits NDJSON events. A subagent
                    # spawn surfaces as a `task` tool call: message.part.updated
                    # events carry part.type "tool" / part.tool "task", with the
                    # planted subagent type in the part's state/input; agent
                    # lifecycle events also name the subagent type. The unique
                    # uuid suffix makes false positives negligible, so any such
                    # event line mentioning the planted name counts as a trigger.
                    part = event.get("part", {})
                    if event.get("type") == "agent.updated" or (
                        isinstance(part, dict) and part.get("type") == "tool"
                        and part.get("tool") == "task"
                    ):
                        if clean_name in json.dumps(event):
                            triggered = True
                    elif event.get("type") == "session.completed":
                        completed = True

            # Fallback: scan the raw event stream for the planted name, in
            # case the event shape drifts from the assumptions above. A
            # genuine reference to the subagent anywhere in the output still
            # counts as a trigger.
            if not triggered and clean_name in buffer:
                triggered = True

            # A non-zero exit without either a trigger signal or a clean
            # completion is a failed run (auth/config/CLI problem), not a
            # non-trigger — surface it instead of silently reporting False.
            exit_code = process.poll()
            if exit_code not in (None, 0) and not triggered and not completed:
                raise RuntimeError(
                    f"opencode run exited {exit_code} without invoking the agent"
                )
        finally:
            # Clean up process on any exit path (return, exception, timeout)
            if process.poll() is None:
                process.kill()
                process.wait()

        return triggered
    finally:
        if agent_file.exists():
            agent_file.unlink()


def run_eval(
    eval_set: list[dict],
    agent_name: str,
    description: str,
    num_workers: int,
    timeout: int,
    project_root: Path,
    runs_per_query: int = 1,
    trigger_threshold: float = 0.5,
    model: str | None = None,
) -> dict:
    """Run the full eval set and return results."""
    results = []

    with ProcessPoolExecutor(max_workers=num_workers) as executor:
        future_to_info = {}
        for item in eval_set:
            for run_idx in range(runs_per_query):
                future = executor.submit(
                    run_single_query,
                    item["query"],
                    agent_name,
                    description,
                    timeout,
                    str(project_root),
                    model,
                )
                future_to_info[future] = (item, run_idx)

        query_triggers: dict[str, list[bool]] = {}
        query_items: dict[str, dict] = {}
        for future in as_completed(future_to_info):
            item, _ = future_to_info[future]
            query = item["query"]
            query_items[query] = item
            if query not in query_triggers:
                query_triggers[query] = []
            try:
                query_triggers[query].append(future.result())
            except Exception as e:
                print(f"Warning: query failed: {e}", file=sys.stderr)
                query_triggers[query].append(False)

    for query, triggers in query_triggers.items():
        item = query_items[query]
        trigger_rate = sum(triggers) / len(triggers)
        should_trigger = item["should_trigger"]
        if should_trigger:
            did_pass = trigger_rate >= trigger_threshold
        else:
            did_pass = trigger_rate < trigger_threshold
        results.append({
            "query": query,
            "should_trigger": should_trigger,
            "trigger_rate": trigger_rate,
            "triggers": sum(triggers),
            "runs": len(triggers),
            "pass": did_pass,
        })

    passed = sum(1 for r in results if r["pass"])
    total = len(results)

    return {
        "agent_name": agent_name,
        "description": description,
        "results": results,
        "summary": {
            "total": total,
            "passed": passed,
            "failed": total - passed,
        },
    }


def main():
    parser = argparse.ArgumentParser(description="Run trigger evaluation for a subagent description")
    parser.add_argument("--eval-set", required=True, help="Path to eval set JSON file")
    parser.add_argument("--agent-path", required=True, help="Path to the subagent .md file (or its containing directory)")
    parser.add_argument("--description", default=None, help="Override description to test")
    parser.add_argument("--num-workers", type=int, default=10, help="Number of parallel workers")
    parser.add_argument("--timeout", type=int, default=30, help="Timeout per query in seconds")
    parser.add_argument("--runs-per-query", type=int, default=3, help="Number of runs per query")
    parser.add_argument("--trigger-threshold", type=float, default=0.5, help="Trigger rate threshold")
    parser.add_argument("--model", default=None, help="Model to use for opencode run (default: user's configured model)")
    parser.add_argument("--verbose", action="store_true", help="Print progress to stderr")
    args = parser.parse_args()

    eval_set = json.loads(Path(args.eval_set).read_text())
    agent_path = Path(args.agent_path)

    name, original_description, _content = parse_agent_md(agent_path)
    description = args.description or original_description
    project_root = find_project_root()

    if args.verbose:
        print(f"Evaluating: {description}", file=sys.stderr)

    output = run_eval(
        eval_set=eval_set,
        agent_name=name,
        description=description,
        num_workers=args.num_workers,
        timeout=args.timeout,
        project_root=project_root,
        runs_per_query=args.runs_per_query,
        trigger_threshold=args.trigger_threshold,
        model=args.model,
    )

    if args.verbose:
        summary = output["summary"]
        print(f"Results: {summary['passed']}/{summary['total']} passed", file=sys.stderr)
        for r in output["results"]:
            status = "PASS" if r["pass"] else "FAIL"
            rate_str = f"{r['triggers']}/{r['runs']}"
            print(f"  [{status}] rate={rate_str} expected={r['should_trigger']}: {r['query'][:70]}", file=sys.stderr)

    print(json.dumps(output, indent=2))


if __name__ == "__main__":
    main()
---
name: subagent-creator
description: Create a new Claude Code subagent (.claude/agents/*.md) for this repo, or test/improve an existing one. Use whenever the user asks to "create a subagent for X", "add a new agent to .claude/agents", "I need a review/testing/scaffolding agent", wants a dedicated agent to run proactively after some kind of change, or wants to benchmark/tune an existing agent's description so it triggers reliably. Covers the full loop: interview the user, draft the agent's frontmatter and system prompt, spawn test runs via the Agent tool, grade and benchmark the results, show them in an HTML viewer, and iterate.
---

# Subagent Creator

Helps write and iteratively improve `.claude/agents/*.md` subagent definitions
for this repo, and mirrors the authoring→eval→iterate loop used by the
upstream `skill-creator` skill, retargeted to subagents. Everything this
skill needs (eval runner, grader, benchmark aggregator, HTML viewer) lives
under this skill's own `scripts/`, `eval-viewer/`, `agents/`, and
`references/` directories — it has no dependency on the `skill-creator`
plugin being installed.

## How a subagent differs from a skill (read this first)

A subagent is a single file, `.claude/agents/<name>.md` — YAML frontmatter
(`name`, `description`, optionally `tools` as a comma-separated allowlist and
`model`), then a markdown body that IS the subagent's system prompt, written
directly as instructions to that agent (no wrapper). See
`.claude/agents/api-module-consistency-reviewer.md` and
`.claude/agents/godog-test-writer.md` in this repo for the exact shape to
match.

Triggering also works differently from skills. There's no `Skill` tool or
`available_skills` listing — instead, the orchestrating Claude is shown a
list of available agent types (just `name` + `description`, the same
one-liner you write in frontmatter) and decides whether to call the `Agent`
tool with `subagent_type` set to one of them, based purely on that
description matching the task at hand. It never sees the body up front. This
means the description carries even more weight than a skill's does — it is
the *only* signal for delegation, so treat it as a discoverability problem in
the same spirit as the writing tips below, adapted for "when should the
orchestrator hand this task off" rather than "when should Claude read this
file."

## Communicating with the user

Match your vocabulary to the user's familiarity. "Trigger", "frontmatter",
and "system prompt" are fine for most users of this repo; "assertion" and
"benchmark" are borderline — check for cues before assuming familiarity.

## Creating a subagent

### Capture intent

Ask one question at a time, waiting for each answer:

1. What should this subagent do, concretely? (e.g. "review new `*api`
   packages for structural consistency" — not just "help with APIs")
2. When should it trigger — what would the user or orchestrator actually
   say/need? Should it also be invoked *proactively* by Claude without the
   user naming it (e.g. "after adding a new package", "before opening a
   PR")? If so, say that explicitly in the description — orchestrators
   default to *not* delegating unless a description clearly earns it.
3. What tools does it need? Default to the full toolset (omit `tools`
   entirely) unless there's a reason to restrict it — e.g. a pure-review
   agent that should never edit files gets `tools: Read, Grep, Glob, Bash`.
4. Does it need a non-default `model` (e.g. a cheap/fast reviewer might want
   `model: haiku`, a hard architecture-review agent might want `model:
   opus`)? Default: omit, inherit the caller's model.
5. Is this authoring-only, or should it also get an eval loop (test
   prompts, grading, benchmarking)? Agents with a checkable job (review
   findings, generated test files, structural conformance) benefit from
   evals; a purely stylistic/conversational agent may not need them.

### Interview and research

Ask about edge cases, expected inputs, what a *wrong* output looks like, and
whether there's an existing convention to hold the new agent to (in this
repo: the "Canonical pattern" section of
`api-module-consistency-reviewer.md` is a good model of "here's the standard,
here's what drift looks like"). Read a couple of existing files in
`.claude/agents/` before drafting — don't invent a new shape from scratch.

### Write the agent file

- **name**: kebab-case, must match the filename (`<name>.md`).
- **description**: what it does AND when to use it — this is the *entire*
  triggering signal (see above). Write it a little "pushy": if this agent
  should be used proactively, say so in the description itself, don't rely
  on the body — the orchestrator never reads the body before deciding.
- **tools**: omit for "all tools"; otherwise a comma-separated allowlist.
- **model**: omit to inherit; set explicitly only with a reason.
- **body**: the system prompt, written directly to the agent ("You are
  reviewing...", not "This skill will..."). Explain the *why* behind rules
  you give it, the same way you would explain to a capable colleague — a
  smart model gets more out of "we hand-write this instead of generating it
  because X" than a bare "ALWAYS do X."

Validate the frontmatter mechanically before running anything:
```bash
python3 .claude/skills/subagent-creator/scripts/quick_validate.py .claude/agents/<name>.md
```

## Running and evaluating test cases

Skip this whole section if the user just wants a quick one-off agent with no
eval loop — write the file, validate it, done.

Otherwise, come up with 2-3 realistic test prompts — the kind of task a real
user would actually hand off. Save them (without assertions yet) to
`evals/evals.json` next to a scratch workspace, e.g.:

```json
{
  "agent_name": "sql-migration-reviewer",
  "evals": [
    {"id": 1, "prompt": "Review migration 0042_add_index.sql for locking issues on a 50M-row table", "expected_output": "Flags the lock risk and suggests CONCURRENTLY", "files": []}
  ]
}
```

Put results in `<agent-name>-workspace/iteration-N/eval-<id>/` (sibling to
the skill, same convention as skill-creator). For each test case spawn two
runs in the same turn:

- **With the candidate subagent**: use the `Agent` tool with `subagent_type`
  set to the new agent's name, save outputs to `.../with_agent/outputs/`.
- **Baseline**: the same prompt handled by a `general-purpose` agent (no
  specialized subagent), saved to `.../without_agent/outputs/`. If you're
  *improving* an existing agent instead of creating a new one, snapshot the
  old file first and use it as the baseline (`old_agent`/`new_agent`
  instead of `without_agent`/`with_agent`) — `aggregate_benchmark.py`
  discovers whatever config directory names you use, so either pair works.

Write `eval_metadata.json` per case (assertions can start empty):
```json
{"eval_id": 0, "eval_name": "descriptive-name", "prompt": "...", "assertions": []}
```

While runs are in flight, draft assertions — objectively checkable claims
about the output (see `references/schemas.md` for the exact shape). Skip
assertions for subjective judgments; that's what human review is for.

As each run finishes, save the task-notification's `total_tokens`/
`duration_ms` to `timing.json` in that run's directory — it's only available
in that notification, nowhere else.

Once everything's done:

1. **Grade** each run against its assertions — spawn a subagent that reads
   `agents/grader.md` and writes `grading.json` (must use the `text`/
   `passed`/`evidence` field names — the viewer depends on them exactly).
2. **Aggregate**:
   ```bash
   python3 -m scripts.aggregate_benchmark <workspace>/iteration-N --agent-name <name>
   ```
   (run from inside `.claude/skills/subagent-creator/`, or adjust
   `PYTHONPATH`/cwd accordingly — it's a package-relative import).
3. **Analyze** — read `agents/analyzer.md` for what to look for (assertions
   that never fail regardless of config, high-variance evals, time/token
   tradeoffs).
4. **View**:
   ```bash
   python3 eval-viewer/generate_review.py <workspace>/iteration-N \
     --agent-name "<name>" --benchmark <workspace>/iteration-N/benchmark.json
   ```
   Add `--previous-workspace <workspace>/iteration-<N-1>` from iteration 2
   on. If there's no display available, use `--static <output_path>` instead
   and read `feedback.json` back once the user downloads it after clicking
   "Submit All Reviews".
5. Tell the user where to look: "Outputs" tab for qualitative review +
   feedback box, "Benchmark" tab for the quantitative comparison.

## Improving the subagent

Read `feedback.json` once the user says they're done. Empty feedback means
that case looked fine.

- **Generalize, don't overfit.** If feedback points at a stubborn gap, don't
  bolt on an ever-growing list of special cases in the body — that's the
  same failure mode as description bloat, just in the system prompt instead
  of the trigger text. Look for a structurally different framing instead.
- **Keep it lean.** Read the transcripts, not just final outputs — if the
  agent wasted turns on something unproductive, that's a sign the body is
  telling it to do the wrong thing, or not explaining why the right thing
  matters.
- **Explain the why.** Same principle as writing the body the first time:
  rigid ALL-CAPS rules are a signal to go back and explain the reasoning
  instead.
- **Look for repeated scaffolding.** If every test run has the subagent
  independently writing the same kind of helper script or doing the same
  multi-step dance, that's worth calling out to the agent directly in its
  body (or bundling as a reference), so future invocations don't reinvent it.

Then rerun all test cases into `iteration-<N+1>/` (including baselines),
relaunch the viewer with `--previous-workspace`, and repeat until the user's
happy or feedback goes quiet.

## Description optimization

If the agent's own description needs tuning for reliable triggering (not the
body — see "How a subagent differs from a skill" above), generate ~20 eval
queries (8-10 should-trigger, 8-10 should-not-trigger, favoring realistic
near-misses over obviously-irrelevant negatives — see skill-creator's own
guidance on writing good eval queries if you need the fuller rationale),
have the user review them, then run:

```bash
python3 -m scripts.run_loop \
  --eval-set <path-to-trigger-eval.json> \
  --agent-path .claude/agents/<name>.md \
  --model <model-id-powering-this-session> \
  --max-iterations 5 \
  --verbose
```

This splits queries 60/40 train/test, evaluates the current description
(each query run 3x for a reliable trigger rate — detected by watching for an
`Agent` tool_use naming a temporary planted agent, see `scripts/run_eval.py`
for the mechanics), proposes improvements via `scripts/improve_description.py`,
and re-evaluates up to 5 times, picking the best description by *test* score
to avoid overfitting. Take `best_description` from the JSON output, update
the agent file's frontmatter, and show the user before/after with scores.

## Done — no packaging step

Unlike skills, subagents aren't packaged or distributed as a bundle — the
finished `.claude/agents/<name>.md` file is already live and ready to use
the moment it's saved and passes `quick_validate.py`. There's nothing
further to build or install.

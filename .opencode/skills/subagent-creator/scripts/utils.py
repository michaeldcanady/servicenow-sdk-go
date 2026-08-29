"""Shared utilities for subagent-creator scripts."""

from pathlib import Path


def parse_agent_md(agent_path: Path) -> tuple[str, str, str]:
    """Parse a subagent definition file, returning (name, description, full_content).

    Unlike a SKILL.md (which lives inside a skill directory), a subagent
    definition is a single loose ``.md`` file, typically
    ``.claude/agents/<name>.md``. ``agent_path`` may point directly at that
    file, or at a directory containing exactly one ``.md`` file.
    """
    if agent_path.is_dir():
        candidates = sorted(agent_path.glob("*.md"))
        if not candidates:
            raise ValueError(f"No .md file found in {agent_path}")
        agent_file = candidates[0]
    else:
        agent_file = agent_path

    content = agent_file.read_text()
    lines = content.split("\n")

    if lines[0].strip() != "---":
        raise ValueError(f"{agent_file} missing frontmatter (no opening ---)")

    end_idx = None
    for i, line in enumerate(lines[1:], start=1):
        if line.strip() == "---":
            end_idx = i
            break

    if end_idx is None:
        raise ValueError(f"{agent_file} missing frontmatter (no closing ---)")

    name = ""
    description = ""
    frontmatter_lines = lines[1:end_idx]
    i = 0
    while i < len(frontmatter_lines):
        line = frontmatter_lines[i]
        if line.startswith("name:"):
            name = line[len("name:"):].strip().strip('"').strip("'")
        elif line.startswith("description:"):
            value = line[len("description:"):].strip()
            # Handle YAML multiline indicators (>, |, >-, |-)
            if value in (">", "|", ">-", "|-"):
                continuation_lines: list[str] = []
                i += 1
                while i < len(frontmatter_lines) and (frontmatter_lines[i].startswith("  ") or frontmatter_lines[i].startswith("\t")):
                    continuation_lines.append(frontmatter_lines[i].strip())
                    i += 1
                description = " ".join(continuation_lines)
                continue
            else:
                description = value.strip('"').strip("'")
        i += 1

    if not name:
        name = agent_file.stem

    return name, description, content

#!/usr/bin/env python3
"""
Quick validation script for subagent .md files - minimal version
"""

import re
import sys
from pathlib import Path

KNOWN_MODEL_ALIASES = {"sonnet", "opus", "haiku", "inherit"}


def _parse_frontmatter(frontmatter_text: str) -> dict:
    """Parse simple flat `key: value` frontmatter without a YAML dependency.

    Subagent frontmatter is flat (name/description/tools/model, all
    single-line strings except description's occasional YAML block-scalar
    form), so a full YAML parser is unneeded — this keeps the skill
    dependency-free.
    """
    frontmatter: dict = {}
    lines = frontmatter_text.split("\n")
    i = 0
    while i < len(lines):
        line = lines[i]
        if not line.strip() or line.strip().startswith("#"):
            i += 1
            continue
        if ":" not in line:
            raise ValueError(f"Malformed frontmatter line (no ':'): {line!r}")
        key, _, value = line.partition(":")
        key = key.strip()
        value = value.strip()
        if value in (">", "|", ">-", "|-"):
            continuation = []
            i += 1
            while i < len(lines) and (lines[i].startswith("  ") or lines[i].startswith("\t") or not lines[i].strip()):
                if lines[i].strip():
                    continuation.append(lines[i].strip())
                i += 1
            frontmatter[key] = " ".join(continuation)
            continue
        frontmatter[key] = value.strip('"').strip("'")
        i += 1
    return frontmatter


def validate_agent(agent_path):
    """Basic validation of a subagent definition file."""
    agent_path = Path(agent_path)

    if agent_path.is_dir():
        candidates = sorted(agent_path.glob("*.md"))
        if not candidates:
            return False, f"No .md file found in {agent_path}"
        agent_file = candidates[0]
    else:
        agent_file = agent_path
        if not agent_file.exists():
            return False, f"{agent_file} not found"

    content = agent_file.read_text()
    if not content.startswith('---'):
        return False, "No YAML frontmatter found"

    match = re.match(r'^---\n(.*?)\n---', content, re.DOTALL)
    if not match:
        return False, "Invalid frontmatter format"

    frontmatter_text = match.group(1)
    body = content[match.end():].strip()

    try:
        frontmatter = _parse_frontmatter(frontmatter_text)
    except ValueError as e:
        return False, f"Invalid frontmatter: {e}"

    ALLOWED_PROPERTIES = {'name', 'description', 'tools', 'model'}
    unexpected_keys = set(frontmatter.keys()) - ALLOWED_PROPERTIES
    if unexpected_keys:
        return False, (
            f"Unexpected key(s) in frontmatter: {', '.join(sorted(unexpected_keys))}. "
            f"Allowed properties are: {', '.join(sorted(ALLOWED_PROPERTIES))}"
        )

    if 'name' not in frontmatter:
        return False, "Missing 'name' in frontmatter"
    if 'description' not in frontmatter:
        return False, "Missing 'description' in frontmatter"

    name = frontmatter.get('name', '')
    if not isinstance(name, str):
        return False, f"Name must be a string, got {type(name).__name__}"
    name = name.strip()
    if name:
        if not re.match(r'^[a-z0-9-]+$', name):
            return False, f"Name '{name}' should be kebab-case (lowercase letters, digits, and hyphens only)"
        if name.startswith('-') or name.endswith('-') or '--' in name:
            return False, f"Name '{name}' cannot start/end with hyphen or contain consecutive hyphens"
        if agent_file.stem != name:
            return False, f"File name '{agent_file.name}' should match the 'name' field ('{name}.md')"

    description = frontmatter.get('description', '')
    if not isinstance(description, str):
        return False, f"Description must be a string, got {type(description).__name__}"
    description = description.strip()
    if not description:
        return False, "Description cannot be empty"
    if '<' in description or '>' in description:
        return False, "Description cannot contain angle brackets (< or >)"
    if len(description) > 1024:
        return False, f"Description is too long ({len(description)} characters). Maximum is 1024 characters."
    if len(description) < 20:
        return False, (
            f"Description is only {len(description)} characters — likely too generic to "
            f"trigger reliably. Describe both what the subagent does and when to use it."
        )

    tools = frontmatter.get('tools')
    if tools is not None:
        if not isinstance(tools, str):
            return False, f"'tools' must be a comma-separated string, got {type(tools).__name__}"
        tool_names = [t.strip() for t in tools.split(",")]
        if any(not t for t in tool_names):
            return False, "'tools' has an empty entry — check for stray commas"

    model = frontmatter.get('model')
    if model is not None:
        if not isinstance(model, str) or not model.strip():
            return False, "'model' must be a non-empty string"
        elif model.strip() not in KNOWN_MODEL_ALIASES:
            print(
                f"Warning: 'model: {model}' is not one of the commonly used aliases "
                f"({', '.join(sorted(KNOWN_MODEL_ALIASES))}) — double check it's intentional.",
                file=sys.stderr,
            )

    if not body:
        return False, "Subagent body (the system prompt) is empty"

    return True, "Subagent definition is valid!"


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python quick_validate.py <agent .md file or directory>")
        sys.exit(1)

    valid, message = validate_agent(sys.argv[1])
    print(message)
    sys.exit(0 if valid else 1)

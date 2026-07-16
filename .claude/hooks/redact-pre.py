#!/usr/bin/env python3
import json, os, re, sys

SECRET_PATTERNS = [
    re.compile(r'(?i)\b[A-Z0-9]{20,}[_-]?[A-Z0-9]{10,}\b'),          # generic long tokens
    re.compile(r'(?i)\b(?:api|secret|token|key|passwd|password)\s*[:=]\s*["\']?([^\s"\']+)'),
    re.compile(r'(?i)sk-[a-z0-9]{20,}'),                              # common key prefix
]

REDACT = "★★★REDACTED★★★"

# Filenames whose content is never shown to Claude unredacted.
SECRET_BEARING_FILENAMES = {
    "settings.local.json",
}

REDACTED_MIRROR_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))  # .claude/


def redact_text(s: str) -> str:
    for pat in SECRET_PATTERNS:
        s = pat.sub(REDACT, s)
    return s


def make_redacted_copy(src_path: str) -> str:
    with open(src_path, "r", encoding="utf-8") as f:
        content = f.read()
    dst_path = os.path.join(REDACTED_MIRROR_DIR, "." + os.path.basename(src_path) + ".redacted")
    with open(dst_path, "w", encoding="utf-8") as f:
        f.write(redact_text(content))
    return dst_path


def main():
    try:
        data = json.load(sys.stdin)
    except Exception as e:
        print(f"Invalid JSON: {e}", file=sys.stderr)
        sys.exit(1)

    tool_name = data.get("tool_name", "")
    tool_input = data.get("tool_input", {})

    if tool_name == "Bash":
        cmd = tool_input.get("command", "") or ""
        if re.search(r'\b(curl|wget)\b.*\s(-H|--header)\s.*(authorization|api-key)', cmd, re.I):
            print("Blocking Bash call that would echo auth headers", file=sys.stderr)
            sys.exit(2)
        if re.search(r'(cat|grep|less|more|printenv|env)[^|]*settings\.local\.json', cmd, re.I) or "GITHUB_PERSONAL_ACCESS_TOKEN" in cmd:
            print("Blocking Bash call that would dump secret-bearing settings.local.json", file=sys.stderr)
            sys.exit(2)

    out = {"hookSpecificOutput": {"hookEventName": "PreToolUse"}}

    if tool_name == "Read":
        fp = tool_input.get("file_path", "") or ""
        base = os.path.basename(fp)
        if base in SECRET_BEARING_FILENAMES:
            try:
                redacted_path = make_redacted_copy(fp)
                out["hookSpecificOutput"]["permissionDecision"] = "allow"
                out["hookSpecificOutput"]["permissionDecisionReason"] = (
                    "Redirected to a redacted copy — secret-like values are masked."
                )
                out["hookSpecificOutput"]["updatedInput"] = {"file_path": redacted_path}
            except OSError as e:
                out["hookSpecificOutput"]["permissionDecision"] = "deny"
                out["hookSpecificOutput"]["permissionDecisionReason"] = f"Could not create redacted copy: {e}"
        elif fp.endswith((".md", ".mdx", ".txt")):
            out["hookSpecificOutput"]["permissionDecision"] = "allow"
            out["hookSpecificOutput"]["permissionDecisionReason"] = "Documentation read auto approved"

    print(json.dumps(out))
    sys.exit(0)


if __name__ == "__main__":
    main()

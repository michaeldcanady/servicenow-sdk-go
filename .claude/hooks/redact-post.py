#!/usr/bin/env python3
import json, os, sys

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
from redact_common import redact_text as _redact_text, REDACT  # noqa: E402


def redact_text(s: str) -> str:
    s = _redact_text(s)
    if len(s) > 50000:  # safety middle truncate for huge outputs
        s = s[:25000] + "\n... " + REDACT + " ..." + s[-25000:]
    return s

def main():
    try:
        data = json.load(sys.stdin)
    except Exception as e:
        print(f"Invalid JSON: {e}", file=sys.stderr)
        return 1

    tool_resp = data.get("tool_response", {})

    # PostToolUse hooks can't rewrite the tool_response that's already in the
    # transcript, so surface the redacted content via additionalContext and
    # warn Claude not to repeat the raw values instead of claiming the
    # transcript itself was sanitized.
    redacted_sections = []
    for k in ("stdout", "stderr", "body", "content"):
        val = tool_resp.get(k)
        if isinstance(val, str):
            redacted = redact_text(val)
            if redacted != val:
                redacted_sections.append(f"--- {k} (redacted) ---\n{redacted}")

    if redacted_sections:
        additional_context = (
            "Secret-like values were detected in this tool's output. "
            "Do not repeat, quote, or forward the raw values above verbatim. "
            "Redacted version:\n" + "\n".join(redacted_sections)
        )
    else:
        additional_context = "No secret-like values detected in this tool's output."

    out = {
        "hookSpecificOutput": {
            "hookEventName": "PostToolUse",
            "additionalContext": additional_context
        }
    }

    print(json.dumps(out))
    return 0

if __name__ == "__main__":
    sys.exit(main())
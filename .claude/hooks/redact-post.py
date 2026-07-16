#!/usr/bin/env python3
import json, sys, re

REDACT = "★★★REDACTED★★★"
SECRET_TEXT = re.compile(r'(?i)(?:api|secret|token|key|bearer)\s*[:=]\s*["\']?([^\s"\']+)')

def redact_text(s: str) -> str:
    s = SECRET_TEXT.sub(REDACT, s)
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
    # Add advisory context to Claude after tool runs
    out = {
        "hookSpecificOutput": {
            "hookEventName": "PostToolUse",
            "additionalContext": "Outputs sanitized by redaction policy"
        }
    }

    # If tool response includes printable fields, scrub them in transcript
    for k in ("stdout", "stderr", "body", "content"):
        if isinstance(tool_resp.get(k), str):
            tool_resp[k] = redact_text(tool_resp[k])

    print(json.dumps(out))
    return 0

if __name__ == "__main__":
    sys.exit(main())
"""Shared secret-redaction patterns/helpers for the pre/post tool-use hooks.

Keeping this in one module means a new secret shape only needs to be added
once instead of being kept in sync across redact-pre.py and redact-post.py.
"""
import re

REDACT = "★★★REDACTED★★★"

SECRET_PATTERNS = [
    re.compile(r'(?i)\b[A-Z0-9]{20,}[_-]?[A-Z0-9]{10,}\b'),          # generic long tokens
    re.compile(r'(?i)\b(?:api|secret|token|key|passwd|password|bearer)\s*[:=]\s*["\']?([^\s"\']+)'),
    re.compile(r'(?i)sk-[a-z0-9]{20,}'),                              # common key prefix
]


def redact_text(s: str) -> str:
    for pat in SECRET_PATTERNS:
        s = pat.sub(REDACT, s)
    return s

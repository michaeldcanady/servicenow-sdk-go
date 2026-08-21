---
name: google-doc-style
description: Applies Google's developer documentation style guide when writing or editing documentation, READMEs, code comments, ADRs, markdown files, or any prose in this repo. Use whenever the user asks to "write docs", "improve this README", "review the documentation", "make this clearer", "write a guide", "polish this text", or involves composing or editing English-language technical prose, headings, lists, code comments, or presentation-ready content. Also use proactively when the agent's own output contains multi-sentence prose or when a task involves explaining code or architecture in human-readable form.
---

# Google Developer Documentation Style

This skill applies [Google's developer documentation style guide](https://developers.google.com/style) to all English-language prose in this repository. The full guide is the authoritative reference; this skill distills the rules most relevant to SDK documentation, code comments, ADRs, READMEs, and markdown pages so you can apply them without leaving your workflow.

## Core philosophy

- **Clarity and consistency above all.** Follow this guide, then project conventions (CLAUDE.md, ADRs), then general references (Merriam-Webster, Chicago Manual of Style).
- **Sound like a knowledgeable friend.** Conversational, friendly, respectful — not formal, not frivolous, not preachy.
- **Write for a global audience.** Simple sentences, no culturally specific references, no idioms. Readers may have varying English proficiency.

## Voice and tone

| Do | Don't |
|---|---|
| Sound natural and approachable | Use slang, internet abbreviations, or pop-culture references |
| Let personality show through subtly | Be cute, zany, or try to be super-entertaining |
| Get to the point quickly | Pad with filler or placeholder phrases |
| Be polite without overdoing it | Use "please" in instructions ("Click **Save**" not "Please click **Save**") |

### Words and phrases to avoid

- **Jargon and buzzwords** unless you define them on first use
- **Figurative language** — metaphors, idioms, ableist language; use literal, precise terms
- **Filler**: *please note*, *at this time*, *simply*, *it's easy*, *just*, *note that*
- **Exclamation marks** — almost never appropriate in technical docs
- **"Let's"** — don't write "Let's do X"; write "Do X" or "You can do X"
- **Starting every sentence the same way** — vary sentence openers

## Language and grammar

### Second person

Address the reader as **you**. Don't use **we** or **let's**.

```
// Good
You can configure the timeout by passing an option.

// Bad
We can configure the timeout by passing an option.
```

### Active voice

Use active voice. Make clear who performs the action.

```
// Good
The server sends an acknowledgment.

// Bad
An acknowledgment is sent by the server.
```

**Passive is acceptable when:**
- The actor is irrelevant: "The database was purged in January."
- You're emphasizing the object: "The file is saved."
- You're de-emphasizing the actor: "Over 50 conflicts were found."

### Present tense

Describe what something *does*, not what it *will do*.

```
// Good
This method returns the record.

// Bad
This method will return the record.
```

### Put conditions before instructions

Don't bury conditions at the end of a sentence.

```
// Good
If the record doesn't exist, the API returns a 404.

// Bad
The API returns a 404 if the record doesn't exist.
```

### American English

Use standard American English spelling and punctuation:
- Serial (Oxford) comma: "a, b, and c"
- Quotation marks: "like this" (not 'like this')
- Dates: unambiguous format — "January 1, 2025" or "2025-01-01" (not "1/1/25")

## Formatting and structure

### Headings and titles

- **Sentence case**: "Create a table record" (not "Create a Table Record")
- Headings are titles, not instructions — no terminal punctuation
- Don't skip heading levels (h1 → h3 is wrong; use h1 → h2 → h3)
- Mark optional sections with "(optional)" in the heading if needed

### Lists

- **Numbered lists** for sequences or steps (order matters)
- **Bulleted lists** for unordered items
- **Description lists** for key-value pairs (where the format supports it)
- Keep parallel structure across list items (all start the same way, same grammatical form)
- Don't use inconsistent end punctuation in list items — either all end with periods or all without

### Code references

- **Inline code** (backticks): method names, class names, filenames, HTTP status codes, console output, placeholders, parameters, IP addresses, port numbers
- **Code blocks** (triple backticks): multi-line code samples
- **Bold** (`**`): UI elements only ("click **Save**")
- **Italics** (`_`): introducing a term for the first time, or discussing a word as a word ("the term _backing store_ refers to...")
- **Underline**: reserve for link text only

### Links

- Use descriptive link text — never "click here" or "this page"
- Put punctuation outside link text
- Format: `[descriptive text](URL)` not `[URL](URL)`

### Numbers

- Spell out numbers zero through nine; use numerals for 10 and above
- Exception: always use numerals for code-related numbers, measurements, or when mixing with other numerals in a sentence
- Use commas in large numbers: 1,000 (not 1.000)

### Exclamation marks

Avoid them. Almost never appropriate in developer documentation.

## Content principles

### Write for a global audience

- Avoid idioms ("out of the box", "on the same page", "touch base")
- Avoid culturally specific references
- Use short, simple sentences
- Define technical terms on first use
- Consistent terminology — use the same word for the same concept throughout

### Write inclusive documentation

- Use inclusive language (avoid "whitelist/blacklist", "master/slave", "grandfathered")
- Don't use figurative language that could be ableist ("sanity check" → "check", "blind spot" → "gap")
- See Google's word list for specific term guidance

### Don't pre-announce

Never document features that don't exist yet. Documentation should describe what *is*, not what *will be*.

### Timeless documentation

Avoid time-sensitive language ("recently", "currently", "now"). The documentation should be accurate when read at any point.

### Prescriptive, not descriptive

Tell users what to do, not what exists. Write "Use X" not "You can use X" when X is the recommended approach.

```
// Good (prescriptive)
Use the `ByID` method to retrieve a single record.

// Less good (descriptive)
You can use the `ByID` method to retrieve a single record.
```

## Procedure formatting

- Start with a brief intro sentence explaining what the procedure does
- Number each step
- One action per step
- Start each step with a verb ("Click **Save**", "Run `go build`")
- Put the condition before the action: "If prompted, click **OK**" (not "Click **OK** if prompted")
- End each step with a period

## Code samples

- Every code sample should be runnable — don't include pseudocode without labeling it
- Keep samples short and focused on the concept being explained
- Use the repo's own patterns (see `website/snippets/*.go` for existing examples)
- Include necessary imports
- Use realistic variable names, not `foo`/`bar`/`baz`
- Format code samples in fenced code blocks with a language tag

## Common mistakes to fix

| Wrong | Right |
|---|---|
| `it's` (possessive) | `its` (possessive) — `it's` = `it is` |
| `which` (restrictive) | `that` (restrictive), `which` (non-restrictive, with comma) |
| `ie.` / `eg.` | `that is` / `for example` — or use the Latin abbreviations only in parenthetical contexts |
| `utilize` | `use` |
| `in order to` | `to` |
| `the reason is because` | `the reason is that` |
| `a number of` | `several` or `many` |
| `essentially` / `basically` | remove — adds nothing |
| `simple` / `easily` / `just` | remove — these are filler that can frustrate readers for whom the task isn't simple |
| `may` (possibility) | `might` (possibility) or `can` (ability) — `may` implies permission |

## Checklist for reviewing prose

Before marking documentation as done, verify:

- [ ] Second person ("you"), not first person ("we") or third person
- [ ] Active voice (with justified passive exceptions)
- [ ] Present tense for describing current behavior
- [ ] Sentence case in all headings
- [ ] Serial comma throughout
- [ ] Code references in backticks, UI elements in bold
- [ ] Descriptive link text (no "click here")
- [ ] No filler words (*simply*, *just*, *please note*, *note that*)
- [ ] No exclamation marks (unless genuinely warranted)
- [ ] No figurative language or idioms
- [ ] Consistent terminology throughout the document
- [ ] Numbers: spelled out 0–9, numerals for 10+
- [ ] Conditions placed before instructions
- [ ] Parallel structure in lists

## Scope

This skill covers English-language prose only. It does not apply to:
- Go source code identifiers (follow Go conventions: `gofmt`, effective Go)
- Commit messages (follow Conventional Commits via CLAUDE.md)
- YAML/JSON structure or formatting
- SQL, shell commands, or other non-prose code

When this skill conflicts with explicit project conventions in CLAUDE.md, the project convention wins.

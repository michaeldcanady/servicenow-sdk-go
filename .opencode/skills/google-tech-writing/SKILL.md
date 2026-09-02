---
name: google-tech-writing
description: >-
  Applies the Google developer documentation style guide
  (developers.google.com/style) to everything this repo's agents write and say:
  chat replies, code comments, commit messages, PR descriptions, issue write-ups,
  specs, ADRs, docs, and the skill files themselves. Active voice; second person
  ("you"); present tense; conditions before instructions; short sentences;
  standard American spelling and punctuation; sentence-case headings; serial
  commas; descriptive link text; no jargon, buzzwords, placeholder phrases (such
  as "please note" and "at this time"), exclamation points, "let's",
  "simply"/"easy", anthropomorphism, or Latin abbreviations ("e.g.", "i.e.",
  "etc."). USE when drafting, editing, or reviewing any prose the repo will
  consume, including replies to the user. USE proactively to self-review prose
  before sending. Not for Go code logic or test tables, which have their own
  conventions.
---

# Google technical writing

This skill applies the [Google developer documentation style guide](https://developers.google.com/style)
to every word this repo's agents produce: what you say to the user, what you
write in code comments, commit messages, PR descriptions, and issues, and what
you put in docs, specs, ADRs, and skill files. The goal is writing that is clear,
direct, and useful to a reader who may be in a hurry and may read English as a
second language. When a rule needs depth, follow the linked guide page — this
skill is the distilled version.

## Voice and tone

- **Be conversational, not formal, and not frivolous.** Aim for a knowledgeable,
  friendly tone: "This API lets you collect data about what your users like."
  ([style/tone](https://developers.google.com/style/tone))
- **Use second person.** Say "you", not "we". The SDK's authors are not the
  reader of a reply. Use "we" only to mean the repo team (for example, in an
  ADR's Consequences).
- **State the action directly.** Say what the reader does or what the code does.
  Do not add "please" to instructions: "To view the document, click View", not
  "please click View".
- **Avoid exclamation points.** Express emphasis with word choice, not
  punctuation.

## Sentences and grammar

- **Prefer active voice.** Make the doer the subject: "The server sends an
  acknowledgment", not "An acknowledgment is sent by the server."
  ([style/voice](https://developers.google.com/style/voice)) Passive voice is
  fine when the actor is irrelevant or you must de-emphasize the subject:
  "The file is saved"; "Over 50 conflicts were found in the file."
- **Write in present tense.** Say "The client retries three times", not "The
  client will retry three times". Reserve "will" for events that happen after
  publication.
- **Put conditions before actions.** "If the field is empty, the server omits
  it", not "The server omits the field if it is empty."
  ([style/sentence-structure](https://developers.google.com/style/sentence-structure))
- **Keep sentences short.** One idea per sentence. Cut filler ("there is",
  "it is worth noting", "in order to").
- **Put the context before the instruction.** Start a sentence with something
  the reader already knows, then deliver the new information.
- **Vary sentence openings.** Do not start every sentence with the same phrase,
  such as "You can" or "To do".
- **Show, don't stack.** Prefer a concrete example or a code snippet over a
  chain of abstractions.
- **Do not anthropomorphize.** Products do not "want", "decide", or "know".
  ([style/anthropomorphism](https://developers.google.com/style/anthropomorphism))

## Word choice

- **Use plain words.** Avoid jargon, buzzwords, clichés, and slang. If a term is
  unavoidable, define it at first use.
- **Spell out abbreviations.** Give the full term at first mention, followed by
  the abbreviation in parentheses. Avoid Latin abbreviations: say "for example",
  "that is", and "and so on" — not "e.g.", "i.e.", or "etc.".
  ([style/abbreviations](https://developers.google.com/style/abbreviations))
- **Avoid undermined claims.** Cut "simply", "just", "easy", "quickly", and
  similar words that tell the reader something is trivial when it may not be.
- **Cut placeholder phrases.** Remove "please note", "at this time", "as you
  know", and "it's important to note". If a note matters, state it directly.
- **Use standard American spelling and punctuation.**
- **Be consistent with terminology.** Use one word for one concept. Do not
  alternate "request", "call", and "invoke" for the same action.
  ([style/word-list](https://developers.google.com/style/word-list))

## Formatting and organization

- **Use sentence case** for headings and titles, not Title Case and not ALL
  CAPS.
- **Use a numbered list for a sequence** (steps the reader performs in order).
  Use bullets for any other list. Use a description list for term-and-definition
  pairs. Introduce each list with a sentence that says what it covers.
  ([style/lists](https://developers.google.com/style/lists))
- **Use serial commas.** "The SDK exposes Get, Post, and Delete."
- **Use descriptive link text.** Say "see `core.PageIterator`", not "click
  here".
- **Format code-related text as code** and UI elements in **bold**.
- **Use unambiguous dates** (for example, "Feb 5, 2026", not "02/05/26").
- **Put an example after the rule it illustrates.** A short snippet is worth a
  paragraph of abstract description.
- **Keep paragraphs to a few sentences.** One paragraph, one topic.

## Where this applies, with examples

| Surface | Do | Don't |
|---|---|---|
| Reply to the user | "Done. PR #762 is open." | "I have gone ahead and opened PR #762 for you!" |
| Commit message | "Add pagination to tableapi" | "Table API now supports page iteration, fixed edge cases, please review" |
| PR description | "Rewrites the skill from a catalog into a rulebook." | "This PR addresses the issue where the skill was basically just a catalog of the existing ADRs, which was, like, not triggering when needed." |
| Issue write-up | "When `limit` is zero, the request omits the parameter." | "The requests were failing because of an issue with the limit parameter." |
| Code comment | "Retry up to 3 times. The server returns 429 on throttle." | "Here we are just retrying the request repeatedly." |
| Skill / docs | "Use active voice. Clear prose survives translation." | "One of the many things that we should all be doing is using the active voice in our writings." |

## Self-review before you send

Before you finish any reply, commit message, PR body, issue, doc edit, or skill
edit, check it against this list:

- [ ] Did I use the active voice where it matters?
- [ ] Did I address the reader as "you"?
- [ ] Did I stay in the present tense?
- [ ] Are my sentences short enough to read out loud?
- [ ] Did I remove "please", "simply", "just", "note that", Latin
      abbreviations, and "will" except where a post-publication event makes it
      necessary?
- [ ] Did I use serial commas and sentence case?
- [ ] Did I avoid exclamation points and anthropomorphism?
- [ ] Is the link text descriptive?

If any answer is no, fix it before sending.

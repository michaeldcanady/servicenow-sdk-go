name: "Feature Request"
description: Request a new capability or enhancement (e.g. support for a new ServiceNow API, auth method, or SDK ergonomics improvement).
title: "[Story]: "
labels: ["type: feature", "state: new"]
body:
  - type: markdown
    attributes:
      value: |
        Thanks for taking the time to request a feature! Please fill this out as completely as you can —
        the more concrete the "so that" clause, the faster this can be scoped and prioritized.

  - type: textarea
    id: user-story
    attributes:
      label: User Story
      description: State this as "As a <persona>, I want <capability>, so that <outcome>."
      placeholder: |
        As a Go developer integrating with ServiceNow's Import Set API,
        I want to insert records into a staging table,
        so that I can trigger transform maps without hand-building HTTP calls.
    validations:
      required: true

  - type: textarea
    id: problem
    attributes:
      label: What problem does this solve?
      description: What can't you do today? What's the workaround, if any?
    validations:
      required: true

  - type: textarea
    id: proposed-solution
    attributes:
      label: Proposed Solution (optional)
      description: If you have an idea of the API surface, request builder shape, or config style, sketch it here.
    validations:
      required: false

  - type: textarea
    id: alternatives
    attributes:
      label: Alternatives Considered (optional)
      description: Other approaches you've tried or considered, and why they fell short.
    validations:
      required: false

  - type: dropdown
    id: module
    attributes:
      label: Affected Module / API
      description: Which part of the SDK does this touch?
      options:
        - table-api
        - attachment-api
        - batch-api
        - cdm (applications/policies/changesets/editor)
        - cmdb
        - case-api
        - authentication / credentials
        - core / request pipeline
        - documentation
        - CI/CD
        - other / not sure
    validations:
      required: true

  - type: input
    id: version
    attributes:
      label: Target Version (optional)
      description: If you have a preference for which SDK version should include this, note it here.
    validations:
      required: false

  - type: textarea
    id: additional-context
    attributes:
      label: Additional Context
      description: Links to related issues/PRs, relevant ServiceNow docs, or anything else useful.
    validations:
      required: false

  - type: checkboxes
    id: checklist
    attributes:
      label: Before submitting
      options:
        - label: I searched existing issues for a duplicate of this request.
          required: true
        - label: I checked the latest release to confirm this doesn't already exist.
          required: true
        - label: I've given this a clear, descriptive title.
          required: true

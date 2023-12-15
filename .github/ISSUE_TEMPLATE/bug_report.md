name: Bug report
description: Create a report to help us improve
title: "[BUG]: "
labels: ["bug"]
assignees: []

body:
- type: markdown
  attributes:
    value: |
      Thank you for reporting a bug in the servicenow-sdk-go module. Please fill out the following information to help us resolve the issue.

- type: input
  id: golang-version
  attributes:
    label: Golang version
    description: What version of Golang are you using?
    placeholder: e.g. 1.17.2
  validations:
    required: true

- type: input
  id: os-name-version
  attributes:
    label: OS name and version
    description: What operating system and version are you running?
    placeholder: e.g. Windows 10, Ubuntu 20.04
  validations:
    required: true

- type: input
  id: module-version
  attributes:
    label: Module version
    description: What version of the servicenow-sdk-go module are you using?
    placeholder: e.g. v0.0.1
  validations:
    required: true

- type: textarea
  id: bug-description
  attributes:
    label: Bug description
    description: Please provide a clear and concise description of the bug.
    placeholder: e.g. The module fails to connect to the ServiceNow instance with an authentication error.
  validations:
    required: true

- type: textarea
  id: minimum-code
  attributes:
    label: Minimum code required to reproduce
    description: Please provide the minimum code required to reproduce the bug.
    placeholder: |
      e.g.
      ```go
      package main

      import (
        "fmt"
        "github.com/michaeldcanady/servicenow-sdk-go"
      )

      func main() {
        client := servicenow.NewClient("https://example.service-now.com", "username", "password")
        record, err := client.GetRecord("incident", "sys_id")
        if err != nil {
          fmt.Println(err)
        } else {
          fmt.Println(record)
        }
      }
      ```
  validations:
    required: true

- type: textarea
  id: expected-behavior
  attributes:
    label: Expected behavior
    description: Please provide a clear and concise description of what you expected to happen.
    placeholder: e.g. The module should return a record object from the ServiceNow instance.
  validations:
    required: true

- type: textarea
  id: actual-behavior
  attributes:
    label: Actual behavior
    description: Please provide a clear and concise description of what actually happened.
    placeholder: e.g. The module returns an error message saying "invalid credentials".
  validations:
    required: true

- type: textarea
  id: screenshots
  attributes:
    label: Screenshots
    description: If applicable, add screenshots to help explain your problem.
  validations:
    required: false

- type: textarea
  id: additional-context
  attributes:
    label: Additional context
    description: Add any other context about the problem here.
  validations:
    required: false

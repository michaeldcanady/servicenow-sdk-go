---
name: unit-test-writer
description: Generates idiomatic Go unit tests using the table-driven pattern and the testify assertion library. Use when adding new features, fixing bugs, or improving test coverage to ensure robust and maintainable code.
---

# Unit Test Writer

## Overview
This skill provides a standardized workflow for writing Go unit tests that are consistent with the `servicenow-sdk-go` project's standards. It emphasizes table-driven tests for readability and `stretchr/testify` for expressive assertions.

## Workflow

1. **Analyze the Code**: Identify the function or method to be tested, its inputs, outputs, and dependencies.
2. **Identify Test Cases**:
    - **Happy Path**: The most common and successful execution path.
    - **Edge Cases**: Empty strings, nil pointers, zero values, maximum/minimum values, boundary conditions.
    - **Error Conditions**: Invalid inputs, dependency failures, context cancellation.
3. **Draft the Test Table**: Define a struct that captures the inputs and expected outputs for each test case.
4. **Implement the Test Function**: Use `t.Run` to iterate through the table and execute each case.
5. **Verify with Assertions**: Use `assert` (for non-fatal failures) or `require` (for fatal failures) from `github.com/stretchr/testify`.

## Implementation Patterns

### Table-Driven Template
```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name    string
        input   string // Adjust type as needed
        want    string // Adjust type as needed
        wantErr bool
        errMsg  string // Optional: for checking specific error messages
    }{
        {
            name:    "happy path",
            input:   "valid",
            want:    "expected",
            wantErr: false,
        },
        {
            name:    "empty input",
            input:   "",
            wantErr: true,
            errMsg:  "input is required",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := FunctionName(tt.input)
            if tt.wantErr {
                assert.Error(t, err)
                if tt.errMsg != "" {
                    assert.Contains(t, err.Error(), tt.errMsg)
                }
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tt.want, got)
            }
        })
    }
}
```

## Guidelines
- **Naming**: Use clear and descriptive names for both the test function and individual test cases.
- **Parallelism**: Use `t.Parallel()` in both the main test function and the subtests when appropriate.
- **Mocking**: For dependencies, use interfaces and generate mocks if necessary (check for existing mocks in the project).
- **Conciseness**: Keep test data focused on the scenario being tested.
- **Assertions**: Prefer `assert.Equal`, `assert.NoError`, `assert.Error`, and `assert.Contains`.

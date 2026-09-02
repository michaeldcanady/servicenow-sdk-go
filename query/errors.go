// Copyright (c) 2026 Michael Canady
// SPDX-License-Identifier: MIT

package query

import "errors"

// Sentinels reported by condition combination. Both fail closed: the affected
// condition renders as invalidQueryPlaceholder rather than panicking or
// emitting a partial query. Match with errors.Is.
var (
	// ErrNoConditions is reported when And or Or is called without any
	// conditions to combine.
	ErrNoConditions = errors.New("at least one condition must be provided")

	// ErrNilCondition is reported when combining against a nil condition,
	// whether as an argument to [Condition.And]/[Condition.Or] or as an
	// element of the variadic package-level combinators.
	ErrNilCondition = errors.New("condition cannot be nil")
)

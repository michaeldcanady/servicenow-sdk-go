// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package conversion

type Mutator[T, S any] func(input T) (S, error)

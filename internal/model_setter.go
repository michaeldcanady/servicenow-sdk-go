// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package internal

type ModelSetter[T any] func(val T) error

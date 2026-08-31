// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package internal

type ErrorThrower interface {
	Throw(typeName string, statusCode int64, contentType string, content []byte) error
}

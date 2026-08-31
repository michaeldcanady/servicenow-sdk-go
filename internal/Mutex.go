// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package internal

type Mutex interface {
	Lock()
	TryLock() bool
	Unlock()
}

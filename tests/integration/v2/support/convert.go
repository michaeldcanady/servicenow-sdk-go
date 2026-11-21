// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package support

import (
	"fmt"
	"math"
)

// ToInt32 converts n to an int32. It returns an error if n does not fit in an
// int32, so a step fails with a clear message instead of silently truncating.
func ToInt32(n int) (int32, error) {
	if n < math.MinInt32 || n > math.MaxInt32 {
		return 0, fmt.Errorf("value %d does not fit in an int32", n)
	}
	return int32(n), nil
}

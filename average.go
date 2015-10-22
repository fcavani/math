// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:        2014-04-25
// Last modification: 2014-

package math

import (
	"time"
)

// AvgInt64 calculates the average of the values in the slice.
func AvgInt64(values []int64) float64 {
	var sum int64
	for _, val := range values {
		sum += val
	}
	return float64(sum) / float64(len(values))
}

func AvgDuration(values []time.Duration) float64 {
	var sum time.Duration
	for _, val := range values {
		sum += val
	}
	return float64(sum) / float64(len(values))
}

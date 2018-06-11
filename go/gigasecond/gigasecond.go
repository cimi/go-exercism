// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond returns a moment in time a gigasecond after the input Time.
func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Duration(int(math.Pow10(9))) * time.Second
	return t.Add(gigasecond)
}

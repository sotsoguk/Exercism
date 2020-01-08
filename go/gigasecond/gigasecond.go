// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds 10^9 seconds (aka "gigasecond" to the given timestamp in seconds)
func AddGigasecond(t time.Time) time.Time {
	// LEARNED: exponential values can be directly written go
	return t.Add(time.Second * 1e9)
}

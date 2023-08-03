package main

import (
	"fmt"
	"time"

	"github.com/prep/average"
)

func main() {
	// Create a SlidingWindow that has a window of 15 minutes, with a
	// granulity of 1 minute.
	sw := average.MustNew(15*time.Minute, time.Minute)
	defer sw.Stop()

	// Do some work.
	sw.Add(15)
	// Do some more work.
	sw.Add(22)
	// Do even more work.
	sw.Add(22)

	fmt.Printf("Average of last  1m: %f\n", sw.Average(time.Minute))
	fmt.Printf("Average of last  5m: %f\n", sw.Average(5*time.Minute))
	fmt.Printf("Average of last 15m: %f\n\n", sw.Average(15*time.Minute))

	total, numSamples := sw.Total(15 * time.Minute)
	fmt.Printf("Counter has a total of %d over %d samples", total, numSamples)
}

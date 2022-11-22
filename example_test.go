package longduration_test

import (
	"fmt"

	"github.com/mashiike/longduration"
)

func ExampleParseDuration() {
	days, _ := longduration.ParseDuration("3d")
	hours, _ := longduration.ParseDuration("25h")
	complex, _ := longduration.ParseDuration("1d2h34m10s")
	fmt.Println(longduration.DurationString(days))
	fmt.Println(longduration.DurationString(hours))
	fmt.Println(longduration.DurationString(complex))
	fmt.Printf("There are %.0f hours in %s.\n", days.Hours(), longduration.DurationString(days))
	fmt.Printf("There are %.0f hours in %s.\n", hours.Hours(), longduration.DurationString(hours))
	fmt.Printf("There are %.0f seconds in %s.\n", complex.Seconds(), longduration.DurationString(complex))
	// Output:
	// 3d
	// 1d1h
	// 1d2h34m10s
	// There are 72 hours in 3d.
	// There are 25 hours in 1d1h.
	// There are 95650 seconds in 1d2h34m10s.
}

package longduration_test

import (
	"fmt"

	"github.com/mashiike/longduration"
)

func ExampleParseDuration() {
	years, _ := longduration.ParseDuration("1y")
	weeks, _ := longduration.ParseDuration("4w")
	days, _ := longduration.ParseDuration("3d")
	hours, _ := longduration.ParseDuration("25h")
	complex, _ := longduration.ParseDuration("1d2h34m10s")
	fmt.Println(longduration.DurationString(years))
	fmt.Println(longduration.DurationString(weeks, longduration.WithWeek()))
	fmt.Println(longduration.DurationString(days))
	fmt.Println(longduration.DurationString(hours))
	fmt.Println(longduration.DurationString(complex))
	fmt.Printf("There are %.4f weeks in %s.\n", longduration.Weeks(years), longduration.DurationString(years, longduration.WithYear()))
	fmt.Printf("There are %.0f days in %s.\n", longduration.Days(weeks), longduration.DurationString(weeks, longduration.WithWeek()))
	fmt.Printf("There are %.0f hours in %s.\n", days.Hours(), longduration.DurationString(days))
	fmt.Printf("There are %.0f hours in %s.\n", hours.Hours(), longduration.DurationString(hours))
	fmt.Printf("There are %.0f seconds in %s.\n", complex.Seconds(), longduration.DurationString(complex))
	// Output:
	// 365d
	// 4w
	// 3d
	// 1d1h
	// 1d2h34m10s
	// There are 52.1429 weeks in 1y.
	// There are 28 days in 4w.
	// There are 72 hours in 3d.
	// There are 25 hours in 1d1h.
	// There are 95650 seconds in 1d2h34m10s.
}

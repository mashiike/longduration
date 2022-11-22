# longduration

[![Documentation](https://godoc.org/github.com/mashiike/longduration?status.svg)](https://godoc.org/github.com/mashiike/longduration)
![Latest GitHub tag](https://img.shields.io/github/tag/mashiike/longduration.svg)
![Github Actions test](https://github.com/mashiike/longduration/workflows/Test/badge.svg?branch=main)
[![Go Report Card](https://goreportcard.com/badge/mashiike/longduration)](https://goreportcard.com/report/mashiike/longduration)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mashiike/longduration/blob/master/LICENSE)

 go utility package for long duration

-----

The standard go package does not have a duration expression beyond hour due to concerns about daylight saving time, leap seconds, etc.
However, in actual use, I would like to see notations such as `1d`, `3d`, `23d` etc.... Therefore, we have decided to use the following expression, which does not take into account daylight saving time, leap seconds, etc.
1d = 24hour, which does not take into account daylight saving time, leap seconds, and so on.

**Note: Do not use this package if you are affected by daylight saving time zones!** 

## Requirements
  * Go 1.16 or higher. support the 3 latest versions of Go.


See [godoc.org/github.com/mashiike/longduration](https://godoc.org/github.com/mashiike/longduration).

-----

## Installation

```shell
$ go get -u github.com/mashiike/longduration
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/mashiike/longduration"
)

func main() {
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
```

## LICENSE

MIT

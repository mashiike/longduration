package longduration

import (
	"fmt"
	"strings"
	"time"
)

type DuraitonStringOptions struct {
	ShowWeek bool
	ShowYear bool
}

// DurationString returns a string representation of the string Duration that also takes into account the number of days.
func DurationString(d time.Duration, optFns ...func(*DuraitonStringOptions)) string {
	opts := &DuraitonStringOptions{
		ShowYear: false,
		ShowWeek: false,
	}
	for _, optFn := range optFns {
		optFn(opts)
	}

	var builder strings.Builder
	remining := d
	if remining.Seconds() < 0 {
		fmt.Fprint(&builder, "-")
		remining = remining.Abs()
	}
	var years, weeks uint64
	if opts.ShowYear {
		years = uint64(Years(remining.Truncate(Year)))
		remining = remining - remining.Truncate(Year)
		if years > 0 {
			fmt.Fprintf(&builder, "%dy", years)
		}
		if remining == 0 {
			return builder.String()
		}
	}
	if opts.ShowWeek {
		weeks = uint64(Weeks(remining.Truncate(Week)))
		remining = remining - remining.Truncate(Week)
		if years > 0 || weeks > 0 {
			fmt.Fprintf(&builder, "%dw", weeks)
		}
		if remining == 0 {
			return builder.String()
		}
	}
	days := uint64(Days(remining.Truncate(Day)))
	remining = remining - remining.Truncate(Day)
	if years > 0 || weeks > 0 || days > 0 {
		fmt.Fprintf(&builder, "%dd", days)
	}
	if remining == 0 {
		return builder.String()
	}
	hours := uint64(remining.Truncate(time.Hour).Hours())
	remining = remining - remining.Truncate(time.Hour)
	if years > 0 || weeks > 0 || days > 0 || hours > 0 {
		fmt.Fprintf(&builder, "%dh", hours)
	}
	if remining == 0 {
		return builder.String()
	}
	minutes := uint64(remining.Truncate(time.Minute).Minutes())
	remining = remining - remining.Truncate(time.Minute)
	if years > 0 || weeks > 0 || days > 0 || hours > 0 || minutes > 0 {
		fmt.Fprintf(&builder, "%dm", minutes)
	}
	if remining == 0 {
		return builder.String()
	}
	fmt.Fprint(&builder, remining.String())
	return builder.String()
}

// WithYear is an option to the DurationString function that also displays the Year.
func WithYear() func(*DuraitonStringOptions) {
	return func(opts *DuraitonStringOptions) {
		opts.ShowYear = true
	}
}

// WithWeek is an option to the DurationString function that also displays the week.
func WithWeek() func(*DuraitonStringOptions) {
	return func(opts *DuraitonStringOptions) {
		opts.ShowWeek = true
	}
}

// Years calculates the number of years to be rounded down
func Years(d time.Duration) float64 {
	years := d / Year
	remining := (d - d.Truncate(Year)).Hours() / (float64(Year) / float64(time.Hour))
	return float64(years) + float64(remining)
}

// Weeks calculates the number of weeks to be rounded down
func Weeks(d time.Duration) float64 {
	weeks := d / Week
	remining := (d - d.Truncate(Week)).Hours() / (float64(Week) / float64(time.Hour))
	return float64(weeks) + float64(remining)
}

// Days calculates the number of days to be rounded down
func Days(d time.Duration) float64 {
	days := d / Day
	remining := (d - d.Truncate(Day)).Hours() / (float64(Day) / float64(time.Hour))
	return float64(days) + float64(remining)
}

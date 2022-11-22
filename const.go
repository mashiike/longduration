package longduration

import "time"

// Duration in days or longer units.
// Ignore the disruption caused by daylight saving time zone transitions.
// Don't use this package if daylight saving time zones are affected.
const (
	Day  = 24 * time.Hour
	Week = 7 * Day
	Year = 365 * Day
)

package longduration

import (
	"errors"
	"time"
)

// ParseDuration parses a duration string; it differs from the standard golang package in that it allows date expressions that do not take into account daylight saving time, etc. For example, `1d` is treated as 24 hours.
func ParseDuration(str string) (time.Duration, error) {
	years, str := trim(str, 'y', 365*24)
	if str == "" {
		if years == 0 {
			return 0, errors.New("invalid format")
		}
		return years, nil
	}
	weeks, str := trim(str, 'w', 7*24)
	if str == "" {
		if years+weeks == 0 {
			return 0, errors.New("invalid format")
		}
		return years + weeks, nil
	}
	days, str := trim(str, 'd', 24)
	if str == "" {
		if years+weeks+days == 0 {
			return 0, errors.New("invalid format")
		}
		return years + weeks + days, nil
	}
	d, err := time.ParseDuration(str)
	return years + weeks + days + d, err

}

func trim(str string, unit rune, coefficient int64) (time.Duration, string) {
	var val int64
	for i, c := range str {
		if '0' <= c && c <= '9' {
			v := int64(c - '0')
			val = val*10 + v
			continue
		}
		if c == unit {
			return time.Duration(val*coefficient) * time.Hour, str[i+1:]
		}
		return 0, str
	}
	return 0, str
}

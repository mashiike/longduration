package longduration

import (
	"errors"
	"time"
)

// ParseDuration parses a duration string; it differs from the standard golang package in that it allows date expressions that do not take into account daylight saving time, etc. For example, `1d` is treated as 24 hours.
func ParseDuration(str string) (time.Duration, error) {
	week, str := trimWeek(str)
	if str == "" {
		return week, nil
	}
	days, parts := trimDay(str)
	if parts != "" {
		d, err := time.ParseDuration(parts)
		return week + days + d, err
	}
	if days == 0 {
		return 0, errors.New("invalid format")
	}
	return days, nil
}

func trimDay(str string) (time.Duration, string) {
	var val int64
	for i, c := range str {
		if '0' <= c && c <= '9' {
			v := int64(c - '0')
			val = val*10 + v
			continue
		}
		if c == 'd' {
			return time.Duration(val) * 24 * time.Hour, str[i+1:]
		}
		return 0, str
	}
	return 0, str
}

func trimWeek(str string) (time.Duration, string) {
	var val int64
	for i, c := range str {
		if '0' <= c && c <= '9' {
			v := int64(c - '0')
			val = val*10 + v
			continue
		}
		if c == 'w' {
			return time.Duration(val) * 7 * 24 * time.Hour, str[i+1:]
		}
		return 0, str
	}
	return 0, str
}

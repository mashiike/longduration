package longduration_test

import (
	"testing"
	"time"

	"github.com/mashiike/longduration"
	"github.com/stretchr/testify/require"
)

func TestParseDurationSuccess(t *testing.T) {
	cases := []struct {
		str      string
		expected time.Duration
	}{
		{
			str:      "1m",
			expected: time.Minute,
		},
		{
			str:      "1h1m",
			expected: time.Hour + time.Minute,
		},
		{
			str:      "1d",
			expected: 24 * time.Hour,
		},
		{
			str:      "1w",
			expected: 7 * 24 * time.Hour,
		},
		{
			str:      "1d1m3s",
			expected: 24*time.Hour + time.Minute + 3*time.Second,
		},
		{
			str:      "1w1d1m3s",
			expected: 8*24*time.Hour + time.Minute + 3*time.Second,
		},
	}

	for _, c := range cases {
		t.Run(c.str, func(t *testing.T) {
			actual, err := longduration.ParseDuration(c.str)
			require.NoError(t, err)
			require.EqualValues(t, c.expected, actual)
		})
	}
}

func TestParseDurationFailed(t *testing.T) {
	cases := []struct {
		str string
	}{
		{
			str: "1",
		},
		{
			str: "s",
		},
		{
			str: "1mins",
		},
		{
			str: "1h1m1d",
		},
		{
			str: "1days",
		},
		{
			str: "11m1d3s",
		},
		{
			str: "d",
		},
	}

	for _, c := range cases {
		t.Run(c.str, func(t *testing.T) {
			_, err := longduration.ParseDuration(c.str)
			require.Error(t, err)
		})
	}
}

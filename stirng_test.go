package longduration_test

import (
	"testing"
	"time"

	"github.com/mashiike/longduration"
	"github.com/stretchr/testify/require"
)

func TestDurationString(t *testing.T) {
	cases := []struct {
		d        time.Duration
		optFns   []func(*longduration.DuraitonStringOptions)
		expected string
	}{
		{
			expected: "-1m",
			d:        -time.Minute,
		},
		{
			expected: "-1h",
			d:        -time.Hour,
		},
		{
			expected: "-1d",
			d:        -24 * time.Hour,
		},
		{
			expected: "-7d",
			d:        -7 * 24 * time.Hour,
		},
		{
			expected: "-7d0h1m3s",
			d:        -1 * (7*24*time.Hour + time.Minute + 3*time.Second),
		},
		{
			expected: "-7d0h1m",
			d:        -1 * (7*24*time.Hour + time.Minute),
		},
		{
			expected: "-1w0d0h1m",
			d:        -1 * (7*24*time.Hour + time.Minute),
			optFns: []func(*longduration.DuraitonStringOptions){
				func(opts *longduration.DuraitonStringOptions) {
					opts.ShowWeek = true
				},
			},
		},
		{
			expected: "-1y",
			d:        -1 * (365 * 24 * time.Hour),
			optFns: []func(*longduration.DuraitonStringOptions){
				func(opts *longduration.DuraitonStringOptions) {
					opts.ShowYear = true
				},
			},
		},
		{
			expected: "1m",
			d:        time.Minute,
		},
		{
			expected: "1h1m",
			d:        time.Hour + time.Minute,
		},
		{
			expected: "1d",
			d:        24 * time.Hour,
		},
		{
			expected: "7d",
			d:        7 * 24 * time.Hour,
		},
		{
			expected: "7d0h1m3s",
			d:        7*24*time.Hour + time.Minute + 3*time.Second,
		},
		{
			expected: "1y",
			d:        (365 * 24 * time.Hour),
			optFns: []func(*longduration.DuraitonStringOptions){
				func(opts *longduration.DuraitonStringOptions) {
					opts.ShowYear = true
				},
			},
		},
		{
			expected: "1y3d",
			d:        (365*24*time.Hour + 3*24*time.Hour),
			optFns: []func(*longduration.DuraitonStringOptions){
				func(opts *longduration.DuraitonStringOptions) {
					opts.ShowYear = true
				},
			},
		},
		{
			expected: "1y0w3d",
			d:        (365*24*time.Hour + 3*24*time.Hour),
			optFns: []func(*longduration.DuraitonStringOptions){
				func(opts *longduration.DuraitonStringOptions) {
					opts.ShowYear = true
					opts.ShowWeek = true
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.expected, func(t *testing.T) {
			actual := longduration.DurationString(c.d, c.optFns...)
			require.EqualValues(t, c.expected, actual)
		})
	}
}

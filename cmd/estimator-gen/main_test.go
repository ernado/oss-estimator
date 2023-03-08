package main

import (
	"testing"
)

func TestHumanize(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{1500, "1.5 K"},
		{13410, "13.41 K"},
		{1345000, "1.34 M"},
		{500, "500"},
		{1000, "1 K"},
		{999, "999"},
		{999499, "999.5 K"},
		{999999, "1 M"},
		{1000001, "1 M"},
	}

	for _, tc := range testCases {
		actual := NearestThousandFormat(float64(tc.input))
		if actual != tc.expected {
			t.Errorf("humanize(%d): expected %s, but got %s", tc.input, tc.expected, actual)
		}
	}
}

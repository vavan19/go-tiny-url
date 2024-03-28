package utils

import (
	"testing"
)

func TestEncodeToBase62(t *testing.T) {
	testCases := []struct {
		id       int
		expected string
	}{
		{id: 1, expected: "0000001"},
		{id: 999, expected: "00000gV"},
		{id: 9999999, expected: "5YC1sJ"},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			result := EncodeToBase62(tc.id)
			if result != tc.expected {
				t.Errorf("EncodeToBase62(%d) = %s; want %s", tc.id, result, tc.expected)
			}
			if len(result) != 7 {
				t.Errorf("Expected result length of 7, got %d", len(result))
			}
		})
	}
}

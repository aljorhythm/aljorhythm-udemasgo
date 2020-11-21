package utility

import "testing"

// AssertIntEquals assert ints
func AssertIntEquals(t *testing.T, msg string, actual, expected int) {
	if actual == expected {
		t.Log(msg)
	} else {
		t.Errorf("FAILED: %s\nExpected %d, got %d", msg, expected, actual)
	}
}

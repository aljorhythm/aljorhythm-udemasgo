package utility

import "testing"

func TestAdd(t *testing.T) {
	res := Add(1, -1)
	if res != 0 {
		t.Errorf("Adding 1 and -1 should give 0, got %d", res)
	}
}

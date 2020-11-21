package main

import (
	"testing"

	"github.com/aljorhythm/aljorhythm-udemasgo/models"
)

func TestMain(t *testing.T) {
	t.Log("Put something to test")
}

func assertIntEquals(t *testing.T, msg string, actual, expected int) {
	if actual == expected {
		t.Log(msg)
	} else {
		t.Errorf("FAILED: %s\nExpected %d, got %d", msg, expected, actual)
	}
}

func TestSlices(t *testing.T) {
	t.Log("Slicings")

	s := []int{0, 2, 3, 4, 5}
	s1 := s[0:2]

	assertIntEquals(t, "s1 length is 2", len(s1), 2)

	assertIntEquals(t, "s[0] is 0", s[0], 0)
	assertIntEquals(t, "s1[0] is 0", s1[0], 0)

	s[0] = 1
	assertIntEquals(t, "s[0] is 1", s[0], 1)
	assertIntEquals(t, "s1[0] is 1", s1[0], 1)
}

func TestMaps(t *testing.T) {
	crewMembers := map[int]models.CrewMember{
		1: models.NewCrewMember(1, "John"),
		2: models.NewCrewMember(2, "Jane"),
	}

	assertIntEquals(t, "there are 2 crewMembers", len(crewMembers), 2)

	if _, found := crewMembers[0]; found {
		t.Error("Should not have found member")
	}
	if member, found := crewMembers[1]; found {
		t.Logf("Found member %s", member.GetName())
		if member != models.NewCrewMember(1, "John") {
			t.Error("Member struct not equal")
		} else {
			t.Logf("Member struct can be equaled")
		}
	}
}

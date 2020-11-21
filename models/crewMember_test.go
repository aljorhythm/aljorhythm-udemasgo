package models

import "testing"

func TestCrewMemberString(t *testing.T) {
	cmember := CrewMember{name: "John", id: 1}
	expected := "1 John"
	actual := cmember.ToString()
	if actual != expected {
		t.Errorf("Expected %s, got %s instead", expected, actual)
	}
}

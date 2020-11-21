package models

import "fmt"

// CrewMember crew member
type CrewMember struct {
	name string
	id   int
}

// ToString string representation
func (crewMember *CrewMember) ToString() string {
	return fmt.Sprintf("%d %s", crewMember.id, crewMember.name)
}

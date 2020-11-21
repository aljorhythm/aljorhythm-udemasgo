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

// GetName get name
func (crewMember *CrewMember) GetName() string {
	return crewMember.name
}

// GetID get id
func (crewMember *CrewMember) GetID() int {
	return crewMember.id
}

// NewCrewMember new crew member
func NewCrewMember(id int, name string) CrewMember {
	return CrewMember{id: id, name: name}
}

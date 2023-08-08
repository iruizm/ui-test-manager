package model

import "github.com/google/uuid"

type Sequence struct {
	Name       string
	Id         uuid.UUID
	Precedents []*Sequence
}

func NewSequence(name string) *Sequence {
	var a Sequence

	a.Id = uuid.New()
	a.Name = name
	a.Precedents = []*Sequence{}

	return &a

}

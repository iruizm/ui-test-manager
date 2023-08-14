package model

import "github.com/google/uuid"

type Sequence struct {
	Name       string
	Id         uuid.UUID
	Content    string
	Precedents []*Sequence
}

func NewSequence(name string, content string) *Sequence {
	var a Sequence

	a.Id = uuid.New()
	a.Name = name
	a.Content = content
	a.Precedents = []*Sequence{}

	return &a

}

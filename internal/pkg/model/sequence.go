package model

import (
	"github.com/google/uuid"
)

type Sequence struct {
	Name       string      `json:"name"`
	Id         uuid.UUID   `json:"id"`
	Content    string      `json:"content"`
	Precedents []*Sequence `json:"precedents"`
}

func NewSequence(name string, content string) *Sequence {
	var a Sequence

	a.Id = uuid.New()
	println("New ID:" + a.Id.String())
	a.Name = name
	a.Content = content
	a.Precedents = []*Sequence{}

	return &a
}

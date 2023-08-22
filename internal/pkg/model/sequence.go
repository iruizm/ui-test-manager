package model

import (
	"github.com/google/uuid"
)

type Sequence struct {
	Name       string       `json:"name"`
	Id         uuid.UUID    `json:"id"`
	Content    string       `json:"content"`
	Precedents []*uuid.UUID `json:"precedents"`
}

func NewSequence(name string, content string) *Sequence {
	var a Sequence

	a.Id = uuid.New()
	a.Name = name
	a.Content = content
	a.Precedents = []*uuid.UUID{}

	return &a
}

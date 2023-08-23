package model

import (
	"github.com/google/uuid"
)

type Sequence struct {
	Name       string      `json:"name"`
	Id         uuid.UUID   `json:"id"`
	Content    string      `json:"content"`
	Precedents []uuid.UUID `json:"precedents"`
}

func NewSequence(name string, content string) *Sequence {
	return &Sequence{
		Id:         uuid.New(),
		Name:       name,
		Content:    content,
		Precedents: make([]uuid.UUID, 0),
	}
}

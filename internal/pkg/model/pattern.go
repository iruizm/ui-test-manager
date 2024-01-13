package model

import (
	"github.com/google/uuid"
)

type Pattern struct {
	Name    string    `json:"name"`
	Id      uuid.UUID `json:"id"`
	Before  string    `json:"before"`
	After   string    `json:"after"`
	Replace bool      `json:"replace"`
}

func NewPattern() *Pattern {
	return &Pattern{
		Name:    "Rule",
		Id:      uuid.New(),
		Before:  "",
		After:   "",
		Replace: false,
	}
}

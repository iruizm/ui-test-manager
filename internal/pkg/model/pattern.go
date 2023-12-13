package model

import (
	"github.com/google/uuid"
)

type Pattern struct {
	Name        string    `json:"name"`
	Id          uuid.UUID `json:"id"`
	Regex       string    `json:"regex"`
	Replacement string    `json:"replacement"`
}

func NewPattern() *Pattern {
	return &Pattern{
		Name:        "",
		Id:          uuid.New(),
		Regex:       "",
		Replacement: "",
	}
}

package model

import "github.com/google/uuid"

type Sequence struct {
	FileName   string
	Id         uuid.UUID
	Precedents []*Sequence
}

func NewSequence(fileName string) *Sequence {
	var a Sequence

	a.Id = uuid.New()
	a.FileName = fileName
	a.Precedents = []*Sequence{}

	return &a

}

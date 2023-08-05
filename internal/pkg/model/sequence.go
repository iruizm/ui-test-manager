package model

type Sequence struct {
	FileName   string
	Id         string
	Precedents []*Sequence
}

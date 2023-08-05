package main

import (
	api "polonium/internal/app"
	"polonium/internal/pkg/model"
	"polonium/internal/pkg/persistence"
)

func main() {
	api.Run()

	var a model.Sequence

	a.Id = "HHHAHAHAH"
	a.FileName = "AAAAA"
	a.Precedents = make([]*model.Sequence, 3)

	persistence.GetSequenceMap().AddSequence(&a)
}

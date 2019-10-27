package mem_model

import ()

type MemModel struct {
	cards          []Card
	cardByL2       map[string]Card
	nextCardId     int
	morphemes      []Morpheme
	morphemeByL2   map[string]Morpheme
	nextMorphemeId int
}

func NewMemModel() *MemModel {
	return &MemModel{
		cards:          []Card{},
		cardByL2:       map[string]Card{},
		nextCardId:     1,
		morphemes:      []Morpheme{},
		morphemeByL2:   map[string]Morpheme{},
		nextMorphemeId: 1,
	}
}

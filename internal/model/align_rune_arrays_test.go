package model

import (
	"gotest.tools/assert"
	"testing"
)

func Test1(t *testing.T) {
	xRunes := []rune("a2 a3")
	yRunes := []rune("a3")

	alignments := AlignRuneArrays(xRunes, yRunes)

	assert.DeepEqual(t, alignments, []Alignment{
		{X: 0, Y: -1},
		{X: 1, Y: -1},
		{X: 2, Y: -1},
		{X: 3, Y: 0},
		{X: 4, Y: 1},
	})
}

func Test2(t *testing.T) {
	xRunes := []rune("duvirotewell")
	yRunes := []rune("dUW vIY rOW tIH w EH l?")

	alignments := AlignRuneArrays(xRunes, yRunes)

	printAlignments(xRunes, yRunes, alignments)

	assert.DeepEqual(t, 1, 0)
}

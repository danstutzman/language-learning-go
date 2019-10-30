package model

import (
	"gotest.tools/assert"
	"log"
	"testing"
)

func printAlignments(xRunes, yRunes []rune, alignments []Alignment) {
	for _, alignment := range alignments {
		xRune := ""
		if alignment.X != -1 {
			xRune = string(xRunes[alignment.X : alignment.X+1])
		}
		yRune := ""
		if alignment.Y != -1 {
			yRune = string(yRunes[alignment.Y : alignment.Y+1])
		}
		log.Printf("'%s' => '%s'", xRune, yRune)
	}
}

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

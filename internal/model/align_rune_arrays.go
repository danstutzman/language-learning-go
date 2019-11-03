package model

import (
	"log"
	"unicode"
)

const (
	STOP    = 0
	LEFT    = 1
	UP      = 2
	UP_LEFT = 3
)

type Alignment struct {
	X int
	Y int
}

var vowelRunes = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'y': true,
	'á': true,
	'é': true,
	'í': true,
	'ó': true,
	'ú': true,
	'ü': true,
}

var consonantRunes = map[rune]bool{
	'b': true,
	'c': true,
	'd': true,
	'f': true,
	'g': true,
	'h': true,
	'j': true,
	'k': true,
	'l': true,
	'm': true,
	'n': true,
	'p': true,
	'q': true,
	'r': true,
	's': true,
	't': true,
	'w': true,
	'x': true,
	'y': true,
	'z': true,
	'ñ': true,
}

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

// Needleman–Wunsch algorithm
func AlignRuneArrays(xRunes, yRunes []rune) []Alignment {
	maxScoreRows := [][]int{}
	arrowDirectionRows := [][]int{}

	for y := 0; y <= len(yRunes); y++ {
		maxScoreRow := []int{}
		arrowDirectionRow := []int{}

		for x := 0; x <= len(xRunes); x++ {
			var maxScore int
			var arrowDirection int

			if x == 0 && y == 0 {
				maxScore = 0
				arrowDirection = STOP
			} else if x == 0 {
				maxScore = -y
				arrowDirection = UP
			} else if y == 0 {
				maxScore = -x
				arrowDirection = LEFT
			} else {
				upLeft := maxScoreRows[y-1][x-1] +
					scoreRuneMatch(xRunes[x-1], yRunes[y-1])
				up := maxScoreRows[y-1][x] - 1
				left := maxScoreRow[x-1] - 1

				// Test upLeft first so we prioritize it
				if upLeft >= up && upLeft >= left {
					maxScore = upLeft
					arrowDirection = UP_LEFT
				} else if up >= upLeft && up >= left {
					maxScore = up
					arrowDirection = UP
				} else if left >= upLeft && left >= up {
					maxScore = left
					arrowDirection = LEFT
				} else {
					panic("Impossible to reach this line")
				}

			}
			maxScoreRow = append(maxScoreRow, maxScore)
			arrowDirectionRow = append(arrowDirectionRow, arrowDirection)
		}
		maxScoreRows = append(maxScoreRows, maxScoreRow)
		arrowDirectionRows = append(arrowDirectionRows, arrowDirectionRow)
	}

	alignments := []Alignment{}
	x := len(xRunes)
	y := len(yRunes)
loop:
	for x >= 0 && y >= 0 {
		switch arrowDirectionRows[y][x] {
		case LEFT:
			alignments = append(alignments, Alignment{X: x - 1, Y: -1})
			x -= 1
		case UP:
			alignments = append(alignments, Alignment{X: -1, Y: y - 1})
			y -= 1
		case UP_LEFT:
			alignments = append(alignments, Alignment{X: x - 1, Y: y - 1})
			x -= 1
			y -= 1
		default:
			break loop
		}
	}

	// Reverse alignments
	for i, j := 0, len(alignments)-1; i < j; i, j = i+1, j-1 {
		alignments[i], alignments[j] = alignments[j], alignments[i]
	}

	return alignments
}

func scoreRuneMatch(xRune, yRune rune) int {
	if xRune == yRune {
		return 4
	}

	xRuneLower := unicode.ToLower(xRune)
	yRuneLower := unicode.ToLower(yRune)
	if xRuneLower == yRuneLower {
		return 3
	}

	if vowelRunes[xRuneLower] && vowelRunes[yRuneLower] ||
		consonantRunes[xRuneLower] && consonantRunes[yRuneLower] {
		return 2
	}

	if (vowelRunes[xRuneLower] || consonantRunes[xRuneLower]) &&
		(vowelRunes[yRuneLower] || consonantRunes[yRuneLower]) {
		return 1
	}

	return -1
}

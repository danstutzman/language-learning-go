package model

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
		return 1
	} else {
		return -1
	}
}

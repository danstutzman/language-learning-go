package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type Alignment struct {
	L1 int
	L2 int
}

func convertAlignments(line string) []Alignment {
	alignments := []Alignment{}
	for _, pair := range strings.Split(line, " ") {
		values := strings.Split(pair, "-")
		l1 := mustAtoi(values[0])
		l2 := mustAtoi(values[1])
		alignments = append(alignments, Alignment{L1: l1, L2: l2})
	}
	return alignments
}

type Split struct {
	L1 int
	L2 int
}

type Span struct {
	L1Begin int
	L1End   int
	L2Begin int
	L2End   int
}

var SPLIT_ON_WHITESPACE = regexp.MustCompile(`\s+`)

func main() {
	if len(os.Args) != 2+1 {
		fmt.Fprintf(os.Stderr, `Usage:
			1st arg is path to fast_align binary
			2nd arg is path to parallel corpus text
		`)
		os.Exit(1)
	}
	fastAlignPath := os.Args[1]
	corpusPath := os.Args[2]

	cmd := exec.Command(fastAlignPath, "-d", "-o", "-i", corpusPath)

	alignmentsCombined, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	spansByLineNum := map[int][]Span{}
	for lineNum0, line := range strings.Split(string(alignmentsCombined), "\n") {
		if line != "" {
			alignments := convertAlignments(line)

			numL1Words := 0
			numL2Words := 0
			for _, alignment := range alignments {
				if alignment.L1 >= numL1Words {
					numL1Words = alignment.L1 + 1
				}
				if alignment.L2 >= numL2Words {
					numL2Words = alignment.L2 + 1
				}
			}

			spans := []Span{}
			for l2Split := 1; l2Split < numL2Words; l2Split++ {
				lastL1End := 1
				if len(spans) > 0 {
					lastL1End = spans[len(spans)-1].L1End
				}

				var nextSpan *Span
				for l1Split := lastL1End + 1; l1Split < numL1Words; l1Split++ {
					crosses := false
					for _, alignment := range alignments {
						if alignment.L1 < l1Split && alignment.L2 >= l2Split ||
							alignment.L1 >= l1Split && alignment.L2 < l2Split {
							crosses = true
							break
						}
					}

					if !crosses {
						if len(spans) > 0 {
							lastSpan := spans[len(spans)-1]
							nextSpan = &Span{
								L1Begin: lastSpan.L1End,
								L1End:   l1Split,
								L2Begin: lastSpan.L2End,
								L2End:   l2Split,
							}
						} else {
							nextSpan = &Span{
								L1Begin: 0,
								L1End:   l1Split,
								L2Begin: 0,
								L2End:   l2Split,
							}
						}
					}
				}
				if nextSpan != nil {
					spans = append(spans, *nextSpan)
				}
			}

			if len(spans) > 0 {
				lastSpan := spans[len(spans)-1]
				spans = append(spans, Span{
					L1Begin: lastSpan.L1End,
					L1End:   -1, // means last word
					L2Begin: lastSpan.L2End,
					L2End:   -1, // means last word
				})
			} else {
				spans = []Span{{
					L1Begin: 0,
					L1End:   -1, // means last word
					L2Begin: 0,
					L2End:   -1, // means last word
				}}
			}

			spansByLineNum[lineNum0+1] = spans
		}
	}

	file, err := os.Open(corpusPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNum += 1

		values := strings.Split(line, " ||| ")
		l1Words := SPLIT_ON_WHITESPACE.Split(strings.TrimSpace(values[0]), -1)
		l2Words := SPLIT_ON_WHITESPACE.Split(strings.TrimSpace(values[1]), -1)

		fmt.Printf("%s\n", line)
		for _, span := range spansByLineNum[lineNum] {
			l1End := span.L1End
			if l1End == -1 {
				l1End = len(l1Words)
			}

			l2End := span.L2End
			if l2End == -1 {
				l2End = len(l2Words)
			}

			fmt.Printf("%-35v %-35v\n",
				strings.Join(l1Words[span.L1Begin:l1End], " "),
				strings.Join(l2Words[span.L2Begin:l2End], " "))
		}
		fmt.Printf("\n")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

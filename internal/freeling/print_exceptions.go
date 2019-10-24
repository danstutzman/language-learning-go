package freeling

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Conjugation struct {
	stem   string
	suffix string
}

func PrintVerbExceptions(freelingDiccPath string) {
	file, err := os.Open(freelingDiccPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum += 1
		line := scanner.Text()

		if lineNum == 1 && line == "<IndexType>" ||
			lineNum == 2 && line == "DB_MAP" ||
			lineNum == 3 && line == "</IndexType>" ||
			lineNum == 4 && line == "<Entries>" ||
			line == "</Entries>" {
			continue
		}

		values := strings.Split(line, " ")
		form := values[0]
		for i := 1; i < len(values); i += 2 {
			lemma := values[i]
			tag := values[i+1]

			expected := true
			if false && strings.HasSuffix(lemma, "ar") && strings.HasPrefix(tag, "V") {
				conjugations := analyzeArVerb(lemma, tag)
				expected = false
				for _, conjugation := range conjugations {
					if conjugation.stem+conjugation.suffix == form {
						expected = true
					}
				}
			}
			if false && strings.HasSuffix(lemma, "er") && strings.HasPrefix(tag, "V") {
				conjugations := analyzeErVerb(lemma, tag)
				expected = false
				for _, conjugation := range conjugations {
					if conjugation.stem+conjugation.suffix == form {
						expected = true
					}
				}
				if false && !expected {
					for _, conjugation := range conjugations {
						fmt.Printf("%-20s %-20s %-10s %v\n", lemma, form, tag, conjugation)
					}
				}
			}
			if strings.HasSuffix(lemma, "ir") && strings.HasPrefix(tag, "V") {
				conjugations := analyzeIrVerb(lemma, tag)
				expected = false
				for _, conjugation := range conjugations {
					if conjugation.stem+conjugation.suffix == form {
						expected = true
					}
				}
			}

			if !expected {
				fmt.Printf("%-20s %-20s %-10s\n", lemma, form, tag)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func analyzeArVerb(lemma, tag string) []Conjugation {
	conjugations := []Conjugation{}
	suffixes := AR_MUTANT_TAG27_TO_SUFFIXES[tag[2:7]]
	if len(suffixes) > 0 && AR_MUTANTS[lemma] != "" {
		stem := AR_MUTANTS[lemma]

		for _, suffix := range suffixes {
			conjugation := Conjugation{stem: stem, suffix: suffix}
			conjugations = append(conjugations, conjugation)
		}
	} else {
		stems := AR_STEM_CHANGES[lemma]
		if len(stems) == 0 || !AR_TAG27_TO_STEM_CHANGE[tag[2:7]] {
			stems = []string{lemma[0 : len(lemma)-2]}
		}

		suffixes = AR_TAG27_TO_SUFFIXES[tag[2:7]]

		for _, stem := range stems {
			for _, suffix := range suffixes {
				if suffix == "e" || suffix == "en" || suffix == "é" ||
					suffix == "emos" || suffix == "es" || suffix == "éis" {
					stem = ENDS_WITH_C.ReplaceAllString(stem, "qu")
					stem = ENDS_WITH_GU.ReplaceAllString(stem, "gü")
					stem = ENDS_WITH_G.ReplaceAllString(stem, "gu")
					stem = ENDS_WITH_Z.ReplaceAllString(stem, "c")
				}

				if lemma == "estar" {
					if suffix == "a" {
						suffix = "á"
					} else if suffix == "an" {
						suffix = "án"
					} else if suffix == "as" {
						suffix = "ás"
					} else if suffix == "e" {
						suffix = "é"
					} else if suffix == "en" {
						suffix = "én"
					} else if suffix == "es" {
						suffix = "és"
					}
				}

				conjugation := Conjugation{stem: stem, suffix: suffix}
				conjugations = append(conjugations, conjugation)
			}
		}
	}
	return conjugations
}

func analyzeErVerb(lemma, tag string) []Conjugation {
	groupInfinitiveStems := groupInfinitiveStemsByInfinitive[lemma]

	groups := map[string]bool{}
	for _, groupInfinitiveStem := range groupInfinitiveStems {
		groups[groupInfinitiveStem.group] = true
	}

	defaultStem := lemma[0 : len(lemma)-2]
	if ENDS_WITH_A_E_OR_O.MatchString(defaultStem) {
		groups["ER_ENDS_WITH_A_E_OR_O"] = true
	}
	if strings.HasSuffix(defaultStem, "ll") ||
		strings.HasSuffix(defaultStem, "ñ") {
		groups["ER_ENDS_WITH_LL_OR_Ñ"] = true
	}
	if strings.HasSuffix(defaultStem, "n") {
		groups["ER_ENDS_WITH_N"] = true
	}
	groups["ER"] = true

	groupTag27Suffixes := findGroupTag27Suffixes(groups, tag[2:7])

	conjugations := []Conjugation{}
	for _, groupTag27Suffix := range groupTag27Suffixes {
		stemsForGroup := []string{}
		for _, groupInfinitiveStem := range groupInfinitiveStems {
			if groupInfinitiveStem.group == groupTag27Suffix.group {
				stemsForGroup = append(stemsForGroup, groupInfinitiveStem.stem)
			}
		}
		if len(stemsForGroup) == 0 {
			stemsForGroup = []string{defaultStem}
		}

		for _, stem := range stemsForGroup {
			conjugation := Conjugation{
				stem:   stem,
				suffix: groupTag27Suffix.suffix,
			}
			conjugations = append(conjugations, conjugation)
		}
	}

	return conjugations
}

func analyzeIrVerb(lemma, tag string) []Conjugation {
	groupInfinitiveStems := groupInfinitiveStemsByInfinitive[lemma]

	groups := map[string]bool{}
	for _, groupInfinitiveStem := range groupInfinitiveStems {
		groups[groupInfinitiveStem.group] = true
	}

	defaultStem := lemma[0 : len(lemma)-2]
	if ENDS_WITH_A_E_O_OR_U.MatchString(defaultStem) {
		groups["IR_ENDS_WITH_A_E_O_OR_U"] = true
	}
	if strings.HasSuffix(defaultStem, "ll") ||
		strings.HasSuffix(defaultStem, "ñ") ||
		strings.HasSuffix(defaultStem, "y") {
		groups["IR_ENDS_WITH_LL_Ñ_OR_Y"] = true
	}
	groups["IR"] = true

	groupTag27Suffixes := findGroupTag27Suffixes(groups, tag[2:7])

	conjugations := []Conjugation{}
	for _, groupTag27Suffix := range groupTag27Suffixes {
		stemsForGroup := []string{}
		for _, groupInfinitiveStem := range groupInfinitiveStems {
			if groupInfinitiveStem.group == groupTag27Suffix.group {
				stemsForGroup = append(stemsForGroup, groupInfinitiveStem.stem)
			}
		}
		if len(stemsForGroup) == 0 {
			stemsForGroup = []string{defaultStem}
		}

		for _, stem := range stemsForGroup {
			conjugation := Conjugation{
				stem:   stem,
				suffix: groupTag27Suffix.suffix,
			}
			conjugations = append(conjugations, conjugation)
		}
	}

	return conjugations
}

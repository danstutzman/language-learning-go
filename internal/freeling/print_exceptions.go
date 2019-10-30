package freeling

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Conjugation struct {
	Stem   string
	Suffix string
	Group  string
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
			if strings.HasPrefix(tag, "V") {
				conjugations := AnalyzeVerb(lemma, tag)
				expected = false
				for _, conjugation := range conjugations {
					if conjugation.Stem+conjugation.Suffix == form {
						expected = true
					}
				}
				if false && !expected {
					for _, conjugation := range conjugations {
						fmt.Printf("%-20s %-20s %-10s %s %s\n", lemma, form, tag, conjugation.Stem, conjugation.Suffix)
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

func AnalyzeVerb(lemma, tag string) []Conjugation {
	conjugations := []Conjugation{}

	uniqueVerbs := findUniqueVerbs(lemma, tag[2:7])
	if len(uniqueVerbs) > 0 {
		for _, uniqueVerb := range uniqueVerbs {
			conjugation := Conjugation{
				Stem:   uniqueVerb.form,
				Suffix: "",
				Group:  "UNIQUE",
			}
			conjugations = append(conjugations, conjugation)
		}
	}

	groupInfinitiveStems := groupInfinitiveStemsByInfinitive[lemma]

	groups := map[string]bool{}
	for _, groupInfinitiveStem := range groupInfinitiveStems {
		groups[groupInfinitiveStem.group] = true
	}

	defaultStem := ENDS_WITH_AR_ER_IR_OR_ÍR.ReplaceAllString(lemma, "")
	if strings.HasSuffix(lemma, "ar") {
		groups["AR"] = true
	} else if strings.HasSuffix(lemma, "er") {
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
	} else if (strings.HasSuffix(lemma, "ir") ||
		strings.HasSuffix(lemma, "ír")) && lemma != "ir" {
		if ENDS_WITH_A_E_O_OR_U.MatchString(defaultStem) {
			groups["IR_ENDS_WITH_A_E_O_OR_U"] = true
		}
		if strings.HasSuffix(defaultStem, "ll") ||
			strings.HasSuffix(defaultStem, "ñ") ||
			strings.HasSuffix(defaultStem, "y") {
			groups["IR_ENDS_WITH_LL_Ñ_OR_Y"] = true
		}
		groups["IR"] = true
	}

	groupTag27Suffixes := findGroupTag27Suffixes(groups, tag[2:7])

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
				Stem:   stem,
				Suffix: groupTag27Suffix.suffix,
				Group:  groupTag27Suffix.group,
			}
			conjugations = append(conjugations, conjugation)
		}
	}

	return conjugations
}

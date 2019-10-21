package freeling

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
			if strings.HasSuffix(lemma, "ar") && strings.HasPrefix(tag, "V") {
				expected = isExpectedArVerb(lemma, form, tag)
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

func isExpectedArVerb(lemma, form, tag string) bool {
	suffix := AR_MUTANT_TAG27_TO_SUFFIX[tag[2:7]]
	if suffix != "" && AR_MUTANTS[lemma] != "" {
		stem := AR_MUTANTS[lemma]

		return (form == stem+suffix) ||
			ENDS_WITH_ESE.ReplaceAllString(form, "era$1") == stem+suffix ||
			ENDS_WITH_ÉSEMOS.ReplaceAllString(form, "éramos") == stem+suffix
	} else {
		stem := AR_STEM_CHANGES[lemma]
		if stem == "" || !AR_TAG27_TO_STEM_CHANGE[tag[2:7]] {
			stem = lemma[0 : len(lemma)-2]
		}

		suffix = AR_TAG27_TO_SUFFIX[tag[2:7]]
		if suffix == "e" || suffix == "en" || suffix == "é" ||
			suffix == "emos" || suffix == "es" || suffix == "éis" {
			stem = ENDS_WITH_C.ReplaceAllString(stem, "qu")
			stem = ENDS_WITH_GU.ReplaceAllString(stem, "gü")
			stem = ENDS_WITH_G.ReplaceAllString(stem, "gu")
			stem = ENDS_WITH_Z.ReplaceAllString(stem, "c")
		}

		return (form == stem+suffix) ||
			ENDS_WITH_ASE.ReplaceAllString(form, "ara$1") == stem+suffix ||
			ENDS_WITH_ÁSEMOS.ReplaceAllString(form, "áramos") == stem+suffix
	}
}

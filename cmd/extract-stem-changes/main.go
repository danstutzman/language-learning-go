package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var BEGINS_WITH_AHI = regexp.MustCompile("^ahi")

var BEGINS_WITH_AHU = regexp.MustCompile("^ahu")

var BEGINS_WITH_ER = regexp.MustCompile("^er")

var ENDS_WITH_AC = regexp.MustCompile("ac$")

var ENDS_WITH_C = regexp.MustCompile("c$")

var ENDS_WITH_G = regexp.MustCompile("g$")

var ENDS_WITH_GU = regexp.MustCompile("gu$")

var ENDS_WITH_GÜ = regexp.MustCompile("gü$")

var ENDS_WITH_OLV = regexp.MustCompile("olv$")

var ENDS_WITH_PON = regexp.MustCompile("pon$")

var ENDS_WITH_QUER = regexp.MustCompile("quer$")

var ENDS_WITH_QUIR = regexp.MustCompile("quir$")

var ENDS_WITH_ROMP = regexp.MustCompile("romp$")

var ENDS_WITH_SC = regexp.MustCompile("sc$")

var ENDS_WITH_TEN = regexp.MustCompile("ten$")

var ENDS_WITH_I = regexp.MustCompile("i$")

var ENDS_WITH_U = regexp.MustCompile("u$")

var ENDS_WITH_UN = regexp.MustCompile("un$")

var ENDS_WITH_Z = regexp.MustCompile("z$")

var IS_CAB = regexp.MustCompile("^cab$")

var IS_HAB = regexp.MustCompile("^hab$")

var IS_POD = regexp.MustCompile("^pod$")

var IS_SAB = regexp.MustCompile("^sab$")

var LAST_SYLLABLE_HAS_AI = regexp.MustCompile(
	"(^|[bcdfghjlmnprstvzñ]|qu)ai([bcdfghjlmnprstvzñ]+|gu)$")

var LAST_SYLLABLE_HAS_AU = regexp.MustCompile(
	"(^|[bcdfghjlmnprstvzñ]|qu)au([bcdfghjlmnprstvzñ]+|gu)$")

var LAST_SYLLABLE_HAS_E = regexp.MustCompile(
	"(^|[bcdfghjlmnprstvzñ]|qu)e([bcdfghjlmnprstvzñ]+|gu)$")

var LAST_SYLLABLE_HAS_EI = regexp.MustCompile(
	"(^|[bcdfghjlmnprstvzñ]|qu)ei([bcdfghjlmnprstvzñ]+|gu)$")

var LAST_SYLLABLE_HAS_I = regexp.MustCompile(
	"(^|[bcdfghjlmnprstvzñ]|qu)i([bcdfghjlmnprstvzñ]+|gu)$")

var LAST_SYLLABLE_HAS_O = regexp.MustCompile(
	"(^|[bcdfghjlmnprstvzñ]|qu)o([bcdfghjlmnprstvzñ]+|gu)$")

var LAST_SYLLABLE_HAS_U = regexp.MustCompile(
	"(^|[bcdfghjlmnprstvzñ]|qu)u([bcdfghjlmnprstvzñ]+|gu)$")

func maybeReplace(stems []string, regex *regexp.Regexp,
	replacement string) []string {

	newStems := []string{}
	for _, stem := range stems {
		if regex.MatchString(stem) {
			newStems = append(newStems, regex.ReplaceAllString(stem, replacement))
		}
	}

	return append(stems, newStems...)
}

func main() {
	if len(os.Args) != 1+1 {
		log.Fatalf("1st arg: path to dicc.src")
	}
	freelingDiccPath := os.Args[1]

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

			if strings.HasPrefix(tag, "V") && strings.HasSuffix(lemma, "ar") {
				expectedStem := lemma[0 : len(lemma)-2]

				stems := []string{expectedStem}
				if tag == "VMIP1P0" || tag == "VMIP1S0" || tag == "VMIP2S0" ||
					tag == "VMIP3P0" || tag == "VMIP3S0" || tag == "VMM02S0" ||
					tag == "VMM03S0" || tag == "VMM03P0" || tag == "VMSP1S0" ||
					tag == "VMSP3S0" || tag == "VMSP3P0" || tag == "VMSP2S0" ||
					tag == "VMIS3S0" || tag == "VMIS3P0" || tag == "VMIS1S0" ||
					tag == "VMIS1P0" || tag == "VMIS2S0" || tag == "VMIS2P0" {
					stems = maybeReplace(stems, LAST_SYLLABLE_HAS_AI, "${1}aí${2}")
					stems = maybeReplace(stems, LAST_SYLLABLE_HAS_AU, "${1}aú${2}")
					stems = maybeReplace(stems, LAST_SYLLABLE_HAS_E, "${1}ie${2}")
					stems = maybeReplace(stems, LAST_SYLLABLE_HAS_E, "${1}i${2}")
					stems = maybeReplace(stems, LAST_SYLLABLE_HAS_EI, "${1}eí${2}")
					stems = maybeReplace(stems, LAST_SYLLABLE_HAS_I, "${1}í${2}")
					stems = maybeReplace(stems, LAST_SYLLABLE_HAS_O, "${1}ue${2}")
					stems = maybeReplace(stems, LAST_SYLLABLE_HAS_O, "${1}üe${2}")
					stems = maybeReplace(stems, LAST_SYLLABLE_HAS_O, "${1}hue${2}")
					stems = maybeReplace(stems, LAST_SYLLABLE_HAS_U, "${1}ue${2}")
					stems = maybeReplace(stems, LAST_SYLLABLE_HAS_U, "${1}ú${2}")
					stems = maybeReplace(stems, ENDS_WITH_I, "í")
					stems = maybeReplace(stems, ENDS_WITH_U, "ú")
					stems = maybeReplace(stems, BEGINS_WITH_AHI, "ahí")
					stems = maybeReplace(stems, BEGINS_WITH_AHU, "ahú")
				}

				if strings.HasSuffix(lemma, "ar") {
					if tag == "VMIS1S0" || tag == "VMM03S0" ||
						tag == "VMM01P0" || tag == "VMM03P0" ||
						tag == "VMSP1S0" || tag == "VMSP3S0" ||
						tag == "VMSP1P0" || tag == "VMSP3P0" ||
						tag == "VMSP2S0" || tag == "VMSP2P0" {
						stems = maybeReplace(stems, ENDS_WITH_C, "qu")
						stems = maybeReplace(stems, ENDS_WITH_GU, "gü")
						stems = maybeReplace(stems, ENDS_WITH_Z, "c")
					}
				} // end if has suffix -ar

				if strings.HasSuffix(lemma, "er") {
					if tag == "VMIP1S0" {
						stems = maybeReplace(stems, ENDS_WITH_C, "") // -go
						stems = maybeReplace(stems, ENDS_WITH_G, "j")
					}
					if tag == "VMSP1S0" || tag == "VMSP3S0" || tag == "VMSP1P0" ||
						tag == "VMSP3P0" || tag == "VMSP2S0" || tag == "VMSP2P0" {
						stems = maybeReplace(stems, ENDS_WITH_G, "j")
						stems = maybeReplace(stems, ENDS_WITH_SC, "zc")
					}
					if tag == "VMIS1S0" || tag == "VMIS3P0" || tag == "VMIS1P0" ||
						tag == "VMIS2S0" || tag == "VMIS2P0" || tag == "VMIS3S0" ||
						tag == "VMSI1S0" || tag == "VMSI3S0" || tag == "VMSI2P0" ||
						tag == "VMSI3P0" || tag == "VMSI2S0" || tag == "VMSF1S0" ||
						tag == "VMSF3S0" || tag == "VMSF2P0" || tag == "VMSF3P0" ||
						tag == "VMSF2S0" || tag == "VMSI1P0" || tag == "VMSF1P0" {
						stems = maybeReplace(stems, ENDS_WITH_AC, "ic")
						stems = maybeReplace(stems, ENDS_WITH_AC, "iz")
						stems = maybeReplace(stems, ENDS_WITH_PON, "pus")
						stems = maybeReplace(stems, ENDS_WITH_QUER, "quis")
						stems = maybeReplace(stems, ENDS_WITH_TEN, "tuv")
						stems = maybeReplace(stems, IS_POD, "pud")
						stems = maybeReplace(stems, IS_SAB, "sup")
						stems = maybeReplace(stems, IS_HAB, "hub")
						stems = maybeReplace(stems, IS_CAB, "cup")
					}
					if tag == "VMIC1S0" || tag == "VMIC3S0" || tag == "VMIC2P0" ||
						tag == "VMIC1P0" || tag == "VMIC3P0" || tag == "VMIC2S0" ||
						tag == "VMIF1P0" || tag == "VMIF3S0" || tag == "VMIF3P0" ||
						tag == "VMIF2S0" || tag == "VMIF1S0" || tag == "VMIF2P0" {
						stems = maybeReplace(stems, ENDS_WITH_C, "r")
					}
					if tag == "VMSP1S0" || tag == "VMSP1P0" || tag == "VMSP2P0" ||
						tag == "VMSP2S0" || tag == "VMSP3S0" || tag == "VMSP3P0" {
						stems = maybeReplace(stems, ENDS_WITH_C, "")
						stems = maybeReplace(stems, IS_CAB, "quep")
						stems = maybeReplace(stems, IS_HAB, "hay")
						stems = maybeReplace(stems, IS_SAB, "sep")
					}
					if tag == "VMM02S0" || tag == "VMM03S0" || tag == "VMM01P0" ||
						tag == "VMM03P0" {
						stems = maybeReplace(stems, ENDS_WITH_C, "z")
						stems = maybeReplace(stems, ENDS_WITH_SC, "zc")
					}
					if tag == "VMM03S0" || tag == "VMM01P0" || tag == "VMM03P0" {
						stems = maybeReplace(stems, ENDS_WITH_C, "g")
						stems = maybeReplace(stems, ENDS_WITH_G, "j")
						stems = maybeReplace(stems, IS_CAB, "quep")
						stems = maybeReplace(stems, IS_SAB, "sep")
					}
					if tag == "VMM02S0" {
						stems = maybeReplace(stems, ENDS_WITH_PON, "pón")
						stems = maybeReplace(stems, ENDS_WITH_TEN, "tén")
					}
					if tag == "VMP00SF" || tag == "VMP00PF" ||
						tag == "VMP00SM" || tag == "VMP00PM" {
						stems = maybeReplace(stems, ENDS_WITH_AC, "ech")
						stems = maybeReplace(stems, ENDS_WITH_PON, "pues")
						stems = maybeReplace(stems, ENDS_WITH_OLV, "uelt")
						stems = maybeReplace(stems, ENDS_WITH_ROMP, "rot")
					}
				} // end if has suffix -er

				if strings.HasSuffix(lemma, "ir") {
					if tag == "VMIP1S0" {
						stems = maybeReplace(stems, ENDS_WITH_C, "") // -go
						//stems = maybeReplace(stems, ENDS_WITH_G, "j")
					}
					if tag == "VMIP1S0" || tag == "VMIP3S0" ||
						tag == "VMIP3P0" || tag == "VMIP2S0" {
						stems = maybeReplace(stems, BEGINS_WITH_ER, "yer")
						stems = maybeReplace(stems, BEGINS_WITH_ER, "ir")
					}
					if tag == "VMIP1S0" {
						stems = maybeReplace(stems, ENDS_WITH_G, "j")
					}
					if tag == "VMIP1S0" {
						stems = maybeReplace(stems, ENDS_WITH_GU, "g")
					}
					if tag == "VMIP3S0" || tag == "VMIP3P0" ||
						tag == "VMIP2S0" || tag == "VMIP1S0" {
						stems = maybeReplace(stems, ENDS_WITH_UN, "ún")
						stems = maybeReplace(stems, ENDS_WITH_QUIR, "quier")
						stems = maybeReplace(stems, ENDS_WITH_GÜ, "guy")
					}
					if tag == "VMIS1S0" || tag == "VMIS3P0" || tag == "VMIS1P0" ||
						tag == "VMIS2S0" || tag == "VMIS2P0" || tag == "VMIS3S0" {
						stems = maybeReplace(stems, ENDS_WITH_C, "j")
					}
				} // end if has suffix -ir

				foundStem := false
				for _, possibleStem := range stems {
					if strings.HasPrefix(form, possibleStem) {
						foundStem = true
						break
					}
				}
				if !foundStem {
					fmt.Printf("%-20s %-20s %-10s %v\n", lemma, form, tag, stems)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

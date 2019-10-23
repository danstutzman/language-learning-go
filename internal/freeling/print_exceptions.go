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
			if strings.HasSuffix(lemma, "ar") && strings.HasPrefix(tag, "V") {
				conjugations := analyzeArVerb(lemma, tag)
				expected = false
				for _, conjugation := range conjugations {
					if conjugation.stem+conjugation.suffix == form {
						expected = true
					}
				}
			}
			if strings.HasSuffix(lemma, "er") && strings.HasPrefix(tag, "V") {
				conjugations := analyzeErVerb(lemma, tag)
				expected = false
				for _, conjugation := range conjugations {
					if conjugation.stem+conjugation.suffix == form {
						expected = true
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
				/*if !expected {
					for _, conjugation := range conjugations {
						fmt.Printf("%v\n", conjugation)
					}
				}*/
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
	suffixes := ER_MUTANT_TAG27_TO_SUFFIXES[tag[2:7]]
	if len(suffixes) > 0 && ER_MUTANTS[lemma] != "" {
		conjugations := []Conjugation{}
		for _, suffix := range suffixes {
			stem := ER_MUTANTS[lemma]
			if strings.HasSuffix(stem, "j") &&
				(strings.HasPrefix(suffix, "ie") || strings.HasPrefix(suffix, "ié")) {
				suffix = suffix[1:len(suffix)] // j- -iera -> -jera
			}
			conjugation := Conjugation{stem: stem, suffix: suffix}
			conjugations = append(conjugations, conjugation)
		}
		return conjugations
	}

	suffix := ER_MUTANT2_TAG27_TO_SUFFIX[tag[2:7]]
	if suffix != "" && ER_MUTANT2S[lemma] != "" {
		stem := ER_MUTANT2S[lemma]
		return []Conjugation{{stem: stem, suffix: suffix}}
	}

	stems := ER_VMP_STEMS[lemma]
	if len(stems) > 0 && strings.HasPrefix(tag, "VMP") {
		suffixes := ER_TAG27_TO_SUFFIXES[tag[2:7]]
		conjugations := []Conjugation{}
		for _, suffix := range suffixes {
			suffix = suffix[2:len(suffix)] // Remove initial -id
			for _, stem := range stems {
				conjugation := Conjugation{stem: stem, suffix: suffix}
				conjugations = append(conjugations, conjugation)
			}
		}
		return conjugations
	}

	stems = ER_VMSP1S_STEMS[lemma]
	if len(stems) > 0 && (tag == "VMSP1S0" || tag == "VMSP3S0" ||
		tag == "VMSP2S0" || tag == "VMSP3P0" || tag == "VMM03S0" ||
		tag == "VMM03P0" || tag == "VMIP1S0") {
		conjugations := []Conjugation{}
		suffixes := ER_TAG27_TO_SUFFIXES[tag[2:7]]
		for _, suffix := range suffixes {
			for _, stem := range stems {
				conjugation := Conjugation{stem: stem, suffix: suffix}
				conjugations = append(conjugations, conjugation)
			}
		}
		return conjugations
	}

	stems = ER_VMSP1P_STEMS[lemma]
	if len(stems) > 0 &&
		(tag == "VMSP1P0" || tag == "VMSP2P0" || tag == "VMM01P0") {
		suffixes := ER_TAG27_TO_SUFFIXES[tag[2:7]]
		conjugations := []Conjugation{}
		for _, suffix := range suffixes {
			for _, stem := range stems {
				conjugation := Conjugation{stem: stem, suffix: suffix}
				conjugations = append(conjugations, conjugation)
			}
		}
		return conjugations
	}

	if ER_VMM2S_FORMS[lemma] != "" && tag == "VMM02S0" {
		return []Conjugation{{stem: ER_VMM2S_FORMS[lemma], suffix: ""}}
	}

	stem := ER_STEM_CHANGES[lemma]
	if stem == "" || !ER_TAG_TO_STEM_CHANGE[tag] {
		stem = lemma[0 : len(lemma)-2]
	}

	suffixes = ER_TAG27_TO_SUFFIXES[tag[2:7]]
	conjugations := []Conjugation{}
	for _, suffix := range suffixes {
		if suffix == "o" {
			stem = ENDS_WITH_G.ReplaceAllString(stem, "j") // Before -g replacements
			stem = ENDS_WITH_AC.ReplaceAllString(stem, "ag")
			stem = ENDS_WITH_EC_IC_OR_OC.ReplaceAllString(stem, "${1}zc")
			stem = ENDS_WITH_N.ReplaceAllString(stem, "ng")
		}

		if ENDS_WITH_A_E_OR_O.MatchString(stem) {
			if suffix == "imos" {
				suffix = "ímos"
			} else if suffix == "iste" {
				suffix = "íste"
			} else if suffix == "isteis" {
				suffix = "ísteis"
			} else if suffix == "ió" {
				suffix = "yó"
			} else if suffix == "ieron" {
				suffix = "yeron"
			} else if suffix == "iendo" {
				suffix = "yendo"
			} else if suffix == "iera" {
				suffix = "yera"
			} else if suffix == "ierais" {
				suffix = "yerais"
			} else if suffix == "ieran" {
				suffix = "yeran"
			} else if suffix == "ieras" {
				suffix = "yeras"
			} else if suffix == "iéramos" {
				suffix = "yéramos"
			} else if suffix == "iere" {
				suffix = "yere"
			} else if suffix == "iereis" {
				suffix = "yereis"
			} else if suffix == "ieren" {
				suffix = "yeren"
			} else if suffix == "ieres" {
				suffix = "yeres"
			} else if suffix == "iéremos" {
				suffix = "yéremos"
			} else if suffix == "iese" {
				suffix = "yese"
			} else if suffix == "ieseis" {
				suffix = "yeseis"
			} else if suffix == "iesen" {
				suffix = "yesen"
			} else if suffix == "ieses" {
				suffix = "yeses"
			} else if suffix == "iésemos" {
				suffix = "yésemos"
			}
		} else if strings.HasSuffix(stem, "ll") || strings.HasSuffix(stem, "ñ") {
			if suffix == "iendo" {
				suffix = "endo"
			} else if suffix == "ieron" {
				suffix = "eron"
			} else if suffix == "ió" {
				suffix = "ó"
			} else if suffix == "ierais" {
				suffix = "erais"
			} else if suffix == "ieras" {
				suffix = "eras"
			} else if suffix == "iera" {
				suffix = "era"
			} else if suffix == "ieran" {
				suffix = "eran"
			} else if suffix == "iéramos" {
				suffix = "éramos"
			} else if suffix == "iereis" {
				suffix = "ereis"
			} else if suffix == "ieres" {
				suffix = "eres"
			} else if suffix == "iere" {
				suffix = "ere"
			} else if suffix == "ieren" {
				suffix = "eren"
			} else if suffix == "iéremos" {
				suffix = "éremos"
			} else if suffix == "iese" {
				suffix = "ese"
			} else if suffix == "ieseis" {
				suffix = "eseis"
			} else if suffix == "iesen" {
				suffix = "esen"
			} else if suffix == "ieses" {
				suffix = "eses"
			} else if suffix == "iésemos" {
				suffix = "ésemos"
			}
		} else if ENDS_WITH_N.MatchString(stem) && strings.HasPrefix(suffix, "er") &&
			suffix != "er" {
			suffix = "dr" + suffix[2:len(suffix)] // ería => dría
		} else if VER_VERBS[lemma] &&
			(strings.HasPrefix(tag, "VMII") || strings.HasPrefix(tag, "VMM03")) {
			suffix = "e" + suffix // ven * VMII1S0 => veía not vía
		}
		conjugation := Conjugation{stem: stem, suffix: suffix}
		conjugations = append(conjugations, conjugation)
	}
	return conjugations
}

func analyzeIrVerb(lemma, tag string) []Conjugation {
	suffixes := IR_MUTANT_TAG_TO_SUFFIXES[tag]
	if len(suffixes) > 0 && len(IR_MUTANTS[lemma]) > 0 {
		stems := IR_MUTANTS[lemma]
		conjugations := []Conjugation{}
		for _, suffix := range suffixes {
			for _, stem := range stems {
				if strings.HasSuffix(stem, "j") &&
					(strings.HasPrefix(suffix, "ie") || strings.HasPrefix(suffix, "ié")) {
					suffix = suffix[1:len(suffix)] // j- -iera -> -jera
				}
				conjugation := Conjugation{stem: stem, suffix: suffix}
				conjugations = append(conjugations, conjugation)
			}
		}
		return conjugations
	}

	suffix := IR_MUTANT2_TAG_TO_SUFFIX[tag]
	if suffix != "" && IR_MUTANT2S[lemma] != "" {
		stem := IR_MUTANT2S[lemma]
		return []Conjugation{{stem: stem, suffix: suffix}}
	}

	stems := IR_VMP_STEMS[lemma]
	if len(stems) > 0 && strings.HasPrefix(tag, "VMP") {
		suffixes := IR_TAG_TO_SUFFIXES[tag]
		conjugations := []Conjugation{}
		for _, stem := range stems {
			for _, suffix := range suffixes {
				suffix = suffix[2:len(suffix)] // Remove initial -id
				conjugation := Conjugation{stem: stem, suffix: suffix}
				conjugations = append(conjugations, conjugation)
			}
		}
		return conjugations
	}

	stems = IR_VMSP1S_STEMS[lemma]
	if len(stems) > 0 &&
		(tag == "VMSP1S0" || tag == "VMSP2S0" || tag == "VMSP3S0" ||
			tag == "VMSP3P0" || tag == "VMM03S0" || tag == "VMM03P0") {
		conjugations := []Conjugation{}
		for _, stem := range stems {
			suffixes := IR_TAG_TO_SUFFIXES[tag]
			for _, suffix := range suffixes {
				conjugation := Conjugation{stem: stem, suffix: suffix}
				conjugations = append(conjugations, conjugation)
			}
		}
		return conjugations
	}

	stems = IR_VMSP1P_STEMS[lemma]
	if len(stems) > 0 &&
		(tag == "VMSP1P0" || tag == "VMSP2P0" || tag == "VMM01P0") {
		suffixes := IR_TAG_TO_SUFFIXES[tag]
		conjugations := []Conjugation{}
		for _, stem := range stems {
			for _, suffix := range suffixes {
				conjugation := Conjugation{stem: stem, suffix: suffix}
				conjugations = append(conjugations, conjugation)
			}
		}
		return conjugations
	}

	if len(IR_VMM2S_FORMS[lemma]) > 0 && tag == "VMM02S0" {
		conjugations := []Conjugation{}
		for _, form := range IR_VMM2S_FORMS[lemma] {
			conjugation := Conjugation{stem: form, suffix: ""}
			conjugations = append(conjugations, conjugation)
		}
		return conjugations
	}

	if len(IR_E_TO_I_STEMS[lemma]) > 0 && (tag == "VMIS3S0" || tag == "VMIS3P0" ||
		tag == "VMG0000" || tag == "VMSF1S0" || tag == "VMSF1P0" ||
		tag == "VMSF2S0" || tag == "VMSF2P0" || tag == "VMSF3S0" ||
		tag == "VMSF3P0" || tag == "VMSI1S0" || tag == "VMSI1P0" ||
		tag == "VMSI2S0" || tag == "VMSI2P0" || tag == "VMSI3S0" ||
		tag == "VMSI3P0") {
		stems = IR_E_TO_I_STEMS[lemma]
	} else if len(IR_STEM_CHANGES[lemma]) > 0 &&
		(tag == "VMIP2S0" || tag == "VMIP3S0" || tag == "VMIP3P0" ||
			tag == "VMIP1S0") {
		stems = IR_STEM_CHANGES[lemma]
	} else {
		stems = []string{lemma[0 : len(lemma)-2]}
	}

	suffixes = IR_TAG_TO_SUFFIXES[tag]

	conjugations := []Conjugation{}
	for _, stem := range stems {
		for _, suffix := range suffixes {
			if ENDS_WITH_A_E_O_OR_U.MatchString(stem) {
				if suffix == "ió" {
					suffix = "yó"
				} else if suffix == "ieron" {
					suffix = "yeron"
				} else if suffix == "iendo" {
					suffix = "yendo"
				} else if suffix == "iera" {
					suffix = "yera"
				} else if suffix == "ierais" {
					suffix = "yerais"
				} else if suffix == "ieran" {
					suffix = "yeran"
				} else if suffix == "ieras" {
					suffix = "yeras"
				} else if suffix == "iéramos" {
					suffix = "yéramos"
				} else if suffix == "iere" {
					suffix = "yere"
				} else if suffix == "iereis" {
					suffix = "yereis"
				} else if suffix == "ieren" {
					suffix = "yeren"
				} else if suffix == "ieres" {
					suffix = "yeres"
				} else if suffix == "iéremos" {
					suffix = "yéremos"
				} else if suffix == "iese" {
					suffix = "yese"
				} else if suffix == "ieseis" {
					suffix = "yeseis"
				} else if suffix == "iesen" {
					suffix = "yesen"
				} else if suffix == "ieses" {
					suffix = "yeses"
				} else if suffix == "iésemos" {
					suffix = "yésemos"
				}
			} else if strings.HasSuffix(stem, "ll") || strings.HasSuffix(stem, "ñ") ||
				strings.HasSuffix(stem, "y") {
				if suffix == "ió" {
					suffix = "ó"
				} else if suffix == "ieron" {
					suffix = "eron"
				} else if suffix == "iendo" {
					suffix = "endo"
				} else if suffix == "iera" {
					suffix = "era"
				} else if suffix == "ierais" {
					suffix = "erais"
				} else if suffix == "ieran" {
					suffix = "eran"
				} else if suffix == "ieras" {
					suffix = "eras"
				} else if suffix == "iéramos" {
					suffix = "éramos"
				} else if suffix == "iere" {
					suffix = "ere"
				} else if suffix == "iereis" {
					suffix = "ereis"
				} else if suffix == "ieren" {
					suffix = "eren"
				} else if suffix == "ieres" {
					suffix = "eres"
				} else if suffix == "iéremos" {
					suffix = "éremos"
				} else if suffix == "iese" {
					suffix = "ese"
				} else if suffix == "ieseis" {
					suffix = "eseis"
				} else if suffix == "iesen" {
					suffix = "esen"
				} else if suffix == "ieses" {
					suffix = "eses"
				} else if suffix == "iésemos" {
					suffix = "ésemos"
				}
			}

			conjugation := Conjugation{stem: stem, suffix: suffix}
			conjugations = append(conjugations, conjugation)
		}
	}
	return conjugations
}

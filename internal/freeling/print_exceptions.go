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
			if strings.HasSuffix(lemma, "er") && strings.HasPrefix(tag, "V") {
				expected = isExpectedErVerb(lemma, form, tag)
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

func isExpectedErVerb(lemma, form, tag string) bool {
	suffix := ER_MUTANT_TAG27_TO_SUFFIX[tag[2:7]]
	if suffix != "" && ER_MUTANTS[lemma] != "" {
		stem := ER_MUTANTS[lemma]
		if strings.HasSuffix(stem, "j") &&
			(strings.HasPrefix(suffix, "ie") || strings.HasPrefix(suffix, "ié")) {
			suffix = suffix[1:len(suffix)] // j- -iera -> -jera
		}
		return (form == stem+suffix) ||
			ENDS_WITH_ESE.ReplaceAllString(form, "era$1") == stem+suffix ||
			ENDS_WITH_ÉSEMOS.ReplaceAllString(form, "éramos") == stem+suffix
	}

	suffix = ER_MUTANT2_TAG27_TO_SUFFIX[tag[2:7]]
	if suffix != "" && ER_MUTANT2S[lemma] != "" {
		stem := ER_MUTANT2S[lemma]
		return form == stem+suffix
	}

	stem := ER_VMP_STEMS[lemma]
	if stem != "" && strings.HasPrefix(tag, "VMP") {
		suffix = ER_TAG_TO_SUFFIX[tag]
		suffix = suffix[2:len(suffix)] // Remove initial -id
		return form == stem+suffix
	}

	stem = ER_VMSP1S_STEMS[lemma]
	if stem != "" && (tag == "VMSP1S0" || tag == "VMSP3S0" ||
		tag == "VMSP2S0" || tag == "VMSP3P0" || tag == "VMM03S0" ||
		tag == "VMM03P0") {
		suffix = ER_TAG_TO_SUFFIX[tag]
		return form == stem+suffix
	}

	stem = ER_VMSP1P_STEMS[lemma]
	if stem != "" && (tag == "VMSP1P0" || tag == "VMSP2P0" || tag == "VMM01P0") {
		suffix = ER_TAG_TO_SUFFIX[tag]
		return form == stem+suffix
	}

	if ER_VMM2S_FORMS[lemma] != "" && tag == "VMM02S0" {
		return form == ER_VMM2S_FORMS[lemma]
	}

	stem = ER_STEM_CHANGES[lemma]
	if stem == "" || !ER_TAG_TO_STEM_CHANGE[tag] {
		stem = lemma[0 : len(lemma)-2]
	}

	suffix = ER_TAG_TO_SUFFIX[tag]
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
		}
	} else if strings.HasSuffix(stem, "ll") || strings.HasSuffix(stem, "ñ") {
		if suffix == "ierais" {
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
		}
	} else if ENDS_WITH_N.MatchString(stem) && strings.HasPrefix(suffix, "er") &&
		suffix != "er" {
		suffix = "dr" + suffix[2:len(suffix)] // ería => dría
	} else if VER_VERBS[lemma] &&
		(strings.HasPrefix(tag, "VMII") || strings.HasPrefix(tag, "VMM03")) {
		suffix = "e" + suffix // ven * VMII1S0 => veía not vía
	}

	return (form == stem+suffix) ||
		ENDS_WITH_ESE.ReplaceAllString(form, "era$1") == stem+suffix ||
		ENDS_WITH_ÉSEMOS.ReplaceAllString(form, "éramos") == stem+suffix
}

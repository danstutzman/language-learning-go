package freeling

import (
	"regexp"
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

var ENDS_WITH_ASE = regexp.MustCompile("ase(is|s|n)?$")
var ENDS_WITH_ÁSEMOS = regexp.MustCompile("ásemos$")
var ENDS_WITH_ESE = regexp.MustCompile("ese(is|s|n)?$")
var ENDS_WITH_ÉSEMOS = regexp.MustCompile("ésemos$")

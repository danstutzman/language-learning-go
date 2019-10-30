/*
The MIT License (MIT)

Copyright (c) 2017 Alex Corvi

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package english

import (
	"fmt"
	"regexp"
)

var regexpAppendix = regexp.MustCompile(`dix$`)
var regexpPooch = regexp.MustCompile(`ooch$`)
var regexpMan = regexp.MustCompile(`(m)an$`)
var regexpPerson = regexp.MustCompile(`(pe)rson$`)
var regexpChild = regexp.MustCompile(`(child)$`)
var regexpOx = regexp.MustCompile(`^(ox)$`)
var regexpAxis = regexp.MustCompile(`(ax|test)is$`)
var regexpFungus = regexp.MustCompile(`(op|ir|umn|am|ll|ct|oc|ng|le|di|ul|ab|rmin|or|in)us$`)
var regexpStatus = regexp.MustCompile(`(alias|status)$`)
var regexpsyllabus = regexp.MustCompile(`(bu)s$`)
var regexpBuffalo = regexp.MustCompile(`(fal|tat|ch|rg|ott|mat|ped|et|can|er|uit|her|ad)o$`)
var regexpBacterium = regexp.MustCompile(`([aeiouy]ri|dat|cul|rat|nasi|edi|rand|ov)um$`)
var regexpCriterion = regexp.MustCompile(`([aoeuiy]|er)ion$`)
var regexpSherion = regexp.MustCompile(`(mat|erio|omen|hedr)on$`)
var regexpAnalysis = regexp.MustCompile(`(is|ps|hes|as|ys|os|ax)is$`)
var regexpCalf = regexp.MustCompile(`(?:([^f])fe|([lrf])f)$`)
var regexpHive = regexp.MustCompile(`(hive)$`)
var regexpAlly = regexp.MustCompile(`([^aeiouy]|qu)y$`)
var regexpAlley = regexp.MustCompile(`([aeiouy]y)$`)
var regexpMatrix = regexp.MustCompile(`(matr|vert|ind)(ix|ex)$`)
var regexpLouse = regexp.MustCompile(`([m|l])ouse$`)
var regexpAlga = regexp.MustCompile(`(alg|lumn|tenn|arv|ebul|pup|rteb|vit)a$`)
var regexpBuzz = regexp.MustCompile(`(uz|qui|ut)(z)$`)
var regexpWaltz = regexp.MustCompile(`(opa|alt)(z)$`)
var regexpFoot = regexp.MustCompile(`^(f|t|g)oo([thse]{1,2})$`)
var regexpPlateau = regexp.MustCompile(`([^aeiouy])eau$`)
var regexpLoaf = regexp.MustCompile(`([aeiouy])f$`)
var regexpArch = regexp.MustCompile(`(x|ch|ss|sh|s)$`)
var regexpO = regexp.MustCompile(`o$`)

var singular2plural = map[string][]string{
	"abacus":              {"abacuses"},
	"abyss":               {"abysses"},
	"addendum":            {"addenda"},
	"agenda":              {"agendas"},
	"agendum":             {"agenda"},
	"agent-provocateur":   {"agents-provocateurs"},
	"aide-de-camp":        {"aides-de-camp"},
	"aircraft":            {"aircraft"},
	"albino":              {"albinos"},
	"album":               {"albums"},
	"alfalfa":             {"alfalfas"},
	"alga":                {"algae"},
	"alumna":              {"alumnae"},
	"alumnus":             {"alumni"},
	"amoeba":              {"amoebae"},
	"analysis":            {"analyses"},
	"anathema":            {"anathemata"},
	"annex":               {"annexes"},
	"antenna":             {"antennas"},
	"antithesis":          {"antitheses"},
	"apex":                {"apices", "apexes"},
	"apparatus":           {"apparatuses"},
	"appendix":            {"appendices", "appendixes"},
	"apple":               {"apples"},
	"aquarium":            {"aquaria", "aquariums"},
	"arch":                {"arches"},
	"armadillo":           {"armadillos"},
	"ass":                 {"asses"},
	"atlas":               {"atlases"},
	"attorne-general":     {"attorneys-general"},
	"aurora":              {"auroras"},
	"auto":                {"autos"},
	"axe":                 {"axes"},
	"axis":                {"axes"},
	"baby":                {"babies"},
	"bacillus":            {"bacilli"},
	"bacterium":           {"bacteria"},
	"balloon":             {"balloons"},
	"banana":              {"bananas"},
	"barracks":            {"barracks"},
	"barracuda":           {"barracudas"},
	"basis":               {"bases"},
	"batch":               {"batches"},
	"beach":               {"beaches"},
	"beau":                {"beaux"},
	"beau-geste":          {"beaux-gestes"},
	"bel-homme":           {"beaux-hommes"},
	"belief":              {"beliefs"},
	"belle-epoque":        {"belles-epoques", "belle-epoques"},
	"bikini":              {"bikinis"},
	"bildungsroman":       {"bildungsromane"},
	"bill-of-attainder":   {"bills-of-attainder"},
	"biscotto":            {"biscotti"},
	"bon-mot":             {"bons-mots"},
	"bon-vivant":          {"bons-vivants"},
	"book":                {"books"},
	"box":                 {"boxes"},
	"brother":             {"brothers"},
	"brush":               {"brushes"},
	"buffalo":             {"buffalos", "buffaloes"},
	"bureau":              {"bureaus", "bureaux"},
	"bus":                 {"buses", "busses"},
	"cactus":              {"cacti"},
	"calf":                {"calves"},
	"cameo":               {"cameos"},
	"candelabrum":         {"candelabra"},
	"canto":               {"cantos"},
	"carton":              {"cartons"},
	"cat-o'-nine-tails":   {"cat-o'-nine-tails"},
	"cello":               {"cellos"},
	"chateau":             {"chateaux", "chateaus"},
	"cherry":              {"cherries"},
	"chick":               {"chicks"},
	"chicken":             {"chickens"},
	"chief":               {"chiefs"},
	"child":               {"children"},
	"church":              {"churches"},
	"château":             {"châteaux", "châteaus"},
	"circus":              {"circuses"},
	"city":                {"cities"},
	"cod":                 {"cods"},
	"codex":               {"codices"},
	"combo":               {"combos"},
	"complex":             {"complexes"},
	"concerto":            {"concerti", "concertos"},
	"copy":                {"copies"},
	"cornea":              {"corneas"},
	"corps":               {"corps"},
	"corpus":              {"corpora", "corpuses"},
	"coup d'etat":         {"coups d'etat"},
	"court-martial":       {"courts-martial"},
	"cri du coeur":        {"cris du coeur"},
	"crisis":              {"crises"},
	"criterion":           {"criteria"},
	"crocus":              {"crocuses"},
	"curriculum":          {"curricula"},
	"daisy":               {"daisies"},
	"datum":               {"data"},
	"des":                 {"deses"},
	"diagnosis":           {"diagnoses"},
	"dictionary":          {"dictionaries"},
	"die":                 {"dice"},
	"director-general":    {"directors-general"},
	"dogma":               {"dogmata"},
	"domino":              {"dominoes"},
	"dukhobor":            {"dukhobortsy"},
	"duo":                 {"duos"},
	"duplex":              {"duplexes"},
	"dwarf":               {"dwarves", "dwarfs"},
	"echo":                {"echoes"},
	"ego":                 {"egos"},
	"elf":                 {"elves"},
	"ellipsis":            {"ellipses"},
	"embargo":             {"embargoes"},
	"emphasis":            {"emphases"},
	"entente-cordiale":    {"ententes-cordiales"},
	"erratum":             {"errata"},
	"fait-accompli":       {"faits-accomplis"},
	"family":              {"families"},
	"faux-pas":            {"faux-pas"},
	"fax":                 {"faxes"},
	"fee simple absolute": {"fees simple absolute"},
	"fez":                 {"fezzes", "fezes"},
	"fireman":             {"firemen"},
	"fish":                {"fishes"},
	"flush":               {"flushes"},
	"fly":                 {"flies"},
	"focus":               {"foci", "focuses"},
	"folio":               {"folios"},
	"foot":                {"feet"},
	"formula":             {"formulas", "formulae"},
	"fungus":              {"fungi", "funguses"},
	"gallows":             {"gallows"},
	"gas":                 {"gases"},
	"gens":                {"gentes"},
	"genu":                {"genua"},
	"genus":               {"genera"},
	"goose":               {"geese"},
	"governor-general":    {"governors-general"},
	"graffito":            {"graffiti"},
	"grief":               {"griefs"},
	"grouse":              {"grouses"},
	"gulf":                {"gulfs"},
	"guru":                {"gurus"},
	"half":                {"halves"},
	"hallux":              {"halluces"},
	"halo":                {"halos"},
	"ham on rye":          {"ham-on-ryes"},
	"handkerchief":        {"handkerchiefs"},
	"head of state":       {"heads of states", "heads of state"},
	"hero":                {"heroes"},
	"hetero":              {"heteros"},
	"hex":                 {"hexes"},
	"hippopotamus":        {"hippopotami", "hippopotamuses"},
	"hoax":                {"hoaxes"},
	"holluschik":          {"holluschikie"},
	"hoof":                {"hooves"},
	"hypothesis":          {"hypotheses"},
	"idee-fixe":           {"idees-fixes"},
	"index":               {"indexes", "indices", "indeces"},
	"inferno":             {"infernos"},
	"insigne":             {"insignia"},
	"iris":                {"irises"},
	"iter":                {"itinera"},
	"jack-in-the-box":     {"jacks-in-the-box", "jack-in-the-boxes"},
	"jack-in-the-pulpit":  {"jacks-in-the-pulpit", "jack-in-the-pulpits"},
	"jack-o'-lantern":     {"jack-o'-lanterns"},
	"kerchief":            {"kerchiefs"},
	"kimono":              {"kimonos"},
	"kiss":                {"kisses"},
	"knife":               {"knives"},
	"krone":               {"kroner"},
	"lady":                {"ladies"},
	"larva":               {"larvae"},
	"lasso":               {"lassos"},
	"leaf":                {"leaves"},
	"lemma":               {"lemmata", "lemmas"},
	"libretto":            {"libretti", "librettos"},
	"life":                {"lives"},
	"loaf":                {"loaves"},
	"locus":               {"loci"},
	"louse":               {"lice"},
	"man":                 {"men"},
	"man-about-town":      {"men-about-town"},
	"man-child":           {"men-children"},
	"man-of-war":          {"men-of-war"},
	"mango":               {"mangoes"},
	"manservant":          {"menservants"},
	"martini":             {"martinis"},
	"matrix":              {"matrices", "matrixes"},
	"matzo":               {"matzoth"},
	"medium":              {"media"},
	"memento":             {"mementos"},
	"memo":                {"memos"},
	"memorandum":          {"memoranda"},
	"menu":                {"menus"},
	"mess":                {"messes"},
	"millennium":          {"millenniums", "millennium", "millennia"},
	"minister-president":  {"ministers-president"},
	"minutia":             {"minutiae"},
	"mischief":            {"mischiefs"},
	"moose":               {"moose"},
	"mosquito":            {"mosquitoes"},
	"motto":               {"mottoes"},
	"mouse":               {"mice"},
	"muff":                {"muffs"},
	"mussolini":           {"mussolinis"},
	"nanny":               {"nannies"},
	"nebula":              {"nebulae", "nebulas"},
	"neurosis":            {"neuroses"},
	"never-was":           {"never-weres"},
	"nova":                {"novas"},
	"nucleus":             {"nuclei"},
	"oaf":                 {"oafs"},
	"oasis":               {"oases"},
	"octopus":             {"octopuses", "octopodes", "octopi"},
	"opus":                {"opera", "operas"},
	"ornis":               {"ornithes"},
	"ovum":                {"ova"},
	"ox":                  {"oxen"},
	"page":                {"pages"},
	"pakistani":           {"pakistanis"},
	"panino":              {"panini"},
	"paparazzo":           {"paparazzi"},
	"paralysis":           {"paralyses"},
	"parenthesis":         {"parentheses"},
	"party":               {"parties"},
	"pass":                {"passes"},
	"passerby":            {"passersby"},
	"penny":               {"pennies"},
	"person":              {"people"},
	"phenomenon":          {"phenomena"},
	"phobia":              {"phobias"},
	"photo":               {"photos"},
	"phylum":              {"phyla"},
	"piano":               {"pianos"},
	"pitch":               {"pitches"},
	"plateau":             {"plateaux", "plateaus"},
	"poppy":               {"poppies"},
	"portfolio":           {"portfolios"},
	"portico":             {"porticos"},
	"potato":              {"potatoes"},
	"pro":                 {"pros"},
	"procurator-fiscal":   {"procurators-fiscal"},
	"prognosis":           {"prognoses"},
	"proof":               {"proofs"},
	"pupa":                {"pupae"},
	"quadrans":            {"quadrantes"},
	"quarto":              {"quartos"},
	"quiz":                {"quizzes"},
	"radius":              {"radii"},
	"referendum":          {"referenda", "referendums"},
	"reflex":              {"reflexes"},
	"rhombus":             {"rhombuses"},
	"roof":                {"roofs"},
	"rubai":               {"rubaiyat"},
	"runner-up":           {"runners-up"},
	"safe":                {"safes"},
	"salmon":              {"salmons"},
	"scarf":               {"scarves", "scarfs"},
	"schema":              {"schemata", "schemas"},
	"schoolchild":         {"schoolchildren"},
	"scratch":             {"scratches"},
	"scrotum":             {"scrota", "scrotums"},
	"self":                {"selves"},
	"seraph":              {"seraphim"},
	"sheaf":               {"sheaves"},
	"shelf":               {"shelves"},
	"ship of the line":    {"ships of the line"},
	"shrimp":              {"shrimps"},
	"shtetl":              {"shtetlach"},
	"silo":                {"silos"},
	"snafu":               {"snafus"},
	"solo":                {"solos"},
	"son of a bitch":      {"sons of bitches", "sons-of-a-bitch"},
	"son-in-law":          {"sons-in-law"},
	"spaghetto":           {"spaghettis", "spaghetti"},
	"splash":              {"splashes"},
	"spy":                 {"spies"},
	"stadium":             {"stadiums"},
	"stereo":              {"stereos"},
	"stigma":              {"stigmata", "stigmas"},
	"stimulus":            {"stimuli"},
	"stitch":              {"stitches"},
	"stoma":               {"stomata", "stomas"},
	"story":               {"stories"},
	"stratum":             {"strata"},
	"studio":              {"studios"},
	"syllabus":            {"syllabi", "syllabuses"},
	"symposium":           {"symposia", "symposiums"},
	"synopsis":            {"synopses"},
	"synthesis":           {"syntheses"},
	"tableau":             {"tableaux", "tableaus"},
	"taco":                {"tacos"},
	"tattoo":              {"tattoos"},
	"tax":                 {"taxes"},
	"tete-a-tete":         {"tete-a-tetes"},
	"that":                {"those"},
	"thesis":              {"theses"},
	"thief":               {"thieves"},
	"this":                {"these"},
	"tomato":              {"tomatoes"},
	"tooth":               {"teeth"},
	"tornado":             {"tornadoes"},
	"torpedo":             {"torpedoes"},
	"torte":               {"torten"},
	"tour-de-force":       {"tours-de-force"},
	"trout":               {"trouts"},
	"try":                 {"tries"},
	"tuna":                {"tunas"},
	"turf":                {"turfs"},
	"tuxedo":              {"tuxedos"},
	"typo":                {"typos"},
	"ushabti":             {"ushabtiu"},
	"vas":                 {"vasa"},
	"vertebra":            {"vertebrae"},
	"vertex":              {"vertices", "vertexes"},
	"veto":                {"vetoes"},
	"video":               {"videos"},
	"vita":                {"vitae"},
	"volcano":             {"volcanoes"},
	"vortex":              {"vortices", "vortexes"},
	"walrus":              {"walruses"},
	"waltz":               {"waltzes"},
	"wash":                {"washes"},
	"watch":               {"watches"},
	"wharf":               {"wharves"},
	"wife":                {"wives"},
	"will-o'-the-wisp":    {"will-o'-the-wisps"},
	"wish":                {"wishes"},
	"wolf":                {"wolves"},
	"woman":               {"women"},
	"woman-doctor":        {"women-doctors"},
	"word":                {"words"},
	"workman":             {"workmen"},
	"wunderkind":          {"wunderkinder"},
	"yemeni":              {"yemenis"},
	"yeshiva":             {"yeshivoth"},
	"yo":                  {"yos"},
	"zecchino":            {"zecchini"},
	"zero":                {"zeros", "zeroes"},
	"zoo":                 {"zoos"},
}

func PluralizeNoun(input string) string {
	plural := singular2plural[input]
	if len(plural) > 0 {
		return plural[0]
	}

	// appendix, spadix, radix
	if regexpAppendix.MatchString(input) {
		return regexpAppendix.ReplaceAllString(input, "dices")
	}
	// pooch
	if regexpPooch.MatchString(input) {
		return regexpPooch.ReplaceAllString(input, "$1chs")
	}
	// man policeman, fireman,
	if regexpMan.MatchString(input) {
		return regexpMan.ReplaceAllString(input, "$1en")
	}
	// person
	if regexpPerson.MatchString(input) {
		return regexpPerson.ReplaceAllString(input, "$1ople")
	}
	// child
	if regexpChild.MatchString(input) {
		return regexpChild.ReplaceAllString(input, "$1ren")
	}
	// ox
	if regexpOx.MatchString(input) {
		return regexpOx.ReplaceAllString(input, "$1en")
	}
	// axis
	if regexpAxis.MatchString(input) {
		return regexpAxis.ReplaceAllString(input, "$1es")
	}
	// fungus locus, nucleus, radius,
	if regexpFungus.MatchString(input) {
		return regexpFungus.ReplaceAllString(input, "$1i")
	}
	// status
	if regexpStatus.MatchString(input) {
		return regexpStatus.ReplaceAllString(input, "$1es")
	}
	// Syllabus thrombus
	if regexpsyllabus.MatchString(input) {
		return regexpsyllabus.ReplaceAllString(input, "$1ses")
	}
	// buffalo cargo, echo, embargo
	if regexpBuffalo.MatchString(input) {
		return regexpBuffalo.ReplaceAllString(input, "$1oes")
	}
	// bacterium curriculum, datum, erratum,
	if regexpBacterium.MatchString(input) {
		return regexpBacterium.ReplaceAllString(input, "$1a")
	}
	// criterion
	if regexpCriterion.MatchString(input) {
		return regexpCriterion.ReplaceAllString(input, "$1ia")
	}
	// sherion
	if regexpSherion.MatchString(input) {
		return regexpSherion.ReplaceAllString(input, "$1a")
	}
	// analysis, basis, crisis, ellipsis, hypothesis
	if regexpAnalysis.MatchString(input) {
		return regexpAnalysis.ReplaceAllString(input, "$1es")
	}
	// calf elf half knife
	if regexpCalf.MatchString(input) {
		return regexpCalf.ReplaceAllString(input, "$1$2ves")
	}
	// hive
	if regexpHive.MatchString(input) {
		return regexpHive.ReplaceAllString(input, "$1s")
	}
	// ally army baby beauty
	if regexpAlly.MatchString(input) {
		return regexpAlly.ReplaceAllString(input, "$1ies")
	}
	// alley attorney essay boy delay
	if regexpAlley.MatchString(input) {
		return regexpAlley.ReplaceAllString(input, "$1s")
	}
	// matrix vertex index
	if regexpMatrix.MatchString(input) {
		return regexpMatrix.ReplaceAllString(input, "$1ices")
	}
	// louse mouse booklouse
	if regexpLouse.MatchString(input) {
		return regexpLouse.ReplaceAllString(input, "$1ice")
	}
	// alga, alumna, antenna, larva
	if regexpAlga.MatchString(input) {
		return regexpAlga.ReplaceAllString(input, "$1ae")
	}
	// buzz fizz klutz quiz topaz waltz
	if regexpBuzz.MatchString(input) {
		return regexpBuzz.ReplaceAllString(input, "$1$2zes")
	}
	// waltz
	if regexpWaltz.MatchString(input) {
		return regexpWaltz.ReplaceAllString(input, "$1zes")
	}
	// foot tooth
	if regexpFoot.MatchString(input) {
		return regexpFoot.ReplaceAllString(input, "$1ee$2")
	}
	// plateau
	if regexpPlateau.MatchString(input) {
		return regexpPlateau.ReplaceAllString(input, "$1eaux")
	}
	// loaf
	if regexpLoaf.MatchString(input) {
		return regexpLoaf.ReplaceAllString(input, "$1ves")
	}
	// arch atlas ax bash bench
	if regexpArch.MatchString(input) {
		return regexpArch.ReplaceAllString(input, "$1es")
	}

	if regexpO.MatchString(input) {
		return regexpO.ReplaceAllString(input, "oes")
	}

	return input + "s"
}

func demoPluralizeNoun() {
	fmt.Println(PluralizeNoun("ox"))
}

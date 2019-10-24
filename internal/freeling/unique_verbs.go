package freeling

type UniqueVerb struct {
	infinitive string
	tag27      string
	form       string
}

var uniqueVerbs = []UniqueVerb{
	{"ir", "IS3S0", "fue"},
	{"ir", "SI1S0", "fuera"},
	{"ir", "SI3S0", "fuera"},
	{"ir", "SI2P0", "fuerais"},
	{"ir", "SI3P0", "fueran"},
	{"ir", "SI2S0", "fueras"},
	{"ir", "SF1S0", "fuere"},
	{"ir", "SF3S0", "fuere"},
	{"ir", "SF2P0", "fuereis"},
	{"ir", "SF3P0", "fueren"},
	{"ir", "SF2S0", "fueres"},
	{"ir", "IS3P0", "fueron"},
	{"ir", "SI1S0", "fuese"},
	{"ir", "SI3S0", "fuese"},
	{"ir", "SI2P0", "fueseis"},
	{"ir", "SI3P0", "fuesen"},
	{"ir", "SI2S0", "fueses"},
	{"ir", "IS1S0", "fui"},
	{"ir", "SI1P0", "fuéramos"},
	{"ir", "SF1P0", "fuéremos"},
	{"ir", "SI1P0", "fuésemos"},
	{"ir", "II1S0", "iba"},
	{"ir", "II3S0", "iba"},
	{"ir", "II2P0", "ibais"},
	{"ir", "II3P0", "iban"},
	{"ir", "II2S0", "ibas"},
	{"ir", "IP3S0", "va"},
	{"ir", "IP2P0", "vais"},
	{"ir", "IP1P0", "vamos"},
	{"ir", "M01P0", "vamos"},
	{"ir", "IP3P0", "van"},
	{"ir", "IP2S0", "vas"},
	{"ir", "IP1S0", "voy"},
	{"ir", "G0000", "yendo"},
	{"ir", "II1P0", "íbamos"},
	{"ir", "M01P0", "vayamos"},
	{"ser", "II1S0", "era"},
	{"ser", "II3S0", "era"},
	{"ser", "II2P0", "erais"},
	{"ser", "II3P0", "eran"},
	{"ser", "II2S0", "eras"},
	{"ser", "IP2S0", "eres"},
	{"ser", "IP3S0", "es"},
	{"ser", "IS3S0", "fue"},
	{"ser", "SI1S0", "fuera"},
	{"ser", "SI3S0", "fuera"},
	{"ser", "SI2P0", "fuerais"},
	{"ser", "SI3P0", "fueran"},
	{"ser", "SI2S0", "fueras"},
	{"ser", "SF1S0", "fuere"},
	{"ser", "SF3S0", "fuere"},
	{"ser", "SF2P0", "fuereis"},
	{"ser", "SF3P0", "fueren"},
	{"ser", "SF2S0", "fueres"},
	{"ser", "IS3P0", "fueron"},
	{"ser", "SI1S0", "fuese"},
	{"ser", "SI3S0", "fuese"},
	{"ser", "SI2P0", "fueseis"},
	{"ser", "SI3P0", "fuesen"},
	{"ser", "SI2S0", "fueses"},
	{"ser", "IS1S0", "fui"},
	{"ser", "IS1P0", "fuimos"},
	{"ser", "IS2S0", "fuiste"},
	{"ser", "IS2P0", "fuisteis"},
	{"ser", "SI1P0", "fuéramos"},
	{"ser", "SF1P0", "fuéremos"},
	{"ser", "SI1P0", "fuésemos"},
	{"ser", "M03S0", "sea"},
	{"ser", "SP1S0", "sea"},
	{"ser", "SP3S0", "sea"},
	{"ser", "M01P0", "seamos"},
	{"ser", "SP1P0", "seamos"},
	{"ser", "M03P0", "sean"},
	{"ser", "SP3P0", "sean"},
	{"ser", "SP2S0", "seas"},
	{"ser", "SP2P0", "seáis"},
	{"ser", "IP2P0", "sois"},
	{"ser", "IP1P0", "somos"},
	{"ser", "IP3P0", "son"},
	{"ser", "IP1S0", "soy"},
	{"ser", "M02S0", "sé"},
	{"ser", "II1P0", "éramos"},
	{"estar", "IP1S0", "estoy"},
	{"haber", "IP3S0", "ha"},
	{"haber", "IP3P0", "han"},
	{"haber", "IP2S0", "has"},
	{"haber", "IP3S0", "hay"},
	{"haber", "IP1S0", "he"},
	{"haber", "IP1P0", "hemos"},
	{"ver", "IP2P0", "veis"},
	{"ver", "IS1S0", "vi"},
	{"ver", "IS3S0", "vio"},
	{"hacer", "IS3S0", "hizo"},
	{"dar", "IP2P0", "dais"},
	{"dar", "SP2P0", "deis"},
	{"dar", "IS1S0", "di"},
	{"dar", "IS3S0", "dio"},
	{"dar", "IP1S0", "doy"},
	{"dar", "M03S0", "dé"},
	{"dar", "SP1S0", "dé"},
	{"dar", "SP3S0", "dé"},
	{"saber", "IP1S0", "sé"},
}

var uniqueVerbsByInfinitiveTag27 = buildUniqueVerbByInfinitiveTag27()

func buildUniqueVerbByInfinitiveTag27() map[string][]UniqueVerb {
	uniqueVerbByInfinitiveTag27 := map[string][]UniqueVerb{}
	for _, row := range uniqueVerbs {
		key := row.infinitive + row.tag27
		uniqueVerbByInfinitiveTag27[key] =
			append(uniqueVerbByInfinitiveTag27[key], row)
	}
	return uniqueVerbByInfinitiveTag27
}

func findUniqueVerbs(infinitive, tag27 string) []UniqueVerb {
	return uniqueVerbsByInfinitiveTag27[infinitive+tag27]
}

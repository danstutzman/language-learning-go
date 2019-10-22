package freeling

var AR_MUTANTS = map[string]string{
	"andar":    "anduv",
	"dar":      "d",
	"desandar": "desanduv",
	"estar":    "estuv",
}

var AR_MUTANT_TAG27_TO_SUFFIXES = map[string][]string{
	"IS1P0": {"imos"},
	"IS1S0": {"e"},
	"IS2P0": {"isteis"},
	"IS2S0": {"iste"},
	"IS3P0": {"ieron"},
	"IS3S0": {"o"},
	"SF1P0": {"iéremos"},
	"SF1S0": {"iere"},
	"SF2P0": {"iereis"},
	"SF2S0": {"ieres"},
	"SF3P0": {"ieren"},
	"SF3S0": {"iere"},
	"SI1P0": {"iéramos", "iésemos"},
	"SI1S0": {"iera", "iese"},
	"SI2P0": {"ierais", "ieseis"},
	"SI2S0": {"ieras", "ieses"},
	"SI3P0": {"ieran", "iesen"},
	"SI3S0": {"iera", "iese"},
}

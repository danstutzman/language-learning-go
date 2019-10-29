package english

import (
	"fmt"
	"regexp"
)

const (
	PRES      = "VBP"
	PAST      = "VBD"
	PAST_PART = "VBN"
	PRES_S    = "VBZ"
	GERUND    = "VBG"
)

var endsWithLongVowelConsonant = regexp.MustCompile(`([uao]m[pb]|[oa]wn|ey|elp|[ei]gn|ilm|o[uo]r|[oa]ugh|igh|ki|ff|oubt|ount|awl|o[alo]d|[iu]rl|upt|[oa]y|ight|oid|empt|act|aud|e[ea]d|ound|[aeiou][srcln]t|ept|dd|[eia]n[dk]|[ioa][xk]|[oa]rm|[ue]rn|[ao]ng|uin|eam|ai[mr]|[oea]w|[eaoui][rscl]k|[oa]r[nd]|ear|er|[^aieouyfm]it|[aeiouy]ir|[^aieouyfm]et|ll|en|vil|om|[^rno].mit|rup|bat|val|.[^skxwb][rvmchslwngb]el)$`)

var endsWithConsonantY = regexp.MustCompile(`[^aeiou]y$`)

var endsWithConsonantE = regexp.MustCompile(`[^aeiouy]e$`)

var endsWithShortVowelConsonant = regexp.MustCompile(`([^dtaeiuo][aeiuo][ptlgnm]|ir|cur|ug|[hj]ar|[^aouiey]ep|[^aeiuo][oua][db])$`)

var endsWithSibilant = regexp.MustCompile(`([ieao]ss|[aeiouy]zz|[aeiouy]ch|nch|rch|[aeiouy]sh|[iae]tch|ax)$`)

var endsWithEE = regexp.MustCompile(`(ee)$`)

var endsWithIE = regexp.MustCompile(`(ie)$`)

var endsWithUE = regexp.MustCompile(`(ue)$`)

var irregulars = map[string][4]string{
	"abhor":         [4]string{"abhorred", "abhorred", "abhors", "abhorring"},
	"abide":         [4]string{"abode", "abode", "abides", "abiding"},
	"acquit":        [4]string{"acquitted", "acquitted", "acquits", "acquitting"},
	"admit":         [4]string{"admitted", "admitted", "admits", "admitting"},
	"affix":         [4]string{"affixed", "affixed", "affixes", "affixing"},
	"apparel":       [4]string{"apparelled", "apparelled", "apparels", "apparelling"},
	"arise":         [4]string{"arose", "arisen", "arises", "arising"},
	"aver":          [4]string{"averred", "averred", "avers", "averring"},
	"awake":         [4]string{"awoke", "awoken", "awakes", "awaking"},
	"babysit":       [4]string{"babysat", "babysat", "babysits", "babysitting"},
	"backbite":      [4]string{"backbit", "backbit", "backbites", "backbiting"},
	"backslide":     [4]string{"backslid", "backslid", "backslides", "backsliding"},
	"bat":           [4]string{"batted", "batted", "bats", "batting"},
	"be":            [4]string{"was", "been", "is", "being"},
	"bear":          [4]string{"bore", "borne", "bears", "bearing"},
	"beat":          [4]string{"beat", "beaten", "beats", "beating"},
	"beckon":        [4]string{"beckoned", "beckoned", "beckons", "beckoning"},
	"become":        [4]string{"became", "become", "becomes", "becoming"},
	"bedim":         [4]string{"bedimmed", "bedimmed", "bedims", "bedimming"},
	"befall":        [4]string{"befell", "befallen", "befalls", "befalling"},
	"begin":         [4]string{"began", "begun", "begins", "beginning"},
	"behold":        [4]string{"beheld", "beheld", "beholds", "beholding"},
	"belch":         [4]string{"belched", "belched", "belches", "belching"},
	"bend":          [4]string{"bent", "bent", "bends", "bending"},
	"benefit":       [4]string{"benefited", "benefited", "benefits", "benefiting"},
	"bereave":       [4]string{"bereft", "bereft", "bereaves", "bereaving"},
	"beseech":       [4]string{"besought", "besought", "beseeches", "beseeching"},
	"bet":           [4]string{"bet", "bet", "bets", "betting"},
	"bias":          [4]string{"biased", "biased", "biases", "biasing"},
	"bib":           [4]string{"bibbed", "bibbed", "bibs", "bibbing"},
	"bid":           [4]string{"bid", "bid", "bids", "bidding"},
	"billet":        [4]string{"billetted", "billetted", "billets", "billetting"},
	"bind":          [4]string{"bound", "bound", "binds", "binding"},
	"bite":          [4]string{"bit", "bitten", "bites", "biting"},
	"blazon":        [4]string{"blazoned", "blazoned", "blazons", "blazoning"},
	"bleed":         [4]string{"bled", "bled", "bleeds", "bleeding"},
	"blow":          [4]string{"blew", "blown", "blows", "blowing"},
	"blur":          [4]string{"blurred", "blurred", "blurs", "blurring"},
	"bobsled":       [4]string{"bobsledded", "bobsledded", "bobsleds", "bobsledding"},
	"bollix":        [4]string{"bollixed", "bollixed", "bollixes", "bollixing"},
	"box":           [4]string{"boxed", "boxed", "boxes", "boxing"},
	"break":         [4]string{"broke", "broken", "breaks", "breaking"},
	"breastfeed":    [4]string{"breastfed", "breastfed", "breastfeeds", "breastfeeding"},
	"breed":         [4]string{"bred", "bred", "breeds", "breeding"},
	"bring":         [4]string{"brought", "brought", "brings", "bringing"},
	"broadcast":     [4]string{"broadcast", "broadcast", "broadcasts", "broadcasting"},
	"browbeat":      [4]string{"browbeat", "browbeat", "browbeats", "browbeating"},
	"buffet":        [4]string{"buffeted", "buffeted", "buffets", "buffeting"},
	"build":         [4]string{"built", "built", "builds", "building"},
	"burn":          [4]string{"burnt", "burnt", "burns", "burning"},
	"burst":         [4]string{"burst", "burst", "bursts", "bursting"},
	"buss":          [4]string{"bussed", "bussed", "busses", "bussing"},
	"bust":          [4]string{"bust", "bust", "busts", "busting"},
	"buy":           [4]string{"bought", "bought", "buys", "buying"},
	"callous":       [4]string{"calloused", "calloused", "callouses", "callousing"},
	"can":           [4]string{"could", "could", "can", "can"},
	"canal":         [4]string{"canaled", "canaled", "canals", "canaling"},
	"cancel":        [4]string{"cancelled", "cancelled", "cancels", "cancelling"},
	"caparison":     [4]string{"caparisoned", "caparisoned", "caparisons", "caparisoning"},
	"cast":          [4]string{"cast", "cast", "casts", "casting"},
	"catalog":       [4]string{"cataloged", "cataloged", "catalogs", "cataloging"},
	"catch":         [4]string{"caught", "caught", "catches", "catching"},
	"cavil":         [4]string{"cavilled", "cavilled", "cavils", "cavilling"},
	"chagrin":       [4]string{"chagrined", "chagrined", "chagrins", "chagrining"},
	"chairman":      [4]string{"chairmaned", "chairmaned", "chairmans", "chairmaning"},
	"channel":       [4]string{"channelled", "channelled", "channels", "channelling"},
	"char":          [4]string{"chared", "chared", "chars", "charing"},
	"chide":         [4]string{"chid", "chidden", "chides", "chiding"},
	"chirrup":       [4]string{"chirrupped", "chirrupped", "chirrups", "chirrupping"},
	"chisel":        [4]string{"chiselled", "chiselled", "chisels", "chiselling"},
	"choir":         [4]string{"choirred", "choirred", "choirs", "choirring"},
	"choose":        [4]string{"chose", "chosen", "chooses", "choosing"},
	"chorus":        [4]string{"chorused", "chorused", "choruses", "chorusing"},
	"cleave":        [4]string{"cleft", "cleft", "cleaves", "cleaving"},
	"climax":        [4]string{"climaxed", "climaxed", "climaxes", "climaxing"},
	"cling":         [4]string{"clung", "clung", "clings", "clinging"},
	"clothe":        [4]string{"clad", "clad", "clothes", "clothing"},
	"clutch":        [4]string{"clutched", "clutched", "clutches", "clutching"},
	"coax":          [4]string{"coaxed", "coaxed", "coaxes", "coaxing"},
	"coif":          [4]string{"coiffed", "coiffed", "coifs", "coiffing"},
	"combat":        [4]string{"combatted", "combatted", "combats", "combatting"},
	"come":          [4]string{"came", "come", "comes", "coming"},
	"comfit":        [4]string{"comfited", "comfited", "comfits", "comfiting"},
	"commix":        [4]string{"commixed", "commixed", "commixes", "commixing"},
	"confer":        [4]string{"conferred", "conferred", "confers", "conferring"},
	"cope":          [4]string{"coped", "cope", "copes", "coping"},
	"cosset":        [4]string{"cossetted", "cossetted", "cossets", "cossetting"},
	"cost":          [4]string{"cost", "cost", "costs", "costing"},
	"counsel":       [4]string{"counselled", "counselled", "counsels", "counselling"},
	"creep":         [4]string{"crept", "crept", "creeps", "creeping"},
	"crib":          [4]string{"cribbed", "cribbed", "cribs", "cribbing"},
	"crochet":       [4]string{"crochetted", "crochetted", "crochets", "crochetting"},
	"crossbreed":    [4]string{"crossbred", "crossbred", "crossbreeds", "crossbreeding"},
	"cuss":          [4]string{"cussed", "cussed", "cusses", "cussing"},
	"custom-make":   [4]string{"custom-made", "custom-made", "custom-makes", "custom-making"},
	"cut":           [4]string{"cut", "cut", "cuts", "cutting"},
	"daydream":      [4]string{"daydreamt", "daydreamt", "daydreams", "daydreaming"},
	"deal":          [4]string{"dealt", "dealt", "deals", "dealing"},
	"debar":         [4]string{"debarred", "debarred", "debars", "debarring"},
	"defer":         [4]string{"deferred", "deferred", "defers", "deferring"},
	"demur":         [4]string{"demurred", "demurred", "demurs", "demurring"},
	"desex":         [4]string{"desexed", "desexed", "desexes", "desexing"},
	"deter":         [4]string{"deterred", "deterred", "deters", "deterring"},
	"develop":       [4]string{"developed", "developed", "develops", "developing"},
	"devil":         [4]string{"devilled", "devilled", "devils", "devilling"},
	"dig":           [4]string{"dug", "dug", "digs", "digging"},
	"dim":           [4]string{"dimmed", "dimmed", "dims", "dimming"},
	"dip":           [4]string{"dipped", "dipped", "dips", "dipping"},
	"discomfit":     [4]string{"discomfited", "discomfited", "discomfits", "discomfiting"},
	"discuss":       [4]string{"discussed", "discussed", "discusses", "discussing"},
	"disembowel":    [4]string{"disembowelled", "disembowelled", "disembowels", "disembowelling"},
	"dishevel":      [4]string{"dishevelled", "dishevelled", "dishevels", "dishevelling"},
	"disprove":      [4]string{"disproved", "disproven", "disproves", "disproving"},
	"do":            [4]string{"did", "done", "does", "doing"},
	"dog":           [4]string{"dogged", "dogged", "dogs", "dogging"},
	"draw":          [4]string{"drew", "drawn", "draws", "drawing"},
	"dream":         [4]string{"dreamt", "dreamt", "dreams", "dreaming"},
	"drink":         [4]string{"drank", "drunk", "drinks", "drinking"},
	"drive":         [4]string{"drove", "driven", "drives", "driving"},
	"drivel":        [4]string{"drivelled", "drivelled", "drivels", "drivelling"},
	"dun":           [4]string{"dunned", "dunned", "duns", "dunning"},
	"dwell":         [4]string{"dwelt", "dwelt", "dwells", "dwelling"},
	"dye":           [4]string{"dyed", "dyed", "dyes", "dyeing"},
	"eat":           [4]string{"ate", "eaten", "eats", "eating"},
	"emblazon":      [4]string{"emblazoned", "emblazoned", "emblazons", "emblazonning"},
	"emit":          [4]string{"emited", "emited", "emits", "emiting"},
	"empanel":       [4]string{"empanelled", "empanelled", "empanels", "empanelling"},
	"endanger":      [4]string{"endangered", "endangered", "endangers", "entangling"},
	"envenom":       [4]string{"envenommed", "envenommed", "envenoms", "envenomming"},
	"equip":         [4]string{"equipped", "equipped", "equips", "equipping"},
	"ex":            [4]string{"exed", "exed", "exes", "exing"},
	"extol":         [4]string{"extolled", "extolled", "extols", "extolling"},
	"eye":           [4]string{"eyed", "eyed", "eyes", "eyeing"},
	"fall":          [4]string{"fell", "fallen", "falls", "falling"},
	"fathom":        [4]string{"fathommed", "fathommed", "fathoms", "fathomming"},
	"fax":           [4]string{"faxed", "faxed", "faxes", "faxing"},
	"feed":          [4]string{"fed", "fed", "feeds", "feeding"},
	"feel":          [4]string{"felt", "felt", "feels", "feeling"},
	"fight":         [4]string{"fought", "fought", "fights", "fighting"},
	"filch":         [4]string{"filched", "filched", "filches", "filching"},
	"filet":         [4]string{"filetted", "filetted", "filets", "filetting"},
	"fillet":        [4]string{"filletted", "filletted", "fillets", "filletting"},
	"find":          [4]string{"found", "found", "finds", "finding"},
	"fit":           [4]string{"fit", "fit", "fits", "fitting"},
	"fix":           [4]string{"fixed", "fixed", "fixes", "fixing"},
	"flee":          [4]string{"fled", "fled", "flees", "fleeing"},
	"flex":          [4]string{"flexed", "flexed", "flexes", "flexing"},
	"fling":         [4]string{"flung", "flung", "flings", "flinging"},
	"flit":          [4]string{"flitted", "flitted", "flits", "flitting"},
	"flummox":       [4]string{"flummoxed", "flummoxed", "flummoxes", "flummoxing"},
	"flux":          [4]string{"fluxed", "fluxed", "fluxes", "fluxing"},
	"fly":           [4]string{"flew", "flown", "flies", "flying"},
	"focus":         [4]string{"focused", "focused", "focuses", "focusing"},
	"forbid":        [4]string{"forbade", "forbidden", "forbids", "forbidding"},
	"forecast":      [4]string{"forecast", "forecast", "forecasts", "forecasting"},
	"foreknow":      [4]string{"foreknew", "foreknew", "foreknows", "foreknowing"},
	"foresee":       [4]string{"foresaw", "foreseen", "foresees", "foreseeing"},
	"foretell":      [4]string{"foretold", "foretold", "foretells", "foretelling"},
	"forget":        [4]string{"forgot", "forgotten", "forgets", "forgetting"},
	"forgive":       [4]string{"forgave", "forgiven", "forgives", "forgiving"},
	"forgo":         [4]string{"forwent", "forwent", "forgos", "forgoing"},
	"forsake":       [4]string{"forsook", "forsaken", "forsakes", "forsaking"},
	"forswear":      [4]string{"forsworn", "forsworn", "forswears", "forswearing"},
	"fox":           [4]string{"foxed", "foxed", "foxes", "foxing"},
	"freeze":        [4]string{"froze", "frozen", "freezes", "freezing"},
	"fret":          [4]string{"fretted", "fretted", "frets", "fretting"},
	"frostbite":     [4]string{"frostbit", "frostbitten", "frostbites", "frostbiting"},
	"fuss":          [4]string{"fussed", "fussed", "fusses", "fussing"},
	"gainsay":       [4]string{"gainsaid", "gainsaid", "gainsays", "gainsaying"},
	"gas":           [4]string{"gassed", "gassed", "gases", "gassing"},
	"get":           [4]string{"got", "got", "gets", "getting"},
	"gibbet":        [4]string{"gibbetted", "gibbetted", "gibbets", "gibbetting"},
	"give":          [4]string{"gave", "given", "gives", "giving"},
	"glom":          [4]string{"glommed", "glommed", "gloms", "glomming"},
	"go":            [4]string{"went", "gone", "goes", "going"},
	"gravel":        [4]string{"gravelled", "gravelled", "gravels", "gravelling"},
	"grind":         [4]string{"ground", "ground", "grinds", "grinding"},
	"grok":          [4]string{"grokked", "grokked", "groks", "grokking"},
	"grovel":        [4]string{"grovelled", "grovelled", "grovels", "grovelling"},
	"grow":          [4]string{"grew", "grown", "grows", "growing"},
	"hand-feed":     [4]string{"hand-fed", "hand-fed", "hand-feeds", "hand-feeding"},
	"handwrite":     [4]string{"handwrote", "handwritten", "handwrites", "handwriting"},
	"hang":          [4]string{"hung", "hung", "hangs", "hanging"},
	"have":          [4]string{"had", "had", "has", "having"},
	"hear":          [4]string{"heard", "heard", "hears", "hearing"},
	"heave":         [4]string{"hove", "hove", "heaves", "heaving"},
	"hew":           [4]string{"hewed", "hewn", "hews", "hewing"},
	"hex":           [4]string{"hexed", "hexed", "hexes", "hexing"},
	"hide":          [4]string{"hid", "hidden", "hides", "hiding"},
	"hit":           [4]string{"hit", "hit", "hits", "hitting"},
	"hoax":          [4]string{"hoaxed", "hoaxed", "hoaxes", "hoaxing"},
	"hold":          [4]string{"held", "held", "holds", "holding"},
	"hurt":          [4]string{"hurt", "hurt", "hurts", "hurting"},
	"immix":         [4]string{"immixed", "immixed", "immixes", "immixing"},
	"impanel":       [4]string{"impanelled", "impanelled", "impanels", "impanelling"},
	"imperil":       [4]string{"imperiled", "imperiled", "imperils", "imperiling"},
	"imprison":      [4]string{"imprisoned", "imprisoned", "imprisons", "imprisoning"},
	"inbreed":       [4]string{"inbred", "inbred", "inbreeds", "inbreeding"},
	"infer":         [4]string{"inferred", "inferred", "infers", "inferring"},
	"infix":         [4]string{"infixed", "infixed", "infixes", "infixing"},
	"inlay":         [4]string{"inlaid", "inlaid", "inlays", "inlaying"},
	"input":         [4]string{"input", "input", "inputs", "inputting"},
	"interbreed":    [4]string{"interbred", "interbred", "interbreeds", "interbreeding"},
	"intermix":      [4]string{"intermixed", "intermixed", "intermixes", "intermixing"},
	"interpret":     [4]string{"interpretted", "interpretted", "interprets", "interpretting"},
	"interweave":    [4]string{"interwove", "interwoven", "interweaves", "interweaving"},
	"interwind":     [4]string{"interwound", "interwound", "interwinds", "interwinding"},
	"intromit":      [4]string{"intromited", "intromited", "intromits", "intromiting"},
	"iron":          [4]string{"ironed", "ironed", "irons", "ironing"},
	"jar":           [4]string{"jared", "jared", "jars", "jaring"},
	"jerry-build":   [4]string{"jerry-built", "jerry-built", "jerry-builds", "jerry-building"},
	"jet":           [4]string{"jetted", "jetted", "jets", "jetting"},
	"jewel":         [4]string{"jewelled", "jewelled", "jewels", "jewelling"},
	"jinx":          [4]string{"jinxed", "jinxed", "jinxes", "jinxing"},
	"junket":        [4]string{"junketted", "junketted", "junkets", "junketting"},
	"keep":          [4]string{"kept", "kept", "keeps", "keeping"},
	"kid":           [4]string{"kidded", "kidded", "kids", "kidding"},
	"kneel":         [4]string{"knelt", "knelt", "kneels", "kneeling"},
	"knit":          [4]string{"knit", "knit", "knits", "knitting"},
	"know":          [4]string{"knew", "known", "knows", "knowing"},
	"label":         [4]string{"labelled", "labelled", "labels", "labelling"},
	"lade":          [4]string{"laded", "laden", "lades", "lading"},
	"larrup":        [4]string{"larrupped", "larrupped", "larrups", "larrupping"},
	"lay":           [4]string{"laid", "laid", "lays", "laying"},
	"lead":          [4]string{"led", "led", "leads", "leading"},
	"lean":          [4]string{"leant", "leant", "leans", "leaning"},
	"leap":          [4]string{"leapt", "leapt", "leaps", "leaping"},
	"learn":         [4]string{"learnt", "learnt", "learns", "learning"},
	"leave":         [4]string{"left", "left", "leaves", "leaving"},
	"lend":          [4]string{"lent", "lent", "lends", "lending"},
	"let":           [4]string{"let", "let", "lets", "letting"},
	"level":         [4]string{"levelled", "levelled", "levels", "levelling"},
	"lie":           [4]string{"lay", "lain", "lies", "laying"},
	"light":         [4]string{"lit", "lit", "lights", "lighting"},
	"lip-read":      [4]string{"lip-read", "lip-read", "lip-reads", "lip-reading"},
	"lose":          [4]string{"lost", "lost", "loses", "losing"},
	"make":          [4]string{"made", "made", "makes", "making"},
	"manumit":       [4]string{"manumited", "manumited", "manumits", "manumiting"},
	"mar":           [4]string{"marred", "marred", "mars", "marring"},
	"market":        [4]string{"marketted", "marketted", "markets", "marketting"},
	"marvel":        [4]string{"marvelled", "marvelled", "marvels", "marvelling"},
	"may":           [4]string{"might", "might", "may", "may"},
	"mean":          [4]string{"meant", "meant", "means", "meaning"},
	"meet":          [4]string{"met", "met", "meets", "meeting"},
	"miscast":       [4]string{"miscast", "miscast", "miscasts", "miscasting"},
	"misdeal":       [4]string{"misdealt", "misdealt", "misdeals", "misdealing"},
	"misdo":         [4]string{"misdid", "misdone", "misdoes", "misdoing"},
	"mishear":       [4]string{"misheard", "misheard", "mishears", "mishearing"},
	"misinterpret":  [4]string{"misinterpretted", "misinterpretted", "misinterprets", "misinterpretting"},
	"mislay":        [4]string{"mislaid", "mislaid", "mislays", "mislaying"},
	"mislead":       [4]string{"misled", "misled", "misleads", "misleading"},
	"mislearn":      [4]string{"mislearnt", "mislearnt", "mislearns", "mislearning"},
	"misread":       [4]string{"misread", "misread", "misreads", "misreading"},
	"misset":        [4]string{"misset", "misset", "missets", "missetting"},
	"misspeak":      [4]string{"misspoke", "misspoken", "misspeaks", "misspeaking"},
	"misspell":      [4]string{"misspelt", "misspelt", "misspells", "misspelling"},
	"misspend":      [4]string{"misspent", "misspent", "misspends", "misspending"},
	"mistake":       [4]string{"mistook", "mistaken", "mistakes", "mistaking"},
	"misteach":      [4]string{"mistaught", "mistaught", "misteaches", "misteaching"},
	"misunderstand": [4]string{"misunderstood", "misunderstood", "misunderstands", "misunderstanding"},
	"miswrite":      [4]string{"miswrote", "miswritten", "miswrites", "miswriting"},
	"mix":           [4]string{"mixed", "mixed", "mixes", "mixing"},
	"mow":           [4]string{"mowed", "mown", "mows", "mowing"},
	"muss":          [4]string{"mussed", "mussed", "musses", "mussing"},
	"must":          [4]string{"must", "must", "must", "must"},
	"net":           [4]string{"netted", "netted", "nets", "netting"},
	"nix":           [4]string{"nixed", "nixed", "nixes", "nixing"},
	"nonplus":       [4]string{"nonplused", "nonplused", "nonpluses", "nonplusing"},
	"offset":        [4]string{"offset", "offset", "offsets", "offsetting"},
	"outbid":        [4]string{"outbid", "outbid", "outbids", "outbidding"},
	"outdo":         [4]string{"outdid", "outdone", "outdoes", "outdoing"},
	"outdraw":       [4]string{"outdrew", "outdrawn", "outdraws", "outdrawing"},
	"outfight":      [4]string{"outfought", "outfought", "outfights", "outfighting"},
	"outfox":        [4]string{"outfoxed", "outfoxed", "outfoxes", "outfoxing"},
	"outgrow":       [4]string{"outgrew", "outgrown", "outgrows", "outgrowing"},
	"output":        [4]string{"output", "output", "outputs", "outputting"},
	"outrun":        [4]string{"outran", "outrun", "outran", "outrunning"},
	"outsell":       [4]string{"outsold", "outsold", "outsells", "outselling"},
	"outshine":      [4]string{"outshone", "outshone", "outshines", "outshining"},
	"outspend":      [4]string{"outspent", "outspent", "outspends", "outspending"},
	"outwit":        [4]string{"outwitted", "outwitted", "outwits", "outwitting"},
	"overbid":       [4]string{"overbid", "overbid", "overbids", "overbidding"},
	"overbuild":     [4]string{"overbuilt", "overbuilt", "overbuilds", "overbuilding"},
	"overbuy":       [4]string{"overbought", "overbought", "overbuys", "overbuying"},
	"overcome":      [4]string{"overcame", "overcome", "overcomes", "overcoming"},
	"overdo":        [4]string{"overdid", "overdone", "overdoes", "overdoing"},
	"overdraw":      [4]string{"overdrew", "overdrawn", "overdraws", "overdrawing"},
	"overdrive":     [4]string{"overdrove", "overdrove", "overdrives", "overdriving"},
	"overeat":       [4]string{"overate", "overeaten", "overeats", "overeating"},
	"overfeed":      [4]string{"overfed", "overfed", "overfeeds", "overfeeding"},
	"overhang":      [4]string{"overhung", "overhung", "overhangs", "overhanging"},
	"overhear":      [4]string{"overheard", "overheard", "overhears", "overhearing"},
	"overlay":       [4]string{"overlaid", "overlaid", "overlays", "overlaying"},
	"overlie":       [4]string{"overlaid", "overlaid", "overlies", "overlying"},
	"overpay":       [4]string{"overpaid", "overpaid", "overpays", "overpaying"},
	"override":      [4]string{"overrode", "overridden", "overrides", "overriding"},
	"overrun":       [4]string{"overran", "overrun", "overruns", "overrunning"},
	"oversee":       [4]string{"oversaw", "overseen", "oversees", "overseeing"},
	"oversell":      [4]string{"oversold", "oversold", "oversells", "overselling"},
	"overshoot":     [4]string{"overshot", "overshot", "overshots", "overshooting"},
	"oversleep":     [4]string{"overslept", "overslept", "oversleeps", "oversleeping"},
	"overspeak":     [4]string{"overspoke", "overspoken", "overspeaks", "overspeaking"},
	"overspend":     [4]string{"overspent", "overspent", "overspends", "overspending"},
	"overspill":     [4]string{"overspilt", "overspilt", "overspills", "overspilling"},
	"overspread":    [4]string{"overspread", "overspread", "overspreads", "overspreading"},
	"overstep":      [4]string{"oversteped", "oversteped", "oversteps", "oversteping"},
	"overtake":      [4]string{"overtook", "overtaken", "overtakes", "overtaking"},
	"overthink":     [4]string{"overthought", "overthought", "overthinks", "overthinking"},
	"overthrow":     [4]string{"overthrew", "overthrown", "overthrows", "overthrowing"},
	"overwind":      [4]string{"overwound", "overwound", "overwind", "overwinding"},
	"overwrite":     [4]string{"overwrote", "overwritten", "overwrite", "overwriting"},
	"panel":         [4]string{"panelled", "panelled", "panels", "panelling"},
	"parallel":      [4]string{"parallelled", "parallelled", "parallels", "parallelling"},
	"partake":       [4]string{"partook", "partaken", "partakes", "partaking"},
	"pay":           [4]string{"paid", "paid", "pays", "paying"},
	"peril":         [4]string{"periled", "periled", "perils", "periling"},
	"permit":        [4]string{"permitted", "permitted", "permits", "permitting"},
	"perplex":       [4]string{"perplexed", "perplexed", "perplexes", "perplexing"},
	"pilot":         [4]string{"piloted", "piloted", "pilots", "piloting"},
	"pit":           [4]string{"pitted", "pitted", "pits", "pitting"},
	"pivot":         [4]string{"pivoted", "pivoted", "pivots", "pivoting"},
	"plead":         [4]string{"pled", "pled", "pleads", "pleading"},
	"plummet":       [4]string{"plummeted", "plummeted", "plummets", "plummeting"},
	"pocket":        [4]string{"pocketted", "pocketted", "pockets", "pocketting"},
	"poison":        [4]string{"poisoned", "poisoned", "poisons", "poisoning"},
	"pommel":        [4]string{"pommelled", "pommelled", "pommels", "pommelling"},
	"prebuild":      [4]string{"prebuilt", "prebuilt", "prebuils", "prebuilding"},
	"prefer":        [4]string{"preferred", "preferred", "prefers", "preferring"},
	"preset":        [4]string{"preset", "preset", "presets", "presetting"},
	"preshrink":     [4]string{"preshrank", "preshrunk", "preshrinks", "preshrinking"},
	"profit":        [4]string{"profited", "profited", "profits", "profitting"},
	"proofread":     [4]string{"proofread", "proofread", "proofreads", "proofreading"},
	"pummel":        [4]string{"pummelled", "pummelled", "pummels", "pummelling"},
	"put":           [4]string{"put", "put", "puts", "putting"},
	"quarrel":       [4]string{"quarrelled", "quarrelled", "quarrels", "quarrelling"},
	"quick-freeze":  [4]string{"quick-froze", "quick-frozen", "quick-freezes", "quick-freezing"},
	"quit":          [4]string{"quit", "quit", "quits", "quitting"},
	"racket":        [4]string{"racketted", "racketted", "rackets", "racketting"},
	"ransom":        [4]string{"ransommed", "ransommed", "ransoms", "ransomming"},
	"ravel":         [4]string{"ravelled", "ravelled", "ravels", "ravelling"},
	"read":          [4]string{"read", "read", "reads", "reading"},
	"reason":        [4]string{"reasoned", "reasoned", "reasons", "reasoning"},
	"rebel":         [4]string{"rebelled", "rebelled", "rebels", "rebelling"},
	"rebid":         [4]string{"rebid", "rebid", "rebids", "rebidding"},
	"rebind":        [4]string{"rebound", "rebound", "rebinds", "rebinding"},
	"rebuild":       [4]string{"rebuilt", "rebuilt", "rebuilds", "rebuilding"},
	"recast":        [4]string{"recast", "recast", "recasts", "recasting"},
	"redo":          [4]string{"redid", "redone", "redoes", "redoing"},
	"redraw":        [4]string{"redrew", "redrawn", "redraws", "redrawing"},
	"refer":         [4]string{"referred", "referred", "refers", "referring"},
	"regret":        [4]string{"regretted", "regretted", "regrets", "regretting"},
	"regrind":       [4]string{"reground", "reground", "regrinds", "regrinding"},
	"regrow":        [4]string{"regrew", "regrown", "regrows", "regrowing"},
	"rehang":        [4]string{"rehung", "rehung", "rehangs", "rehanging"},
	"rehear":        [4]string{"reheard", "reheard", "rehears", "rehearing"},
	"reknit":        [4]string{"reknit", "reknit", "reknits", "reknitting"},
	"relax":         [4]string{"relaxed", "relaxed", "relaxes", "relaxing"},
	"relay":         [4]string{"relaid", "relaid", "relays", "relaying"},
	"remake":        [4]string{"remade", "remade", "remakes", "remaking"},
	"remit":         [4]string{"remited", "remited", "remits", "remiting"},
	"rend":          [4]string{"rent", "rent", "rends", "rending"},
	"repay":         [4]string{"repaid", "repaid", "repays", "repaying"},
	"reread":        [4]string{"reread", "reread", "rereads", "rereading"},
	"rerun":         [4]string{"reran", "rerun", "reruns", "rerunning"},
	"resell":        [4]string{"resold", "resold", "resells", "reselling"},
	"resend":        [4]string{"resent", "resent", "resends", "resending"},
	"reset":         [4]string{"reset", "reset", "resets", "resetting"},
	"retake":        [4]string{"retook", "retaken", "retakes", "retaking"},
	"reteach":       [4]string{"retaught", "retaught", "reteaches", "reteaching"},
	"retell":        [4]string{"retold", "retold", "retells", "retelling"},
	"rethink":       [4]string{"rethought", "rethought", "rethinks", "rethinking"},
	"retread":       [4]string{"retread", "retread", "retreads", "retreading"},
	"retreat":       [4]string{"retreat", "retreat", "retreats", "retreating"},
	"retrofit":      [4]string{"retrofit", "retrofit", "retrofits", "retrofitting"},
	"retroflex":     [4]string{"retroflexed", "retroflexed", "retroflexes", "retroflexing"},
	"revel":         [4]string{"revelled", "revelled", "revels", "revelling"},
	"rewind":        [4]string{"rewound", "rewound", "rewinds", "rewinding"},
	"rib":           [4]string{"ribbed", "ribbed", "ribs", "ribbing"},
	"ricochet":      [4]string{"ricochetted", "ricochetted", "ricochets", "ricochetting"},
	"rid":           [4]string{"rid", "rid", "rids", "ridding"},
	"ride":          [4]string{"rode", "ridden", "rides", "riding"},
	"ring":          [4]string{"rang", "rung", "rings", "ringing"},
	"rise":          [4]string{"rose", "risen", "rises", "rising"},
	"rival":         [4]string{"rivalled", "rivalled", "rivals", "rivalling"},
	"rivet":         [4]string{"rivetted", "rivetted", "rivets", "rivetting"},
	"rocket":        [4]string{"rocketted", "rocketted", "rockets", "rocketting"},
	"run":           [4]string{"ran", "run", "runs", "running"},
	"sand-cast":     [4]string{"sand-cast", "sand-cast", "sand-casts", "sand-casting"},
	"saw":           [4]string{"sawed", "sawn", "saws", "sawing"},
	"say":           [4]string{"said", "said", "says", "saying"},
	"scar":          [4]string{"scarred", "scarred", "scars", "scarring"},
	"see":           [4]string{"saw", "seen", "sees", "seeing"},
	"seek":          [4]string{"sought", "sought", "seeks", "seeking"},
	"sell":          [4]string{"sold", "sold", "sells", "selling"},
	"send":          [4]string{"sent", "sent", "sends", "sending"},
	"set":           [4]string{"set", "set", "sets", "setting"},
	"sew":           [4]string{"sewed", "sewn", "sews", "sewing"},
	"sex":           [4]string{"sexed", "sexed", "sexes", "sexing"},
	"shake":         [4]string{"shook", "shaken", "shakes", "shaking"},
	"shall":         [4]string{"should", "should", "shall", "shall"},
	"shave":         [4]string{"shove", "shaven", "shaves", "shaving"},
	"shear":         [4]string{"shore", "shorn", "shears", "shearing"},
	"shed":          [4]string{"shed", "shed", "sheds", "shedding"},
	"shine":         [4]string{"shone", "shone", "shines", "shining"},
	"shit":          [4]string{"shit", "shit", "shits", "shitting"},
	"shoe":          [4]string{"shod", "shod", "shoes", "shoeing"},
	"shoot":         [4]string{"shot", "shot", "shoots", "shooting"},
	"show":          [4]string{"showed", "shown", "shows", "showing"},
	"shrink":        [4]string{"shrank", "shrunk", "shrinks", "shrinking"},
	"shrivel":       [4]string{"shrivelled", "shrivelled", "shrivels", "shrivelling"},
	"shut":          [4]string{"shut", "shut", "shuts", "shutting"},
	"sidestep":      [4]string{"sidesteped", "sidesteped", "sidesteps", "sidesteping"},
	"signal":        [4]string{"signaled", "signaled", "signals", "signaling"},
	"sing":          [4]string{"sang", "sung", "sings", "singing"},
	"sink":          [4]string{"sank", "sunk", "sinks", "sinking"},
	"sit":           [4]string{"sat", "sat", "sits", "sitting"},
	"skid":          [4]string{"skidded", "skidded", "skids", "skidding"},
	"slay":          [4]string{"slew", "slain", "slays", "slaying"},
	"sleep":         [4]string{"slept", "slept", "sleeps", "sleeping"},
	"slide":         [4]string{"slid", "slid", "slides", "sliding"},
	"sling":         [4]string{"slung", "slung", "slings", "slinging"},
	"slink":         [4]string{"slunk", "slunk", "slinks", "slinking"},
	"slit":          [4]string{"slit", "slit", "slits", "slitting"},
	"slur":          [4]string{"slurred", "slurred", "slurs", "slurring"},
	"smell":         [4]string{"smelt", "smelt", "smells", "smelling"},
	"smite":         [4]string{"smote", "smitten", "smites", "smiting"},
	"sneak":         [4]string{"snuck", "snuck", "sneaks", "sneaking"},
	"speak":         [4]string{"spoke", "spoken", "speaks", "speaking"},
	"speed":         [4]string{"sped", "sped", "speeds", "speeding"},
	"spell":         [4]string{"spelt", "spelt", "spells", "spelling"},
	"spend":         [4]string{"spent", "spent", "spends", "spending"},
	"spill":         [4]string{"spilt", "spilt", "spills", "spilling"},
	"spin":          [4]string{"spun", "spun", "spins", "spinning"},
	"spiral":        [4]string{"spiraled", "spiraled", "spirals", "spiraling"},
	"spit":          [4]string{"spat", "spat", "spits", "spitting"},
	"split":         [4]string{"split", "split", "splits", "splitting"},
	"spoil":         [4]string{"spoilt", "spoilt", "spoils", "spoiling"},
	"spread":        [4]string{"spread", "spread", "spreads", "spreading"},
	"spring":        [4]string{"sprang", "sprung", "springs", "springing"},
	"spur":          [4]string{"spurred", "spurred", "spurs", "spurring"},
	"squat":         [4]string{"squatted", "squatted", "squats", "squatting"},
	"stand":         [4]string{"stood", "stood", "stands", "standing"},
	"steal":         [4]string{"stole", "stolen", "steals", "stealing"},
	"stem":          [4]string{"stemmed", "stemmed", "stems", "stemming"},
	"stick":         [4]string{"stuck", "stuck", "sticks", "sticking"},
	"sting":         [4]string{"stung", "stung", "stings", "stinging"},
	"stink":         [4]string{"stank", "stunk", "stinks", "stinking"},
	"stop":          [4]string{"stopped", "stopped", "stops", "stopping"},
	"strew":         [4]string{"strewed", "strewn", "strews", "strewing"},
	"stride":        [4]string{"strode", "stridden", "strides", "striding"},
	"strike":        [4]string{"struck", "struck", "strikes", "striking"},
	"string":        [4]string{"strung", "strung", "strings", "stringing"},
	"strive":        [4]string{"strove", "striven", "strives", "striving"},
	"stymie":        [4]string{"stymied", "stymied", "stymies", "stymieing"},
	"stymy":         [4]string{"stymied", "stymied", "stymies", "stymieing"},
	"sublet":        [4]string{"sublet", "sublet", "sublets", "subletting"},
	"submit":        [4]string{"submitted", "submitted", "submits", "submitting"},
	"suds":          [4]string{"sudsed", "sudsed", "sudses", "sudsing"},
	"summon":        [4]string{"summoned", "summoned", "summons", "summoning"},
	"swear":         [4]string{"swore", "sworn", "swears", "swearing"},
	"sweat":         [4]string{"sweat", "sweat", "sweats", "sweating"},
	"sweep":         [4]string{"swept", "swept", "sweeps", "sweeping"},
	"swell":         [4]string{"swelled", "swollen", "swells", "swelling"},
	"swim":          [4]string{"swam", "swum", "swims", "swimming"},
	"swing":         [4]string{"swung", "swung", "swings", "swinging"},
	"tag":           [4]string{"tagged", "tagged", "tags", "tagging"},
	"tailor-make":   [4]string{"tailor-made", "tailor-made", "tailor-makes", "tailor-making"},
	"take":          [4]string{"took", "taken", "takes", "taking"},
	"tan":           [4]string{"tanned", "tanned", "tans", "tanning"},
	"tap":           [4]string{"tapped", "tapped", "taps", "tapping"},
	"target":        [4]string{"targetted", "targetted", "targets", "targetting"},
	"tat":           [4]string{"tatted", "tatted", "tats", "tatting"},
	"tax":           [4]string{"taxed", "taxed", "taxes", "taxing"},
	"teach":         [4]string{"taught", "taught", "teaches", "teaching"},
	"tear":          [4]string{"tore", "torn", "tears", "tearing"},
	"telefax":       [4]string{"telefaxed", "telefaxed", "telefaxes", "telefaxing"},
	"tell":          [4]string{"told", "told", "tells", "telling"},
	"test-drive":    [4]string{"test-drove", "test-driven", "test-drives", "test-driving"},
	"test-fly":      [4]string{"test-flew", "test-flown", "test-flies", "test-flying"},
	"think":         [4]string{"thought", "thought", "thinks", "thinking"},
	"thrive":        [4]string{"throve", "thriven", "thrives", "thriving"},
	"throw":         [4]string{"threw", "thrown", "throws", "throwing"},
	"thrust":        [4]string{"thrust", "thrust", "thrusts", "thrusting"},
	"ticket":        [4]string{"ticketted", "ticketted", "tickets", "ticketting"},
	"tip":           [4]string{"tipped", "tipped", "tips", "tipping"},
	"toe":           [4]string{"toed", "toed", "toes", "toeing"},
	"tog":           [4]string{"togged", "togged", "togs", "toging"},
	"trammel":       [4]string{"trammelled", "trammelled", "trammels", "trammelling"},
	"transfer":      [4]string{"transferred", "transferred", "transfers", "transferring"},
	"transfix":      [4]string{"transfixed", "transfixed", "transfixes", "transfixing"},
	"travel":        [4]string{"travelled", "travelled", "travels", "travelling"},
	"tread":         [4]string{"trod", "trodden", "treads", "treading"},
	"tug":           [4]string{"tuged", "tuged", "tugs", "tuging"},
	"tunnel":        [4]string{"tunnelled", "tunnelled", "tunnels", "tunnelling"},
	"twit":          [4]string{"twitted", "twitted", "twits", "twitting"},
	"typecast":      [4]string{"typecast", "typecast", "typecasts", "typecasting"},
	"typeset":       [4]string{"typeset", "typeset", "typesets", "typesetting"},
	"typewrite":     [4]string{"typewrote", "typewritten", "typewrites", "typewriting"},
	"unbend":        [4]string{"unbent", "unbent", "unbends", "unbending"},
	"unbind":        [4]string{"unbound", "unbound", "unbinds", "unbinding"},
	"unbosom":       [4]string{"unbosommed", "unbosommed", "unbosoms", "unbosomming"},
	"unclothe":      [4]string{"unclad", "unclad", "unclothes", "unclothing"},
	"underbid":      [4]string{"underbid", "underbid", "underbids", "underbidding"},
	"undercut":      [4]string{"undercut", "undercut", "undercuts", "undercutting"},
	"underfeed":     [4]string{"underfed", "underfed", "underfeeds", "underfeeding"},
	"undergo":       [4]string{"underwent", "undergone", "undergoes", "undergoing"},
	"underlie":      [4]string{"underlay", "underlain", "underlies", "underlaying"},
	"undersell":     [4]string{"undersold", "undersold", "undersells", "underselling"},
	"understand":    [4]string{"understood", "understood", "understands", "understanding"},
	"undertake":     [4]string{"undertook", "undertaken", "undertakes", "undertaking"},
	"underwrite":    [4]string{"underwrote", "underwritten", "underwrites", "underwriting"},
	"undo":          [4]string{"undid", "undone", "undoes", "undoing"},
	"unfit":         [4]string{"unfited", "unfited", "unfits", "unfiting"},
	"unfreeze":      [4]string{"unfroze", "unfrozen", "unfreezes", "unfreezing"},
	"unknit":        [4]string{"unknit", "unknit", "unknits", "unknitting"},
	"unlax":         [4]string{"unlaxed", "unlaxed", "unlaxes", "unlaxing"},
	"unmake":        [4]string{"unmade", "unmade", "unmakes", "unmaking"},
	"unravel":       [4]string{"unravelled", "unravelled", "unravels", "unravelling"},
	"unsay":         [4]string{"unsaid", "unsaid", "unsays", "unsaying"},
	"unsex":         [4]string{"unsexed", "unsexed", "unsexes", "unsexing"},
	"unwind":        [4]string{"unwound", "unwound", "unwinds", "unwinding"},
	"uphold":        [4]string{"upheld", "upheld", "upholds", "upholding"},
	"upset":         [4]string{"upset", "upset", "upsets", "upsetting"},
	"vex":           [4]string{"vexed", "vexed", "vexes", "vexing"},
	"wake":          [4]string{"woke", "woken", "wakes", "waking"},
	"wallop":        [4]string{"walloped", "walloped", "wallops", "walloping"},
	"wax":           [4]string{"waxed", "waxed", "waxes", "waxing"},
	"waylay":        [4]string{"waylaid", "waylaid", "waylays", "waylaying"},
	"wear":          [4]string{"wore", "worn", "wears", "wearing"},
	"weave":         [4]string{"wove", "woven", "weaves", "weaving"},
	"web":           [4]string{"webbed", "webbed", "webs", "webbing"},
	"wed":           [4]string{"wed", "wed", "weds", "wedding"},
	"weep":          [4]string{"wept", "wept", "weeps", "weeping"},
	"wend":          [4]string{"went", "went", "wends", "wending"},
	"wet":           [4]string{"wet", "wet", "wets", "wetting"},
	"whet":          [4]string{"whetted", "whetted", "whets", "whetting"},
	"will":          [4]string{"would", "would", "will", "willing"},
	"win":           [4]string{"won", "won", "wins", "winning"},
	"wind":          [4]string{"wound", "wound", "winds", "winding"},
	"withdraw":      [4]string{"withdrew", "withdrawn", "withdraws", "withdrawing"},
	"withhold":      [4]string{"withheld", "withheld", "withholds", "withholding"},
	"withstand":     [4]string{"withstood", "withstood", "withstands", "withstanding"},
	"worship":       [4]string{"worshiped", "worshiped", "worships", "worshiping"},
	"wring":         [4]string{"wrung", "wrung", "wrings", "wringing"},
	"write":         [4]string{"wrote", "written", "writes", "writing"},
	"xerox":         [4]string{"xeroxed", "xeroxed", "xeroxes", "xeroxing"},
	"yak":           [4]string{"yakked", "yakked", "yaks", "yakking"},
	"yen":           [4]string{"yenned", "yenned", "yens", "yenning"},
	"zinc":          [4]string{"zincked", "zincked", "zincs", "zincking"},
}

func ConjugateVerb(vb, to string) string {
	conjugations, wasFound := irregulars[vb]
	if wasFound {
		if to == PRES_S {
			return conjugations[2]
		} else if to == GERUND {
			return conjugations[3]
		} else if to == PAST_PART {
			return conjugations[0]
		} else if to == PAST {
			return conjugations[1]
		} else {
			return vb
		}
	}

	if endsWithLongVowelConsonant.MatchString(vb) {
		if to == PRES_S {
			return vb + "s"
		} else if to == GERUND {
			return vb + "ing"
		} else if to == PAST_PART || to == PAST {
			return vb + "ed"
		} else {
			return vb
		}
	}

	if endsWithConsonantY.MatchString(vb) {
		base := vb[0 : len(vb)-1]
		if to == PRES_S {
			return base + "ies"
		}
		if to == GERUND {
			return vb + "ing"
		}
		if to == PAST_PART || to == PAST {
			return base + "ied"
		}
		return vb
	}

	if endsWithConsonantE.MatchString(vb) {
		base := vb[0 : len(vb)-1]
		if to == PRES_S {
			return vb + "s"
		}
		if to == GERUND {
			return base + "ing"
		}
		if to == PAST_PART || to == PAST {
			return base + "ed"
		}
		return vb
	}

	if endsWithShortVowelConsonant.MatchString(vb) {
		if to == PRES_S {
			return vb + "s"
		}
		if to == GERUND {
			return vb + vb[len(vb)-1:len(vb)] + "ing"
		}
		if to == PAST_PART || to == PAST {
			return vb + vb[len(vb)-1:len(vb)] + "ed"
		}
		return vb
	}

	if endsWithSibilant.MatchString(vb) {
		if to == PRES_S {
			return vb + "es"
		}
		if to == GERUND {
			return vb + "ing"
		}
		if to == PAST_PART || to == PAST {
			return vb + "ed"
		}
		return vb
	}

	if endsWithEE.MatchString(vb) {
		if to == PRES_S {
			return vb + "s"
		}
		if to == GERUND {
			return vb + "ing"
		}
		if to == PAST_PART || to == PAST {
			return vb + "d"
		}
		return vb
	}

	if endsWithIE.MatchString(vb) {
		if to == PRES_S {
			return vb + "s"
		}
		if to == GERUND {
			return vb[0:len(vb)-2] + "ying"
		}
		if to == PAST_PART || to == PAST {
			return vb + "d"
		}
		return vb
	}

	if endsWithUE.MatchString(vb) {
		if to == PRES_S {
			return vb + "s"
		}
		if to == GERUND {
			return vb[0:len(vb)-1] + "ing"
		}
		if to == PAST_PART || to == PAST {
			return vb + "d"
		}
		return vb
	}

	// Default
	if to == PRES_S {
		return vb + "s"
	}
	if to == GERUND {
		return vb + "ing"
	}
	if to == PAST_PART || to == PAST {
		return vb + "ed"
	}
	return vb
}

func main() {
	vb := "see"
	fmt.Println(ConjugateVerb(vb, PRES_S))
	fmt.Println(ConjugateVerb(vb, GERUND))
	fmt.Println(ConjugateVerb(vb, PAST_PART))
	fmt.Println(ConjugateVerb(vb, PAST))
	fmt.Println(ConjugateVerb(vb, PRES))
}

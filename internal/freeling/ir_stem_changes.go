package freeling

var IR_STEM_CHANGES = map[string]string{
	"acomedir":     "acomid",
	"adherir":      "adhier",
	"adquirir":     "adquier",
	"advenir":      "advien",
	"advertir":     "adviert",
	"afluir":       "afluy",
	"antedecir":    "antedic",
	"antevenir":    "antevien",
	"argüir":       "arguy",
	"arrepentir":   "arrepient",
	"asentir":      "asient",
	"astreñir":     "astriñ",
	"atribuir":     "atribuy",
	"avenir":       "avien",
	"bendecir":     "bendic",
	"cernir":       "ciern",
	"circunferir":  "circunfier",
	"ceñir":        "ciñ",
	"cohibir":      "cohíb",
	"colegir":      "colig",
	"comedir":      "comid",
	"competir":     "compit",
	"concebir":     "concib",
	"concernir":    "conciern",
	"concluir":     "concluy",
	"condecir":     "condic",
	"conferir":     "confier",
	"confluir":     "confluy",
	"consentir":    "consient",
	"conseguir":    "consigu",
	"constituir":   "constituy",
	"constreñir":   "constriñ",
	"construir":    "construy",
	"contradecir":  "contradic",
	"contravenir":  "contravien",
	"contribuir":   "contribuy",
	"controvertir": "controviert",
	"convenir":     "convien",
	"convertir":    "conviert",
	"corregir":     "corrig",
	"costreñir":    "costriñ",
	"deferir":      "defier",
	"derretir":     "derrit",
	"derruir":      "derruy",
	"desadvertir":  "desadviert",
	"desavenir":    "desavien",
	"desceñir":     "desciñ",
	"descomedir":   "descomid",
	"desconvenir":  "desconvien",
	"desdecir":     "desdic",
	"desmedir":     "desmid",
	"desmentir":    "desmient",
	"desobstruir":  "desobstruy",
	"despedir":     "despid",
	"destituir":    "destituy",
	"desteñir":     "destiñ",
	"destruir":     "destruy",
	"desvestir":    "desvist",
	"devenir":      "devien",
	"decir":        "dic",
	"diferir":      "difier",
	"difluir":      "difluy",
	"digerir":      "digier",
	"diluir":       "diluy",
	"diminuir":     "diminuy",
	"discernir":    "disciern",
	"disconvenir":  "disconvien",
	"disentir":     "disient",
	"disminuir":    "disminuy",
	"distribuir":   "distribuy",
	"divertir":     "diviert",
	"dormir":       "duerm",
	"elegir":       "elig",
	"embestir":     "embist",
	"engerir":      "engier",
	"entredecir":   "entredic",
	"estatuir":     "estatuy",
	"estreñir":     "estriñ",
	"excluir":      "excluy",
	"expedir":      "expid",
	"fluir":        "fluy",
	"fruir":        "fruy",
	"gemir":        "gim",
	"gruir":        "gruy",
	"hendir":       "hiend",
	"herir":        "hier",
	"hervir":       "hierv",
	"henchir":      "hinch",
	"heñir":        "hiñ",
	"huir":         "huy",
	"imbuir":       "imbuy",
	"impedir":      "impid",
	"incluir":      "incluy",
	"inferir":      "infier",
	"influir":      "influy",
	"ingerir":      "ingier",
	"injerir":      "injier",
	"inmiscuir":    "inmiscuy",
	"inquirir":     "inquier",
	"inserir":      "insier",
	"instituir":    "instituy",
	"instruir":     "instruy",
	"interferir":   "interfier",
	"intervenir":   "intervien",
	"intuir":       "intuy",
	"invertir":     "inviert",
	"investir":     "invist",
	"erguir":       "irgu",
	"luir":         "luy",
	"maldecir":     "maldic",
	"malherir":     "malhier",
	"medir":        "mid",
	"mentir":       "mient",
	"morir":        "muer",
	"obstruir":     "obstruy",
	"ocluir":       "ocluy",
	"perquirir":    "perquier",
	"perseguir":    "persigu",
	"pervertir":    "perviert",
	"pedir":        "pid",
	"preconcebir":  "preconcib",
	"predecir":     "predic",
	"preelegir":    "preelig",
	"preferir":     "prefier",
	"premorir":     "premuer",
	"presentir":    "presient",
	"prevenir":     "previen",
	"proferir":     "profier",
	"prohibir":     "prohíb",
	"proseguir":    "prosigu",
	"prostituir":   "prostituy",
	"provenir":     "provien",
	"readquirir":   "readquier",
	"reargüir":     "rearguy",
	"recluir":      "recluy",
	"reconstituir": "reconstituy",
	"reconstruir":  "reconstruy",
	"reconvenir":   "reconvien",
	"reconvertir":  "reconviert",
	"redargüir":    "redarguy",
	"redecir":      "redic",
	"redistribuir": "redistribuy",
	"reelegir":     "reelig",
	"reexpedir":    "reexpid",
	"referir":      "refier",
	"refluir":      "refluy",
	"rehervir":     "rehierv",
	"rehenchir":    "rehinch",
	"rehuir":       "rehuy",
	"remedir":      "remid",
	"repetir":      "repit",
	"requerir":     "requier",
	"resentir":     "resient",
	"reseguir":     "resigu",
	"restituir":    "restituy",
	"reteñir":      "retiñ",
	"retribuir":    "retribuy",
	"revenir":      "revien",
	"revertir":     "reviert",
	"revestir":     "revist",
	"reunir":       "reún",
	"regir":        "rig",
	"rendir":       "rind",
	"reñir":        "riñ",
	"sentir":       "sient",
	"seguir":       "sigu",
	"servir":       "sirv",
	"sobrevenir":   "sobrevien",
	"subseguir":    "subsigu",
	"substituir":   "substituy",
	"subvenir":     "subvien",
	"subvertir":    "subviert",
	"sugerir":      "sugier",
	"supervenir":   "supervien",
	"sustituir":    "sustituy",
	"teñir":        "tiñ",
	"transferir":   "transfier",
	"trasferir":    "trasfier",
	"travestir":    "travist",
	"tribuir":      "tribuy",
	"venir":        "vien",
	"vestir":       "vist",
	"zaherir":      "zahier",
}

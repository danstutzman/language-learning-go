package freeling

var AR_STEM_CHANGES = map[string][]string{
	"aberrar":          {"abierr"},
	"abuñolar":         {"abuñuel"},
	"acensuar":         {"acensú"},
	"acentuar":         {"acentú"},
	"acertar":          {"aciert"},
	"aclocar":          {"acluec"},
	"acrecentar":       {"acrecient"},
	"actuar":           {"actú"},
	"acollar":          {"acuell"},
	"acordar":          {"acuerd"},
	"acornar":          {"acuern"},
	"acostar":          {"acuest"},
	"adecuar":          {"adecú"},
	"afollar":          {"afuell"},
	"aforar":           {"afuer"},
	"agorar":           {"agüer"},
	"ahijar":           {"ahíj"},
	"ahilar":           {"ahíl"},
	"ahincar":          {"ahínc"},
	"ahitar":           {"ahít"},
	"ahuchar":          {"ahúch"},
	"ahumar":           {"ahúm"},
	"ahusar":           {"ahús"},
	"ajorar":           {"ajuer"},
	"alebrar":          {"aliebr"},
	"alentar":          {"alient"},
	"almorzar":         {"almuerz"},
	"aliar":            {"alí"},
	"amnistiar":        {"amnistí"},
	"amohinar":         {"amohín"},
	"ampliar":          {"amplí"},
	"amoblar":          {"amuebl"},
	"amolar":           {"amuel"},
	"aneblar":          {"aniebl"},
	"ansiar":           {"ansí"},
	"anticuar":         {"anticú"},
	"apacentar":        {"apacient"},
	"apercollar":       {"apercuell"},
	"apernar":          {"apiern"},
	"apretar":          {"apriet"},
	"apropincuar":      {"apropincú"},
	"aprobar":          {"aprueb"},
	"apostar":          {"apuest", "apost"},
	"arcaizar":         {"arcaíz"},
	"arrendar":         {"arriend"},
	"arriar":           {"arrí"},
	"arruar":           {"arrú"},
	"asentar":          {"asient"},
	"aserrar":          {"asierr"},
	"asestar":          {"asiest"},
	"aspaventar":       {"aspavient"},
	"asolar":           {"asuel", "asol"},
	"asonar":           {"asuen"},
	"ataviar":          {"ataví"},
	"atenuar":          {"atenú"},
	"aterrar":          {"atierr", "aterr"},
	"atesar":           {"aties"},
	"atestar":          {"atiest", "atest"},
	"atravesar":        {"atravies"},
	"atraillar":        {"atraíll"},
	"atronar":          {"atruen"},
	"atorar":           {"atuer", "ator"},
	"autografiar":      {"autografí"},
	"avaluar":          {"avalú"},
	"avergonzar":       {"avergüenz"},
	"averiar":          {"averí"},
	"aventar":          {"avient"},
	"aviar":            {"aví"},
	"azolar":           {"azuel"},
	"airar":            {"aír"},
	"aislar":           {"aísl"},
	"aullar":           {"aúll"},
	"aunar":            {"aún"},
	"aupar":            {"aúp"},
	"baquiar":          {"baquí"},
	"beldar":           {"bield"},
	"biografiar":       {"biografí"},
	"cablegrafiar":     {"cablegrafí"},
	"cabrahigar":       {"cabrahíg"},
	"calentar":         {"calient"},
	"caligrafiar":      {"caligrafí"},
	"calofriar":        {"calofrí"},
	"cartografiar":     {"cartografí"},
	"chirriar":         {"chirrí"},
	"cegar":            {"cieg"},
	"cerrar":           {"cierr"},
	"cimentar":         {"cimient"},
	"circunvolar":      {"circunvuel"},
	"clocar":           {"cluec"},
	"comenzar":         {"comienz"},
	"comprobar":        {"comprueb"},
	"conceptuar":       {"conceptú"},
	"concertar":        {"conciert"},
	"concordar":        {"concuerd"},
	"confesar":         {"confies"},
	"confiar":          {"confí"},
	"consolar":         {"consuel"},
	"contextuar":       {"contextú"},
	"continuar":        {"continú"},
	"contrariar":       {"contrarí"},
	"criar":            {"crí"},
	"cuantiar":         {"cuantí"},
	"cuchichiar":       {"cuchichí"},
	"colgar":           {"cuelg"},
	"colar":            {"cuel"},
	"contar":           {"cuent"},
	"costar":           {"cuest"},
	"ciar":             {"cí"},
	"dactilografiar":   {"dactilografí"},
	"decentar":         {"decient"},
	"degollar":         {"degüell"},
	"demostrar":        {"demuestr"},
	"denegar":          {"denieg"},
	"denostar":         {"denuest"},
	"derrengar":        {"derrieng"},
	"desacertar":       {"desaciert"},
	"desacordar":       {"desacuerd"},
	"desaforar":        {"desafuer"},
	"desafiar":         {"desafí"},
	"desahijar":        {"desahíj"},
	"desalentar":       {"desalient"},
	"desamoblar":       {"desamuebl"},
	"desapretar":       {"desapriet"},
	"desaprobar":       {"desaprueb"},
	"desaporcar":       {"desapuerc"},
	"desarrendar":      {"desarriend"},
	"desasentar":       {"desasient"},
	"desasosegar":      {"desasosieg"},
	"desataviar":       {"desataví"},
	"desatentar":       {"desatient"},
	"desaterrar":       {"desatierr"},
	"desaviar":         {"desaví"},
	"desainar":         {"desaín"},
	"descafeinar":      {"descafeín"},
	"descarriar":       {"descarrí"},
	"desconceptuar":    {"desconceptú"},
	"desconcertar":     {"desconciert"},
	"desconfiar":       {"desconfí"},
	"desconsolar":      {"desconsuel"},
	"descriar":         {"descrí"},
	"descolgar":        {"descuelg"},
	"descollar":        {"descuell"},
	"descontar":        {"descuent"},
	"descordar":        {"descuerd"},
	"descornar":        {"descuern"},
	"descostar":        {"descuest"},
	"desdentar":        {"desdient"},
	"desembaular":      {"desembaúl"},
	"desempedrar":      {"desempiedr"},
	"desencerrar":      {"desencierr"},
	"desengrosar":      {"desengrues"},
	"desenterrar":      {"desentierr"},
	"desgobernar":      {"desgobiern"},
	"deshabituar":      {"deshabitú"},
	"deshelar":         {"deshiel"},
	"desherbar":        {"deshierb"},
	"desherrar":        {"deshierr"},
	"desosar":          {"deshues"},
	"deslendrar":       {"desliendr"},
	"desliar":          {"deslí"},
	"desmajolar":       {"desmajuel"},
	"desmembrar":       {"desmiembr"},
	"desnevar":         {"desniev"},
	"despedrar":        {"despiedr"},
	"despernar":        {"despiern"},
	"despertar":        {"despiert"},
	"desplegar":        {"desplieg"},
	"despoblar":        {"despuebl"},
	"destentar":        {"destient"},
	"desterrar":        {"destierr"},
	"destrocar":        {"destruec", "detroc"},
	"desoldar":         {"desueld"},
	"desollar":         {"desuell"},
	"desvariar":        {"desvarí"},
	"desvergonzar":     {"desvergüenz"},
	"desvirtuar":       {"desvirtú"},
	"desviar":          {"desví"},
	"devaluar":         {"devalú"},
	"dentar":           {"dient"},
	"dezmar":           {"diezm"},
	"discontinuar":     {"discontinú"},
	"discordar":        {"discuerd"},
	"disonar":          {"disuen"},
	"dolar":            {"duel"},
	"efectuar":         {"efectú"},
	"ejecutoriar":      {"ejecutorí"},
	"emparentar":       {"emparient", "emparent"},
	"empedrar":         {"empiedr"},
	"empezar":          {"empiez"},
	"emporcar":         {"empuerc"},
	"encerrar":         {"encierr"},
	"enclocar":         {"encluec"},
	"encomendar":       {"encomiend"},
	"encontrar":        {"encuentr"},
	"encordar":         {"encuerd"},
	"encorar":          {"encuer"},
	"encovar":          {"encuev"},
	"endentar":         {"endient"},
	"enfatuar":         {"enfatú"},
	"enfriar":          {"enfrí"},
	"engrosar":         {"engrues", "engros"},
	"enhastiar":        {"enhastí"},
	"enlejiar":         {"enlejí"},
	"enmelar":          {"enmiel"},
	"enmendar":         {"enmiend"},
	"enraizar":         {"enraíz"},
	"enrocar":          {"enruec", "enroc"},
	"enrodar":          {"enrued"},
	"enriar":           {"enrí"},
	"ensangrentar":     {"ensangrient"},
	"ensarmentar":      {"ensarmient"},
	"ensoñar":          {"ensueñ"},
	"enterrar":         {"entierr"},
	"entrecerrar":      {"entrecierr"},
	"entrepernar":      {"entrepiern"},
	"entortar":         {"entuert"},
	"enviar":           {"enví"},
	"enzainar":         {"enzaín"},
	"escarmentar":      {"escarmient"},
	"esforzar":         {"esfuerz"},
	"esgrafiar":        {"esgrafí"},
	"espurriar":        {"espurrí"},
	"espiar":           {"espí"},
	"esquiar":          {"esquí"},
	"estenografiar":    {"estenografí"},
	"estregar":         {"estrieg"},
	"estriar":          {"estrí"},
	"evacuar":          {"evacú"},
	"evaluar":          {"evalú"},
	"exceptuar":        {"exceptú"},
	"expatriar":        {"expatrí"},
	"expiar":           {"expí"},
	"extasiar":         {"extasí"},
	"extenuar":         {"extenú"},
	"extraviar":        {"extraví"},
	"ferrar":           {"fierr"},
	"filiar":           {"filí"},
	"fluctuar":         {"fluctú"},
	"fotografiar":      {"fotografí"},
	"fotolitografiar":  {"fotolitografí"},
	"fregar":           {"frieg"},
	"follar":           {"fuell", "foll"},
	"forzar":           {"fuerz"},
	"fiar":             {"fí"},
	"gloriar":          {"glorí"},
	"gobernar":         {"gobiern"},
	"graduar":          {"gradú"},
	"guiar":            {"guí"},
	"habituar":         {"habitú"},
	"hacendar":         {"haciend"},
	"hastiar":          {"hastí"},
	"helar":            {"hiel"},
	"herbar":           {"hierb"},
	"herrar":           {"hierr"},
	"historiar":        {"historí"},
	"holgar":           {"huelg"},
	"hollar":           {"huell"},
	"improbar":         {"imprueb"},
	"incensar":         {"inciens"},
	"individuar":       {"individú"},
	"infatuar":         {"infatú"},
	"insinuar":         {"insinú"},
	"interactuar":      {"interactú"},
	"inventariar":      {"inventarí"},
	"judaizar":         {"judaíz"},
	"jugar":            {"jueg"},
	"licuar":           {"licú"},
	"litofotografiar":  {"litofotografí"},
	"litografiar":      {"litografí"},
	"liar":             {"lí"},
	"malcriar":         {"malcrí"},
	"mancornar":        {"mancuern"},
	"manifestar":       {"manifiest"},
	"maullar":          {"maúll"},
	"mecanografiar":    {"mecanografí"},
	"menstruar":        {"menstrú"},
	"merendar":         {"meriend"},
	"melar":            {"miel"},
	"mentar":           {"mient"},
	"moblar":           {"muebl"},
	"mostrar":          {"muestr"},
	"miar":             {"mí"},
	"negar":            {"nieg"},
	"oblicuar":         {"oblicú"},
	"ortografiar":      {"ortografí"},
	"paliar":           {"palí"},
	"parahusar":        {"parahús"},
	"patiquebrar":      {"patiquiebr"},
	"perniquebrar":     {"perniquiebr"},
	"perpetuar":        {"perpetú"},
	"pensar":           {"piens"},
	"pipiar":           {"pipí"},
	"plegar":           {"plieg"},
	"porfiar":          {"porfí"},
	"precalentar":      {"precalient"},
	"preceptuar":       {"preceptú"},
	"prohijar":         {"prohíj"},
	"promiscuar":       {"promiscú"},
	"probar":           {"prueb"},
	"poblar":           {"puebl"},
	"puntuar":          {"puntú"},
	"piar":             {"pí"},
	"quebrar":          {"quiebr"},
	"radiografiar":     {"radiografí"},
	"radiotelegrafiar": {"radiotelegrafí"},
	"reasentar":        {"reasient"},
	"recalentar":       {"recalient"},
	"recentar":         {"recient"},
	"recomendar":       {"recomiend"},
	"recomenzar":       {"recomienz"},
	"recriar":          {"recrí"},
	"recolar":          {"recuel"},
	"recontar":         {"recuent"},
	"recordar":         {"recuerd"},
	"recostar":         {"recuest"},
	"redituar":         {"reditú"},
	"reencontrar":      {"reencuentr"},
	"reenviar":         {"reenví"},
	"refregar":         {"refrieg"},
	"reforzar":         {"refuerz"},
	"regimentar":       {"regimient"},
	"regoldar":         {"regueld"},
	"rehollar":         {"rehuell"},
	"rehilar":          {"rehíl"},
	"rehusar":          {"rehús"},
	"remendar":         {"remiend"},
	"renegar":          {"renieg"},
	"renovar":          {"renuev"},
	"repatriar":        {"repatrí"},
	"repensar":         {"repiens"},
	"replegar":         {"replieg"},
	"reprobar":         {"reprueb"},
	"repoblar":         {"repuebl"},
	"requebrar":        {"requiebr"},
	"resfriar":         {"resfrí"},
	"resegar":          {"resieg"},
	"resembrar":        {"resiembr"},
	"resquebrar":       {"resquiebr"},
	"restregar":        {"restrieg"},
	"resollar":         {"resuell"},
	"resonar":          {"resuen"},
	"retemblar":        {"retiembl"},
	"retostar":         {"retuest"},
	"revaluar":         {"revalú"},
	"reventar":         {"revient"},
	"revolcar":         {"revuelc"},
	"regar":            {"rieg"},
	"rociar":           {"rocí"},
	"rodar":            {"rued"},
	"rogar":            {"rueg"},
	"salpimentar":      {"salpimient"},
	"sarmentar":        {"sarmient"},
	"sainar":           {"saín"},
	"semicerrar":       {"semicierr"},
	"segar":            {"sieg"},
	"sembrar":          {"siembr"},
	"sentar":           {"sient"},
	"serrar":           {"sierr"},
	"situar":           {"sitú"},
	"sobrecalentar":    {"sobrecalient"},
	"sobrehilar":       {"sobrehíl"},
	"sobresembrar":     {"sobresiembr"},
	"sobrevolar":       {"sobrevuel"},
	"sorregar":         {"sorrieg"},
	"sosegar":          {"sosieg"},
	"subarrendar":      {"subarriend"},
	"soldar":           {"sueld"},
	"solar":            {"suel"},
	"soltar":           {"suelt"},
	"sonar":            {"suen"},
	"soñar":            {"sueñ"},
	"superpoblar":      {"superpuebl"},
	"taquigrafiar":     {"taquigrafí"},
	"tatuar":           {"tatú"},
	"taimar":           {"taím"},
	"telegrafiar":      {"telegrafí"},
	"temblar":          {"tiembl"},
	"tentar":           {"tient"},
	"trascolar":        {"trascuel"},
	"trascordar":       {"trascuerd"},
	"trasegar":         {"trasieg"},
	"trastrocar":       {"trastruec"},
	"trasoñar":         {"trasueñ"},
	"traillar":         {"traíll"},
	"tropezar":         {"tropiez"},
	"trocar":           {"truec"},
	"tronar":           {"truen"},
	"triar":            {"trí"},
	"tonar":            {"tuen"},
	"tostar":           {"tuest"},
	"tumultuar":        {"tumultú"},
	"unisonar":         {"unisuen"},
	"usufructuar":      {"usufructú"},
	"vaciar":           {"vací"},
	"valuar":           {"valú"},
	"variar":           {"varí"},
	"volcar":           {"vuelc"},
	"volar":            {"vuel"},
	"errar":            {"yerr"},
}

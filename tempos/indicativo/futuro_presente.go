package indicativo

func FuturoPresente(verbo string) []string {
	base := verbo
	return []string{
		"eu " + base + "ei",
		"tu " + base + "ás",
		"ele " + base + "á",
		"nós " + base + "emos",
		"vós " + base + "eis",
		"eles " + base + "ão",
	}
}

package indicativo

func FuturoPreterito(verbo string) []string {
	base := verbo
	return []string{
		"eu " + base + "ia",
		"tu " + base + "ias",
		"ele " + base + "ia",
		"nós " + base + "íamos",
		"vós " + base + "íeis",
		"eles " + base + "iam",
	}
}

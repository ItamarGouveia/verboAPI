package indicativo

func Presente(verbo string) []string {
	if len(verbo) < 2 {
		return []string{"[Verbo inválido]"}
	}

	base := verbo[:len(verbo)-2]

	switch verbo[len(verbo)-2:] {
	case "ar":
		return []string{
			"eu " + base + "o",
			"tu " + base + "as",
			"ele " + base + "a",
			"nós " + base + "amos",
			"vós " + base + "ais",
			"eles " + base + "am",
		}
	case "er":
		return []string{
			"eu " + base + "o",
			"tu " + base + "es",
			"ele " + base + "e",
			"nós " + base + "emos",
			"vós " + base + "eis",
			"eles " + base + "em",
		}
	case "ir":
		return []string{
			"eu " + base + "o",
			"tu " + base + "es",
			"ele " + base + "e",
			"nós " + base + "imos",
			"vós " + base + "is",
			"eles " + base + "em",
		}
	default:
		return []string{"[Verbo irregular ou não suportado]"}
	}
}

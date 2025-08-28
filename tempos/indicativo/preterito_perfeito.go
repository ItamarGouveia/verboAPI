package indicativo

func PreteritoPerfeito(verbo string) []string {
	base := verbo[:len(verbo)-2]
	switch verbo[len(verbo)-2:] {
	case "ar":
		return []string{"eu " + base + "ei", "tu " + base + "aste", "ele " + base + "ou", "nós " + base + "amos", "vós " + base + "astes", "eles " + base + "aram"}
	case "er":
		return []string{"eu " + base + "i", "tu " + base + "este", "ele " + base + "eu", "nós " + base + "emos", "vós " + base + "estes", "eles " + base + "eram"}
	case "ir":
		return []string{"eu " + base + "i", "tu " + base + "iste", "ele " + base + "iu", "nós " + base + "imos", "vós " + base + "istes", "eles " + base + "iram"}
	default:
		return []string{"[Verbo irregular ou não suportado]"}
	}
}

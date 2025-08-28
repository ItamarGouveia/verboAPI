package indicativo

func PreteritoMaisQuePerfeito(verbo string) []string {
	base := verbo[:len(verbo)-2]
	switch verbo[len(verbo)-2:] {
	case "ar":
		return []string{"eu " + base + "ara", "tu " + base + "aras", "ele " + base + "ara", "nós " + base + "áramos", "vós " + base + "áreis", "eles " + base + "aram"}
	case "er", "ir":
		return []string{"eu " + base + "era", "tu " + base + "eras", "ele " + base + "era", "nós " + base + "êramos", "vós " + base + "êreis", "eles " + base + "eram"}
	default:
		return []string{"[Verbo irregular ou não suportado]"}
	}
}

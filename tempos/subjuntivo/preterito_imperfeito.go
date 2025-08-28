package subjuntivo

func PreteritoImperfeito(verbo string) []string {
	base := verbo[:len(verbo)-2]
	switch verbo[len(verbo)-2:] {
	case "ar":
		return []string{"se eu " + base + "asse", "se tu " + base + "asses", "se ele " + base + "asse", "se nós " + base + "ássemos", "se vós " + base + "ásseis", "se eles " + base + "assem"}
	case "er", "ir":
		return []string{"se eu " + base + "esse", "se tu " + base + "esses", "se ele " + base + "esse", "se nós " + base + "êssemos", "se vós " + base + "êsseis", "se eles " + base + "essem"}
	default:
		return []string{"[Verbo irregular ou não suportado]"}
	}
}

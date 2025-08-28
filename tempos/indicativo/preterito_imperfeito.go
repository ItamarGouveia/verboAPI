package indicativo

func PreteritoImperfeito(verbo string) []string {
	base := verbo[:len(verbo)-2]
	switch verbo[len(verbo)-2:] {
	case "ar":
		return []string{"eu " + base + "ava", "tu " + base + "avas", "ele " + base + "ava", "nós " + base + "ávamos", "vós " + base + "áveis", "eles " + base + "avam"}
	case "er", "ir":
		return []string{"eu " + base + "ia", "tu " + base + "ias", "ele " + base + "ia", "nós " + base + "íamos", "vós " + base + "íeis", "eles " + base + "iam"}
	default:
		return []string{"[Verbo irregular ou não suportado]"}
	}
}

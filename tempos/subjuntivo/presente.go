package subjuntivo

func Presente(verbo string) []string {
	base := verbo[:len(verbo)-2]
	switch verbo[len(verbo)-2:] {
	case "ar":
		return []string{"que eu " + base + "e", "que tu " + base + "es", "que ele " + base + "e", "que nós " + base + "emos", "que vós " + base + "eis", "que eles " + base + "em"}
	case "er", "ir":
		return []string{"que eu " + base + "a", "que tu " + base + "as", "que ele " + base + "a", "que nós " + base + "amos", "que vós " + base + "ais", "que eles " + base + "am"}
	default:
		return []string{"[Verbo irregular ou não suportado]"}
	}
}

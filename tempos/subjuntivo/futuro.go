package subjuntivo

func Futuro(verbo string) []string {
	base := verbo[:len(verbo)-2]
	switch verbo[len(verbo)-2:] {
	case "ar":
		return []string{"quando eu " + base + "ar", "quando tu " + base + "ares", "quando ele " + base + "ar", "quando nós " + base + "armos", "quando vós " + base + "ardes", "quando eles " + base + "arem"}
	case "er", "ir":
		return []string{"quando eu " + base + "er", "quando tu " + base + "eres", "quando ele " + base + "er", "quando nós " + base + "ermos", "quando vós " + base + "erdes", "quando eles " + base + "erem"}
	default:
		return []string{"[Verbo irregular ou não suportado]"}
	}
}

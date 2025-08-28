package imperativo

func Negativo(verbo string) []string {
	base := verbo[:len(verbo)-2]
	switch verbo[len(verbo)-2:] {
	case "ar":
		return []string{
			"não " + base + "es tu",
			"não " + base + "e você",
			"não " + base + "emos nós",
			"não " + base + "eis vós",
			"não " + base + "em vocês"}
	case "er", "ir":
		return []string{
			"não " + base + "as tu",
			"não " + base + "a você",
			"não " + base + "amos nós",
			"não " + base + "ais vós",
			"não " + base + "am vocês"}
	default:
		return []string{"[Verbo irregular ou não suportado]"}
	}
}

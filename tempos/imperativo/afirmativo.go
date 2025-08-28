package imperativo

func Afirmativo(verbo string) []string {
	base := verbo[:len(verbo)-2]
	switch verbo[len(verbo)-2:] {
	case "ar":
		return []string{
			base + "a" + " tu",
			base + "e" + " você",
			base + "emos" + " nós",
			base + "ai" + " vós ",
			base + "em" + " vocês"}
	case "er":
		return []string{
			base + "e" + " tu",
			base + "a" + " você",
			base + "amos" + " nós",
			base + "ei" + " vós ",
			base + "am" + " vocês"}
	case "ir":
		return []string{
			base + "e" + " tu",
			base + "a" + " você",
			base + "amos" + " nós",
			base + "i" + " vós ",
			base + "am" + " vocês"}

	default:
		return []string{"[Verbo irregular ou não suportado]"}
	}
}

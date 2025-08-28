package infinitivo

func Pessoal(verbo string) []string {
	base := verbo[:len(verbo)-2]
	switch verbo[len(verbo)-2:] {
	case "ar":
		return []string{
			"por " + base + "ar eu",
			"por " + base + "ares tu",
			"por " + base + "ar ele",
			"por " + base + "armos nós",
			"por " + base + "ardes vós",
			"por " + base + "arem eles"}
	case "er", "ir":
		return []string{
			"por " + base + "er eu",
			"por " + base + "eres tu",
			"por " + base + "er ele",
			"por " + base + "ermos nós",
			"por " + base + "erdes vós",
			"por " + base + "erem eles"}

	default:
		return []string{"[Verbo irregular ou não suportado]"}
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"verbo/tempos/imperativo"
	"verbo/tempos/indicativo"
	"verbo/tempos/infinitivo"
	"verbo/tempos/subjuntivo"
)

// Estrutura de resposta JSON
type Conjugacoes struct {
	Indicativo map[string][]string `json:"indicativo"`
	Subjuntivo map[string][]string `json:"subjuntivo"`
	Imperativo map[string][]string `json:"imperativo"`
	Infinitivo map[string][]string `json:"infinitivo"`
}

// Handler da API
func conjugateHandler(w http.ResponseWriter, r *http.Request) {
	verbo := r.URL.Query().Get("verbo")
	if verbo == "" {
		http.Error(w, "Parâmetro 'verbo' é obrigatório", http.StatusBadRequest)
		return
	}

	res := Conjugacoes{
		Indicativo: map[string][]string{
			"presente":             indicativo.Presente(verbo),
			"preterito_imperfeito": indicativo.PreteritoImperfeito(verbo),
			"preterito_perfeito":   indicativo.PreteritoPerfeito(verbo),
			"mais_que_perfeito":    indicativo.PreteritoMaisQuePerfeito(verbo),
			"futuro_presente":      indicativo.FuturoPresente(verbo),
			"futuro_preterito":     indicativo.FuturoPreterito(verbo),
		},
		Subjuntivo: map[string][]string{
			"presente":             subjuntivo.Presente(verbo),
			"preterito_imperfeito": subjuntivo.PreteritoImperfeito(verbo),
			"futuro":               subjuntivo.Futuro(verbo),
		},
		Imperativo: map[string][]string{
			"afirmativo": imperativo.Afirmativo(verbo),
			"negativo":   imperativo.Negativo(verbo),
		},
		Infinitivo: map[string][]string{
			"pessoal": infinitivo.Pessoal(verbo),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/conjugate", conjugateHandler)

	port := 8080
	fmt.Printf("Servidor de conjugação rodando em http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

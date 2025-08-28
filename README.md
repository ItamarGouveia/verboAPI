# verboAPI

![Go](https://img.shields.io/badge/Go-1.21-blue) ![License](https://img.shields.io/badge/License-MIT-green)

**verboAPI** é uma API escrita em Go para conjugação de verbos em português. Ela permite obter as formas conjugadas de verbos de forma rápida e prática, servindo como base para aplicações de ensino de idiomas, chatbots, ou qualquer projeto que precise trabalhar com conjugação verbal.

## Funcionalidades

- Conjugação de verbos no **presente do indicativo** (em desenvolvimento para outros tempos verbais).
- Estrutura simples e modular.
- API leve e rápida, totalmente em **Go**.

## Limitações Atuais

- Verbos não regulares e defectivos ainda não estão totalmente suportados.
- Planejamos corrigir essas questões nas próximas atualizações.

## Estrutura do Projeto

```
verboAPI/
├── main.go          # Arquivo principal do servidor
├── go.mod           # Gerenciamento de dependências
├── go.sum           # Checksums das dependências
└── tempos/          # Lógica de conjugação dos tempos verbais
```

## Como usar

1. Clonar o projeto:

```bash
git clone https://github.com/ItamarGouveia/verboAPI.git
cd verboAPI
```

2. Instalar dependências:

```bash
go mod tidy
```

3. Executar a API:

```bash
go run main.go
```

A API estará disponível em `http://localhost:8080`.

## Endpoints Disponíveis

- `/conjugate?verbo=<verbo>` → Retorna a conjugação completa do verbo.

## Exemplo de uso

```bash
curl -s "http://localhost:8080/conjugate?verbo=amar"
```

Resposta esperada:

```json
{
  "indicativo": {
    "futuro_presente": ["eu amarei","tu amarás","ele amará","nós amaremos","vós amareis","eles amarão"],
    "futuro_preterito": ["eu amaria","tu amarias","ele amaria","nós amaríamos","vós amaríeis","eles amariam"],
    "mais_que_perfeito": ["eu amara","tu amaras","ele amara","nós amáramos","vós amáreis","eles amaram"],
    "presente": ["eu amo","tu amas","ele ama","nós amamos","vós amais","eles amam"],
    "preterito_imperfeito": ["eu amava","tu amavas","ele amava","nós amávamos","vós amáveis","eles amavam"],
    "preterito_perfeito": ["eu amei","tu amaste","ele amou","nós amamos","vós amastes","eles amaram"]
  },
  "subjuntivo": {
    "futuro": ["quando eu amar","quando tu amares","quando ele amar","quando nós amarmos","quando vós amardes","quando eles amarem"],
    "presente": ["que eu ame","que tu ames","que ele ame","que nós amemos","que vós ameis","que eles amem"],
    "preterito_imperfeito": ["se eu amasse","se tu amasses","se ele amasse","se nós amássemos","se vós amásseis","se eles amassem"]
  },
  "imperativo": {
    "afirmativo": ["ama tu","ame você","amemos nós","amai vós","amem vocês"],
    "negativo": ["não ames tu","não ame você","não amemos nós","não ameis vós","não amem vocês"]
  },
  "infinitivo": {
    "pessoal": ["por amar eu","por amares tu","por amar ele","por amarmos nós","por amardes vós","por amarem eles"]
  }
}
```

## Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir **issues** ou enviar **pull requests**.

1. Fork este repositório
2. Crie uma branch para sua feature (`git checkout -b feature/nome-da-feature`)
3. Faça commit das suas alterações (`git commit -m 'Adiciona nova feature'`)
4. Faça push para a branch (`git push origin feature/nome-da-feature`)
5. Abra um Pull Request

## Licença

Este projeto está sob a licença **MIT**. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

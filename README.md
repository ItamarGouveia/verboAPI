# verboAPI

![Go](https://img.shields.io/badge/Go-1.21-blue) ![License](https://img.shields.io/badge/License-MIT-green)

**verboAPI** é uma API escrita em Go para conjugação de verbos em português. Ela permite obter as formas conjugadas de verbos de forma rápida e prática, servindo como base para aplicações de ensino de idiomas, chatbots, ou qualquer projeto que precise trabalhar com conjugação verbal.

## Funcionalidades

- Conjugação de verbos no **presente do indicativo** (em desenvolvimento para outros tempos verbais).
- Estrutura simples e modular.
- API leve e rápida, totalmente em **Go**.

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

4. Exemplo de uso (Presente do Indicativo):

```bash
curl http://localhost:8080/conjugate?verbo=falar
```

Resposta esperada:

```json
{
  "eu": "falo",
  "tu": "falas",
  "ele/ela": "fala",
  "nós": "falamos",
  "vós": "falais",
  "eles/elas": "falam"
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


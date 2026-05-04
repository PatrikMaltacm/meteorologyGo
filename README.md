# MeteorologyGo 🌦️

Uma API robusta e leve desenvolvida em **Go** para coleta e gerenciamento de dados de estações meteorológicas caseiras. Projetada para integrar dispositivos IoT como **ESP32**, **ESP8266** e **Arduino** utilizando sensores de ambiente.

---

## 🚀 Sobre o Projeto

O MeteorologyGo serve como o backend central para entusiastas de IoT. Se você possui uma estação meteorológica baseada em microcontroladores com sensores como:

- **BME280** — Pressão, Temperatura e Umidade
- **DHT11 / DHT22** — Temperatura e Umidade

Esta API é o lugar ideal para armazenar, consultar e gerenciar esses dados de forma simples e eficiente.

---

## 🛠️ Tecnologias Utilizadas

| Tecnologia | Descrição |
|---|---|
| [Go (Golang)](https://go.dev/) | Linguagem principal — alta performance e concorrência nativa |
| [Gin Gonic](https://github.com/gin-gonic/gin) | Framework web rápido e minimalista |
| [PostgreSQL](https://www.postgresql.org/) | Banco de dados relacional robusto para armazenamento dos dados |
| [go-playground/validator](https://github.com/go-playground/validator) | Validação rigorosa dos dados recebidos dos sensores |
| [godotenv](https://github.com/joho/godotenv) | Gerenciamento de variáveis de ambiente via `.env` |

---

## 📁 Estrutura do Projeto

A aplicação segue o padrão idiomático do Go ([Standard Go Project Layout](https://github.com/golang-standards/project-layout)):

```
meteorologyGo/
├── cmd/
│   └── server/
│       └── main.go          # Ponto de entrada da aplicação
├── internal/
│   ├── database/
│   │   └── database.go      # Configuração e conexão com o banco de dados (PostgreSQL)
│   ├── handler/
│   │   ├── routes.go        # Definição das rotas da API
│   │   ├── station.go       # Handlers HTTP para as estações
│   │   └── weather.go       # Handlers HTTP para os dados meteorológicos
│   └── model/
│       ├── station.go       # Modelos e estruturas de dados de estações
│       └── weather.go       # Modelos e estruturas de dados meteorológicos
├── go.mod
├── go.sum
└── README.md
```

---

## 📡 Integração e Rotas da API

### Estações

Você pode cadastrar e consultar suas estações meteorológicas.

**Criar uma estação:**
`POST /api/v1/station`
```json
{
  "lat": -23.5505,
  "long": -46.6333
}
```

**Listar todas as estações:**
`GET /api/v1/station/all`

### Dados Meteorológicos

A API recebe os dados via `POST` no endpoint `/api/v1/weather`. Exemplo de payload:

```json
{
  "station_id": "uuid-da-estacao",
  "pressure": 1013,
  "humidity": 65,
  "temp": 22,
  "lat": -23.5505,
  "long": -46.6333
}
```

**Consultar dados:**
- `GET /api/v1/weather` (retorna a última leitura registrada)
- `GET /api/v1/weather/all` (retorna todo o histórico de leituras)

---

## ⚙️ Como Começar

### Pré-requisitos

- [Go](https://go.dev/dl/) `v1.20` ou superior

### Instalação

1. **Clone o repositório:**

```bash
git clone https://github.com/PatrikMaltacm/meteorologyGo.git
cd meteorologyGo
```

2. **Instale as dependências:**

```bash
go mod tidy
```

3. **Execute a aplicação:**

```bash
go run cmd/server/main.go
```

4. **Configure suas variáveis de ambiente:**
Crie um arquivo `.env` na raiz do projeto com a URL de conexão do PostgreSQL, por exemplo:
```env
DATABASE_URL=postgres://user:password@localhost:5432/meteorology?sslmode=disable
```

---

## 🔓 Open Source & Contribuição

Este projeto é **Open Source** e está aberto para qualquer um utilizar, modificar e melhorar. Sinta-se à vontade para:

- 🐛 Abrir **Issues** para reportar bugs ou sugerir melhorias.
- 🔀 Enviar **Pull Requests** com novas funcionalidades (ex: suporte a outros bancos, painéis de visualização, etc).

---

## 📄 Licença

Distribuído sob a licença **MIT**. O uso é livre para projetos pessoais ou comerciais.

---

<p align="center">Desenvolvido por <strong>Patrik Malta</strong></p>
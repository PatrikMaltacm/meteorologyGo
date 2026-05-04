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
| [BigCache](https://github.com/allegro/bigcache) | Cache em memória ultra-rápido, reduzindo o tempo de resposta do último dado para ~4ms |
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
│   ├── cache/
│   │   └── bigcache.go      # Configuração do cache em memória (BigCache)
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

> **⚠️ Nota sobre Otimização de Memória (Decimais como Inteiros):** Para economizar memória e tráfego, as métricas dos sensores devem ser enviadas e armazenadas como números inteiros. 
> - **Umidade (`humidity`)**: Espera um `uint16`. Ex: `80.5%` deve ser enviado como `805`.
> - **Temperatura (`temp`)**: Espera um `int16`. Ex: `22.5°C` deve ser enviado como `225`.
> - **Pressão (`pressure`)**: Espera um `uint32`. Ex: `1013.25 hPa` deve ser enviado como `101325`.
> 
> **Atenção:** Latitude (`lat`) e Longitude (`long`) continuam sendo enviados como decimais padrão (`float64`). Lembre-se de reverter os valores inteiros para decimais no seu frontend (dividindo por 10 ou 100) ao exibir os dados!

```json
{
  "station_id": "uuid-da-estacao",
  "pressure": 101325,
  "humidity": 650,
  "temp": 225,
  "lat": -23.5505,
  "long": -46.6333
}
```

**Consultar dados:**
- `GET /api/v1/weather` (retorna a última leitura registrada com altíssima performance em ~4ms graças ao cache em memória)
- `GET /api/v1/weather/all` (retorna todo o histórico de leituras)

---

## 💡 Escalabilidade e Cache

A API utiliza o **BigCache** para armazenar a última leitura de dados meteorológicos em memória, garantindo tempos de resposta extremamente rápidos (em média `~4ms`). 

> **Observação:** O limite de memória do cache está configurado por padrão em `512 MB` (`HardMaxCacheSize: 512` no arquivo `internal/cache/bigcache.go`). Se a sua aplicação escalar de forma significativa (com milhares de estações registrando dados massivamente), considere aumentar esse valor de acordo com os recursos de hardware do seu servidor para evitar a expiração/remoção prematura de dados cacheados.

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
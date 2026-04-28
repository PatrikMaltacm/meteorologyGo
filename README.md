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
| [SQLite](https://www.sqlite.org/) | Banco de dados local, leve e sem configuração complexa |
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
│   │   └── database.go      # Configuração e conexão com o banco de dados (SQLite)
│   ├── handler/
│   │   ├── routes.go        # Definição das rotas da API
│   │   └── weather.go       # Handlers HTTP para os dados meteorológicos
│   └── model/
│       └── weather.go       # Modelos e estruturas de dados
├── database.db              # Arquivo do banco de dados SQLite (gerado em runtime)
├── go.mod
├── go.sum
└── README.md
```

---

## 📡 Integração com Sensores

A API recebe dados via `POST` no endpoint `/sendData`. Exemplo de payload:

```json
{
  "pressure": 1013,
  "humidity": 65,
  "temp": 22
}
```

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

4. **Inicialize o banco de dados** (somente na primeira execução), acessando via `POST`:

```
http://localhost:8080/setup
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
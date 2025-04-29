# API de Contatos em Go

Uma API RESTful para gerenciamento de contatos construída com Go, Gin e PostgreSQL.

## 📋 Índice

- [Visão Geral](#visão-geral)
- [Tecnologias](#tecnologias)
- [Recursos](#recursos)
- [Instalação](#instalação)
- [Configuração do Banco de Dados](#configuração-do-banco-de-dados)
- [Executando a Aplicação](#executando-a-aplicação)
- [Documentação da API](#documentação-da-api)
- [Endpoints](#endpoints)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Migrações](#migrações)
- [Contribuindo](#contribuindo)
- [Licença](#licença)

## 🔍 Visão Geral

Este projeto é uma API de gerenciamento de contatos que permite criar, listar, atualizar e excluir contatos. A API segue boas práticas de desenvolvimento, utiliza uma arquitetura em camadas e inclui documentação Swagger.

## 🛠️ Tecnologias

- [Go](https://golang.org/) - Linguagem de programação
- [Gin](https://github.com/gin-gonic/gin) - Framework web
- [PostgreSQL](https://www.postgresql.org/) - Banco de dados
- [Swagger](https://swagger.io/) - Documentação da API
- [Docker](https://www.docker.com/) - Containerização

## ✨ Recursos

- CRUD completo para contatos
- Documentação interativa com Swagger
- Implementação de migrações de banco de dados
- Arquitetura em camadas (Handler, Service, Repository)
- Containerização com Docker

## 🚀 Instalação

### Pré-requisitos

- Go 1.24+
- PostgreSQL
- Docker e Docker Compose (opcional)

### Clonando o repositório

```bash
git clone https://github.com/Felipe8297/go-contacts-api.git
cd go-contacts-api
```

### Instalando dependências

```bash
go mod download
```

## 🗄️ Configuração do Banco de Dados

### Usando Docker (Recomendado)

```bash
docker-compose up -d
```

### Configuração Manual

Crie um banco de dados PostgreSQL e configure as variáveis de ambiente conforme o arquivo `.env.example`.

### Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=nome_db
```

## ▶️ Executando a Aplicação

### Executando as migrações

```bash
go run cmd/migrate/main.go
```

### Iniciando o servidor

```bash
go run cmd/api/main.go
```

O servidor estará disponível em http://localhost:8080

## 📚 Documentação da API

A documentação Swagger está disponível em:

```
http://localhost:8080/swagger/index.html
```

## 🔌 Endpoints

| Método | URL | Descrição |
|--------|-----|-----------|
| GET | /contacts | Lista todos os contatos |
| GET | /contacts/:id | Obtém um contato específico |
| POST | /contacts | Cria um novo contato |
| PUT | /contacts/:id | Atualiza um contato existente |
| DELETE | /contacts/:id | Remove um contato |

## 📁 Estrutura do Projeto

```
go-contacts-api/
│
├── cmd/
│   ├── api/                # Ponto de entrada da API
│   └── migrate/            # Ferramenta de migração
│
├── docs/                   # Documentação Swagger gerada
│
├── internal/
│   ├── contacts/           # Módulo de contatos
│   │   ├── handler.go      # Manipuladores de requisições
│   │   ├── model.go        # Modelos/entidades
│   │   ├── repository.go   # Camada de acesso a dados
│   │   └── service.go      # Lógica de negócios
│   │
│   └── pkg/
│       ├── db/             # Conexão com banco de dados
│       └── migrations/     # Migrações do banco de dados
│
├── .env.example            # Exemplo de variáveis de ambiente
├── .gitignore              # Arquivos ignorados pelo Git
├── docker-compose.yaml     # Configuração Docker
├── go.mod                  # Dependências Go
└── README.md               # Documentação do projeto
```

## 🔄 Migrations

As migrations são executadas automaticamente quando a aplicação é iniciada. Os arquivos de migração estão localizados em `internal/pkg/migrations/`.

## 📄 Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

Desenvolvido por [Felipe Sousa](https://github.com/Felipe8297) 
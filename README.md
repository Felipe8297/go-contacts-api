# API de Contatos em Go

Uma API RESTful para gerenciamento de contatos construída com Go, Gin, PostgreSQL, Prometheus e Grafana.

## 📋 Índice

- [Visão Geral](#visão-geral)
- [Tecnologias](#tecnologias)
- [Recursos](#recursos)
- [Instalação](#instalação)
- [Configuração do Banco de Dados](#configuração-do-banco-de-dados)
- [Monitoramento e Observabilidade](#monitoramento-e-observabilidade)
- [Executando a Aplicação](#executando-a-aplicação)
- [Documentação da API](#documentação-da-api)
- [Endpoints](#endpoints)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Migrações](#migrações)
- [Contribuindo](#contribuindo)
- [Licença](#licença)

## 🔍 Visão Geral

Este projeto é uma API de gerenciamento de contatos que permite criar, listar, atualizar e excluir contatos. A API segue boas práticas de desenvolvimento, utiliza uma arquitetura em camadas, inclui documentação Swagger e monitoramento com Prometheus e Grafana.

## 🛠️ Tecnologias

- [Go](https://golang.org/) - Linguagem de programação
- [Gin](https://github.com/gin-gonic/gin) - Framework web
- [PostgreSQL](https://www.postgresql.org/) - Banco de dados
- [Swagger](https://swagger.io/) - Documentação da API
- [Prometheus](https://prometheus.io/) - Monitoramento de métricas
- [Grafana](https://grafana.com/) - Visualização de métricas
- [Docker](https://www.docker.com/) - Containerização

## ✨ Recursos

- CRUD completo para contatos
- Documentação interativa com Swagger
- Implementação de migrações de banco de dados
- Arquitetura em camadas (Handler, Service, Repository)
- Monitoramento de métricas com Prometheus
- Dashboards de métricas com Grafana
- Containerização com Docker

## 🚀 Instalação

### Pré-requisitos

- Docker e Docker Compose

### Clonando o repositório

```bash
git clone https://github.com/Felipe8297/go-contacts-api.git
cd go-contacts-api
```

## 🗄️ Configuração do Banco de Dados

A configuração do banco de dados é feita automaticamente pelo Docker Compose. As credenciais padrão estão definidas no arquivo [`docker-compose.yaml`](docker-compose.yaml):

- Usuário: `docker`
- Senha: `docker`
- Banco: `contactsdb`

Se necessário, ajuste as variáveis de ambiente no próprio arquivo.

## 📊 Monitoramento e Observabilidade

O projeto já está configurado para expor métricas no endpoint `/metrics`, que podem ser coletadas pelo Prometheus e visualizadas no Grafana.

- **Prometheus**: coleta métricas da API automaticamente.
- **Grafana**: permite criar dashboards para visualização das métricas.

### Acessando as ferramentas

- **Prometheus**: [http://localhost:9090](http://localhost:9090)
- **Grafana**: [http://localhost:3000](http://localhost:3000)  
  Usuário padrão: `admin`  
  Senha padrão: `admin`

### Exemplo de métricas expostas

- `http_requests_total`
- `http_request_duration_seconds`
- `database_operations_total`
- `database_operation_duration_seconds`

Você pode criar dashboards no Grafana utilizando o Prometheus como fonte de dados.

## ▶️ Executando a Aplicação

### Usando Docker Compose

Com todos os serviços configurados, basta executar:

```bash
docker-compose up -d
```

Isso irá subir os containers da API, banco de dados, Prometheus e Grafana.  
A API estará disponível em [http://localhost:8080](http://localhost:8080).

### Parando os serviços

```bash
docker-compose down
```

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
| GET | /metrics | Métricas Prometheus |

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
├── prometheus/             # Configuração do Prometheus
├── .env.example            # Exemplo de variáveis de ambiente
├── .gitignore              # Arquivos ignorados pelo Git
├── docker-compose.yaml     # Configuração Docker
├── Dockerfile              # Build da aplicação
├── go.mod                  # Dependências Go
└── README.md               # Documentação do projeto
```

## 🔄 Migrations

As migrations são executadas automaticamente quando a aplicação é iniciada dentro do container. Os arquivos de migração estão localizados em `internal/pkg/migrations/`.

## 📄 Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

Desenvolvido por [Felipe Silva](https://github.com/Felipe8297)
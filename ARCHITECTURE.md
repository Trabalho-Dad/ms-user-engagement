# 🎯 Screaming Architecture - ms-feedbacks

## 📢 O que é "Screaming Architecture"?

**Screaming Architecture** é um estilo onde a estrutura grita **o propósito do sistema**, não a tecnologia.

### Antes (Estrutura "silenciosa"):
```
internal/
├── domain/
├── usecase/
├── http/
├── infrastructure/
```
❌ Olhando para isso, não fica claro: "Qual é o propósito dessa aplicação?"

### Depois (Screaming Architecture):
```
internal/
├── feedbacks/               ✨ GRITA: "Isso é sobre Feedbacks!"
│   ├── domain/              ✨ GRITA: "Regras de feedback"
│   ├── application/         ✨ GRITA: "Casos de uso de feedback"
│   └── adapters/            ✨ GRITA: "Como feedback é recebido/persistido"
```
✅ Agora fica claro: "Essa é uma aplicação de feedback!"

### Nomes Explícitos Gritam Ainda Mais:
```
domain/
├── feedback.go              ✨ GRITA: "Entidade Feedback"
├── feedbackRepository.go    ✨ GRITA: "Repositório de Feedback"

application/
├── feedbackService.go       ✨ GRITA: "Serviço de Feedback"

adapters/http/
├── feedbackHttpHandler.go   ✨ GRITA: "Handler HTTP de Feedback"

adapters/repository/
├── feedbackMemoryRepository.go  ✨ GRITA: "Feedback em Memória"
```

---

## 🎯 O que é Screaming Architecture?

Screaming Architecture é um estilo de organização de código onde a **estrutura de pastas "grita" sobre o propósito da aplicação** - não sobre tecnologias. Olhando para os diretórios, fica imediatamente claro que é um serviço de **feedback**.

Baseado nos princípios de Clean Architecture (Uncle Bob), mas com foco em deixar o negócio evidente.

---

## Estrutura de Diretórios

```
ms-feedbacks/
├── cmd/
│   └── main.go                  # Entry point simples
│
├── internal/
│   ├── feedbacks/               # ⭐ O DOMÍNIO PRINCIPAL GRITA
│   │   ├── domain/              # 📋 Regras de Negócio Puras
│   │   │   ├── feedback.go                     # ✨ Entidade Feedback
│   │   │   └── feedbackRepository.go          # ✨ Contrato do Repositório
│   │   │
│   │   ├── application/         # 🔄 Casos de Uso (Orquestração)
│   │   │   └── feedbackService.go             # ✨ Service de Feedback
│   │   │
│   │   └── adapters/            # 🔌 Implementações Técnicas
│   │       ├── http/            # HTTP adapters (entrada)
│   │       │   └── feedbackHttpHandler.go     # ✨ Handler HTTP
│   │       │
│   │       └── repository/      # Repository adapters (saída)
│   │           ├── feedbackMemoryRepository.go   # ✨ Em memória
│   │           └── postgres/                     # (futuro) PostgreSQL
│   │               ├── connection.go
│   │               └── feedbackPostgresRepository.go
│   │
│   ├── shared/                  # 🛠️ Utilitários Compartilhados
│   │   ├── http/
│   │   │   ├── router.go        # Configuração de rotas Gin
│   │   │   └── health_handler.go
│   │   ├── middleware/
│   │   ├── auth/
│   │   ├── context/
│   │   ├── env_loader.go        # Carregamento de variáveis
│   │   └── logger.go            # Logging
│   │
│   └── app/
│       └── bootstrap.go         # Composition Root (DI - Injeção de Dependências)
│
├── shared/                      # 📦 Utilitários Globais
│   ├── env_loader.go
│   └── logger.go
│
├── go.mod
├── go.sum
├── LICENSE
└── ARCHITECTURE.md
```

---

## Camadas Explicadas

### 1. **Domain Layer** (`feedbacks/domain/`)
✅ **Pura, isenta de dependências externas**

- Contém entidades (`feedback.go`: Feedback) e interfaces (`feedbackRepository.go`: FeedbackRepository)
- Todas as regras de validação de negócio ficam aqui
- Não conhece sobre HTTP, banco de dados, frameworks
- Pode ser reutilizada em qualquer contexto (CLI, gRPC, etc)

**Arquivos:**
- `feedback.go`: Entidade Feedback com construtor `New()` e validações
- `feedbackRepository.go`: Interface que define o contrato de persistência

### 2. **Application Layer** (`feedbacks/application/`)
🔄 **Orquestração da lógica de negócio**

- Implementa casos de uso concretos
- Coordena a interação entre domain + infrastructure
- Contém `FeedbackService` que orquestra criação/listagem de feedbacks
- Ainda agnóstica a tecnologia (HTTP/gRPC/CLI)

**Arquivos:**
- `feedbackService.go`: FeedbackService com métodos `Create()` e `List()`

### 3. **Adapters Layer** (`feedbacks/adapters/`)
🔌 **Implementações técnicas - Isoladas e Substituíveis**

#### **HTTP Adapter** (`adapters/http/`)
- `feedbackHttpHandler.go`: FeedbackHttpHandler que transforma requisições HTTP em DTOs
- Valida entrada, chama service, retorna resposta

#### **Repository Adapter** (`adapters/repository/`)
- `feedbackMemoryRepository.go`: Em memória (para dev/testes)
- `postgres/`: (futuro) Implementação com PostgreSQL

**Vantagem**: Trocar de banco de dados é apenas trocar o adapter!

### 4. **Shared Layer** (`internal/shared/`)
🛠️ **Utilitários reutilizáveis**

- **HTTP**: Router, handlers compartilhados, middlewares
- **Middleware**: Authorization, logging, etc
- **Auth**: Lógica de autenticação compartilhada
- **Context**: Helpers para contexto HTTP
- **Config**: Env loader, logger

### 5. **Bootstrap** (`app/bootstrap.go`)
🚀 **Composition Root - Injeção de Dependências**

Responsável por:
- Instanciar repositório (FeedbackMemoryRepository ou futuro FeedbackPostgresRepository)
- Instanciar service (FeedbackService)
- Montar router
- Retornar app pronto para rodar

```go
func New() *App {
    repo := repository.NewFeedbackMemoryRepository()
    service := application.NewFeedbackService(repo)
    router := sharedhttp.NewRouter(service)
    return &App{router, logger}
}
```

---

## Fluxo de uma Requisição

```
HTTP GET /v1/feedbacks
    ↓
[router.go] → NewRouter() configura rota
    ↓
[feedbackHttpHandler.go] → FeedbackHttpHandler.List() recebe requisição
    ↓
[feedbackService.go] → FeedbackService.List() executa lógica
    ↓
[feedback.go] → Validações de negócio
    ↓
[feedbackMemoryRepository.go] → Adapter retorna dados
    ↓
Handler serializa → JSON response
```

---

## Vantagens da Screaming Architecture

✅ **Independência de Frameworks**
- Trocar de Gin para Echo? Apenas adapte `http/`
- Trocar de memória para PostgreSQL? Apenas adapte `repository/`

✅ **Testabilidade**
- Domain é 100% testável (sem dependências)
- Service é testável com mock de repository
- Handlers testáveis com mock de service

✅ **Clareza de Intenção**
- Olhar para diretórios grita: "Isso é um serviço de feedback!"
- Novos desenvolvedores entendem rápido a estrutura

✅ **Escalabilidade**
- Fácil adicionar novos casos de uso em `application/`
- Fácil adicionar novos adapters (gRPC, AMQP, etc)

✅ **Manutenibilidade**
- Mudanças em uma camada não afetam as outras
- Responsabilidades bem definidas

---

## Como Adicionar Novos Recursos

### Exemplo: Adicionar autenticação de feedback

1. **Domain** (`domain/`)
   - Atualizar `feedback.go`: adicionar campo `AuthorID`
   - Validar AuthorID em `Validate()`

2. **Application** (`application/`)
   - Atualizar `feedbackService.go`
   - Adicionar campo em `CreateFeedbackInput`
   - Validar permissões em `FeedbackService.Create()`

3. **Adapters** (`adapters/http/`)
   - Atualizar `feedbackHttpHandler.go`
   - Extrair autor do JWT no handler
   - Passar para service

4. **Infrastructure** (`adapters/repository/`)
   - Atualizar `feedbackMemoryRepository.go` (para testes)
   - Adicionar coluna em schema do Postgres
   - Atualizar `feedbackPostgresRepository.go` (futuro)

---

## Comandos Úteis

```bash
# Compilar
go build -o ms-feedbacks ./cmd

# Rodar
./ms-feedbacks

# Testar domain
go test ./internal/feedbacks/domain

# Testar application
go test ./internal/feedbacks/application

# Testar tudo
go test ./...
```

---

## Próximos Passos

- [ ] Implementar `feedbackPostgresRepository.go` com conexão PostgreSQL
- [ ] Adicionar autenticação/autorização em `shared/auth/`
- [ ] Adicionar testes unitários para domain, application, adapters
- [ ] Configurar logging estruturado em `shared/logger.go`
- [ ] Adicionar métricas/observabilidade
- [ ] Criar novo domínio (ex: `orders/`, `customers/`) seguindo mesma estrutura

---

**Baseado em:**
- Clean Architecture - Uncle Bob
- Screaming Architecture - Uncle Bob
- Domain-Driven Design (DDD)

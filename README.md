# go_mod_api — API REST em Go (Produtos)

Projeto de estudo com foco em **começar a trabalhar com REST API em Golang**, usando **Docker** para subir os containers da aplicação e do banco, e o **DBeaver** para visualizar/administrar o PostgreSQL.

## Objetivo

Praticar a construção de uma API REST em Go do zero, aplicando uma separação simples de camadas (repository, usecase, controller) e persistindo dados em um PostgreSQL rodando em container.

## Stack utilizada

- **Go** 1.26.8
- **Gin** (`github.com/gin-gonic/gin`) — framework web/router HTTP
- **lib/pq** (`github.com/lib/pq`) — driver PostgreSQL para `database/sql`
- **PostgreSQL 12** — banco de dados (via imagem oficial `postgres:12`)
- **Docker / Docker Compose** — orquestração dos containers
- **DBeaver** — cliente gráfico usado para inspecionar o banco durante o desenvolvimento

## Arquitetura do projeto

O projeto segue uma separação em camadas inspirada em Clean Architecture, simplificada para fins didáticos:

```
cmd/            → ponto de entrada da aplicação (main.go)
controller/     → recebe a requisição HTTP, valida entrada e devolve resposta
usecase/        → regra de negócio (orquestra o repository)
repository/     → acesso a dados (queries SQL)
model/          → structs de domínio (Product, Response)
db/             → conexão com o banco de dados
```

Fluxo de uma requisição:

```
Rota (main.go) → Controller → UseCase → Repository → Banco de Dados
```

## Estrutura de arquivos

```
.
├── cmd/
│   └── main.go
├── controller/
│   └── product_name.go
├── usecase/
│   └── product_usercase.go
├── repository/
│   └── product_repository.go
├── model/
│   ├── product.go
│   └── response.go
├── db/
│   └── conn.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum
```

## Endpoints disponíveis

| Método | Rota             | Descrição                              |
|--------|-------------------|-----------------------------------------|
| GET    | `/ping`           | Testa se o servidor está no ar          |
| GET    | `/products`       | Lista todos os produtos                 |
| GET    | `/product/:prudctId` | Busca um produto pelo ID              |
| POST   | `/product`        | Cria um novo produto                    |

### Exemplo de payload — `POST /product`

```json
{
  "name": "Batata frita",
  "price": 20
}
```

### Exemplo de resposta — `GET /product/:prudctId`

```json
{
  "id_product": 1,
  "name": "Batata frita",
  "price": 20
}
```

## Banco de dados

O container `go_db` sobe uma imagem `postgres:12` com as seguintes credenciais (definidas no `docker-compose.yml` e replicadas em `db/conn.go`):

- **Host:** `go_db` (nome do serviço no Docker Compose, resolvido via rede interna do Docker)
- **Porta:** `5432`
- **Usuário:** `postgres`
- **Senha:** `1234`
- **Database:** `postgres`

Os dados são persistidos em um volume Docker (`pgdata`), então o conteúdo do banco sobrevive mesmo que o container seja recriado.

### Tabela `product`

A aplicação espera uma tabela `product` com pelo menos as colunas:

```sql
CREATE TABLE product (
    id           SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    price        NUMERIC NOT NULL
);
```

> Essa tabela foi criada manualmente durante os testes, conectando no banco pelo **DBeaver** (host `localhost`, porta `5432`, usuário/senha conforme acima, já que a porta do Postgres está exposta no host via `docker-compose.yml`).

## Como rodar o projeto

1. Ter o Docker e o Docker Compose instalados.
2. Na raiz do projeto, subir os containers:

   ```bash
   docker-compose up --build
   ```

3. A API sobe em `http://localhost:8000` e o Postgres fica exposto em `localhost:5432` (podendo ser acessado pelo DBeaver ou outro cliente).
4. Testar se subiu corretamente:

   ```bash
   curl http://localhost:8000/ping
   ```

## Pontos de aprendizado deste projeto

- Estruturação básica de uma API REST em Go com Gin.
- Separação de responsabilidades em camadas (repository / usecase / controller).
- Uso do pacote `database/sql` + driver `lib/pq` para falar com PostgreSQL.
- Uso do Docker Compose para orquestrar aplicação + banco de dados juntos.
- Uso do DBeaver como ferramenta de inspeção/administração do banco durante o desenvolvimento.

## Possíveis melhorias futuras

- Corrigir o nome do parâmetro de rota `:prudctId` → `:productId`.
- Implementar os métodos de **UPDATE** e **DELETE** (hoje só existem GET e POST).
- Mover credenciais do banco para variáveis de ambiente em vez de deixá-las hardcoded em `db/conn.go`.
- Tratar melhor os erros nos controllers (hoje alguns `ctx.JSON` de erro não têm `return` logo em seguida, ex.: `GetProducts`).
- Adicionar validações no `model.Product` (ex.: `binding:"required"` nas tags do Gin).

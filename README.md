# drink-counter-api 🍺

## Objetivo

Basicamente, esse projeto tem como objetivo sacanear Miguel, um amigo meu que bebe muita bebida alcóolica. Aproveitei essa oportunidade para aprender a desenvolver utilizando GoLang, aprendendo a mexer nessa linguagem extremamente poderosa e em seus frameworks/libs.

Aproveitando a oportunidade, também decidi que faria esse projeto da maneira mais documentada e certinha que eu conseguisse _(mesmo sabendo que ainda teriam coisas que provavelmente eu deixaria passar)_. Por isso, comecei a fazer a documentação e a escrever os PRs utilizando Markdown e colocarei aqui tudo relacionado a arquitetura, as tecnologias e as bibliotecas utilizadas na construção da **`drink-counter-api`**.

## 1. Tecnologias Utilizadas

A API em si é desenvolvida em **Go** e utiliza um conjunto de bibliotecas e ferramentas para gerenciamento de rotas, interação com o banco de dados (PostgreSQL), autenticação e validação de dados. A seguir, estão as principais tecnologias empregadas:

| Categoria                 | Tecnologia                                                            | Descrição                                                                    |
| :------------------------ | :-------------------------------------------------------------------- | :--------------------------------------------------------------------------- |
| **Linguagem**             | Go (Golang)                                                           | Linguagem de programação para o desenvolvimento da API.                      |
| **Framework Web**         | [Gorilla Mux](https://github.com/gorilla/mux)                         | Roteador HTTP robusto para mapeamento de URLs e manipulação de requisições.  |
| **ORM**                   | [GORM](https://gorm.io/)                                              | ORM (Object-Relational Mapping) para interação com o banco de dados.         |
| **Banco de Dados**        | [PostgreSQL](https://www.postgresql.org/)                             | Banco de Dados relacional.                                                   |
| **Migrações**             | [go-gormigrate](https://github.com/go-gormigrate/gormigrate)          | Biblioteca para gerenciar migrações de esquema do banco de dados com GORM.   |
| **Validação**             | [go-playground/validator](https://github.com/go-playground/validator) | Biblioteca para validação de estruturas de dados (structs) com base em tags. |
| **Autenticação**          | [golang-jwt/jwt](https://github.com/golang-jwt/jwt)                   | Implementação de JSON Web Tokens (JWT) para autenticação e autorização.      |
| **CORS**                  | [gorilla/handlers](https://github.com/gorilla/handlers)               | Middleware para lidar com Cross-Origin Resource Sharing (CORS).              |
| **Variáveis de Ambiente** | [joho/godotenv](https://github.com/joho/godotenv)                     | Carrega variáveis de ambiente de um arquivo `.env` para a aplicação.         |

## 2. Arquitetura

A arquitetura da API segue um padrão modular em **_features_**, com a lógica de negócio organizada em pacotes distintos para facilitar a manutenção e escalabilidade. A estrutura de diretórios reflete essa modularidade:

```
drink-counter-api/
├── driver/             # Conexão com DB e migrações
├── posts/              # Lógica de negócio para posts
├── users/              # Lógica de negócio para usuários (rotas, modelos, schemas, serviços, utils)
├── .../                # Outras features seguindo o mesmo padrão de users e posts
├── utils/              # Utilitários globais (erros, variáveis de ambiente, paginação)
├── server.go           # Ponto de entrada da aplicação
├── go.mod              # Módulos e dependências do Go
├── go.sum              # Checksums das dependências
├── .env.example        # Exemplo de arquivo de variáveis de ambiente
├── docker-compose.yaml # Configuração do Docker Compose para o banco de dados
└── README.md           # Este arquivo de documentação
```

Com isso apresentado, com excessão do módulos das features, vou descrever alguns módulos que valem a pena ter uma descrição mais detalhada de seu funcionamento.

### 2.1. Ponto de Entrada (`server.go`)

O arquivo `server.go`, é o arquivo principal da aplicação, sendo responsável por:

- Carregar variáveis de ambiente (`utils.LoadEnv()`).
- Estabelecer a conexão com o banco de dados PostgreSQL (`driver.Connect()`).
- Executar as migrações do banco de dados (`driver.RunMigrations()`).
- Configurar o roteador principal utilizando `Gorilla Mux`.
- Inicializar os módulos da API através de suas funções `Init`(Ex: `users.Init(main_router, db)`)
- Configurar o middleware CORS (`github.com/gorilla/handlers`) para permitir requisições de diferentes origens.
- Iniciar o servidor HTTP na porta configurada.

### 2.2. Módulo de Driver (`driver/`)

O diretório `driver/` é responsável pela interação com o banco de dados:

- **`driver/driver.go`**: Contém a função `Connect()` para estabelecer a conexão com o banco de dados PostgreSQL usando GORM e a função `Close()` para fechar a conexão.

- **`driver/migrate.go`**: Gerencia as migrações do banco de dados utilizando `go-gormigrate`. Define as migrações para criar as tabelas `User`, `Post` e `Comment` no banco de dados.

### 2.3. Módulo de Utilitários (`utils/`)

O diretório `utils/` agrupa funções e constantes de uso geral em toda a aplicação:

- **`utils/utils.go`**: Inclui funções para carregar variáveis de ambiente (`LoadEnv()`), calcular offsets para paginação (`CalculateOffset()`), e constantes como `DATEFORMAT` e `PAGESIZE`.

- **`utils/db_errors/db_errors.go`**: Centraliza o tratamento de erros de banco de dados, mapeando erros comuns do GORM para respostas HTTP apropriadas.

- **`utils/schema_errors/schema_errors.go`**: Lida com erros de validação de esquemas e erros de parsing JSON, fornecendo mensagens de erro padronizadas para o cliente.

### 2.4. Features do Projeto (ex: `/users`, `/posts`, etc)

Para entender melhor como funciona cada feature, visite a [página de documentação](https://even-season-1cf.notion.site/drink-counter-api-3349546016ee804badf6c1492af6c7cf) do projeto que foi feita utilizando a ferramenta Notion.

## 3. Configuração e Execução

### 3.1. Variáveis de Ambiente

O projeto utiliza um arquivo `.env` para gerenciar variáveis de ambiente. Um exemplo (`.env.example`) é fornecido com as seguintes variáveis:

- `DATABASE_URL`: String de conexão com o banco de dados PostgreSQL.
- `PORT`: Porta em que a API será executada (ex: `8080`).
- `HOST`: Host da aplicação (ex: `http://localhost`).

### 3.2. Docker Compose

O arquivo `docker-compose.yaml` facilita a configuração de um ambiente de desenvolvimento com um banco de dados PostgreSQL:

```yaml
version: "3.8"
services:
  db:
    image: postgres
    restart: always
    environment:
      POSTGRES_DB: drink_counter
      POSTGRES_USER: miguel_feira
      POSTGRES_PASSWORD: f4lion
    ports:
      - "5432:5432"
    volumes:
      - drink_counter_postgres_data:/var/lib/postgresql

volumes:
  drink_counter_postgres_data:
```

Para iniciar o banco de dados PostgreSQL usando Docker Compose, execute:

```bash
docker-compose up -d
```

### 3.3. Execução da Aplicação

Após configurar as variáveis de ambiente e iniciar o banco de dados, a aplicação pode ser executada com:

```bash
go run server.go
```

## Referências (Documentações)

[1] Gorilla Mux. Disponível em: [https://github.com/gorilla/mux](https://github.com/gorilla/mux)
[2] GORM. Disponível em: [https://gorm.io/](https://gorm.io/)
[3] go-gormigrate. Disponível em: [https://github.com/go-gormigrate/gormigrate](https://github.com/go-gormigrate/gormigrate)
[4] go-playground/validator. Disponível em: [https://github.com/go-playground/validator](https://github.com/go-playground/validator)
[5] golang-jwt/jwt. Disponível em: [https://github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt)
[6] gorilla/handlers. Disponível em: [https://github.com/gorilla/handlers](https://github.com/gorilla/handlers)
[7] joho/godotenv. Disponível em: [https://github.com/joho/godotenv](https://github.com/joho/godotenv)

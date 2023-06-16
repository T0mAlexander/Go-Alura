# Projeto Go-Store

## Descrição

Este projeto é um exemplo de loja virtual com CRUD embarcado que performa as operações de persistências de dados mais básicas que são criação, edição, atualização e exclusão de produtos fictícios

## Bibliotecas e dependências

- **lib/pq:** driver de banco de dados para Golang bastante utilizado em desenvolvimento de aplicações que demandam acesso e interações com o PostgreSQL

## Pré-requisitos

- **[Docker](https://docs.docker.com/get-docker/)** instalado na máquina com versão mínima em 24.0.0

- **[Golang](https://go.dev/dl/)** instalado na máquina com versão mínima em 1.20

## Instruções de uso

1. Clone o repositório

```bash
git clone -b go-store https://github.com/T0mAlexander/Go-Alura
```

2. Navegue até o diretório

```bash
cd Go-Alura
```

3. Construa e inicialize o container do Docker em sua máquina a partir do arquivo `docker-compose.yml`

```bash
docker-compose up -d
```

4. Compile e instancie o servidor

```bash
go run main.go
```

> Nota: o servidor irá ser executado na **porta 3333**. Para acessar a aplicação web, digite na sua barra de endereço [localhost:3333](http://localhost:3333)

## Modelagem do banco de dados

| Coluna        | Tipo    | Descrição             |
|---------------|---------|-----------------------|
| Id            | Int     | Identificador único   |
| Amount        | Int     | Quantidade            |
| Name          | String  | Nome do produto       |
| Description   | String  | Descrição do produto  |
| Price         | Float64 | Preço do produto      |

## Termos de uso

Este projeto é de livre uso para outros sem nenhuma restrição para cópias ou forks 👍🏻

### Autor: Tom Alexander

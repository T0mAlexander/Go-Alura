# Go API Rest

## Descrição

Projeto exemplar com backend e API REST construída com Golang e React

## Biblioteca e dependências

**GORM:** biblioteca de mapeamento objeto-relacional (ORM) para Go que facilita a interação com bancos de dados relacionais, fornecendo uma camada de abstração para criar, recuperar, atualizar e excluir registros de forma simplificada e eficiente.

**GORM/Postgres:** driver que permite a integração entre o GORM e o banco de dados PostgreSQL, fornecendo funcionalidades específicas para conectar, executar consultas e manipular dados.

**React**: biblioteca JavaScript de código aberto para construção de interfaces de usuário interativas e reativas.

**Node.js:** ambiente de execução JavaScript assíncrono e orientado a eventos, projetado para criar aplicativos de rede escaláveis.

## Pré-requisitos

* [Docker](https://docs.docker.com/get-docker/) instalado na máquina com versão mínima em 24.0.0

* [Golang](about:blank) instalado na máquina com versão mínima em 1.20

* [Node.js](https://nodejs.org/en/download) instalado na máquina com versão mínima em 17.0.0

## Instruções de uso

1. Clone o repositório

```bash
git clone -b go-api-rest https://github.com/T0mAlexander/Go-Alura
```

2. Navegue até o diretório

```bash
cd Go-Alura
```

3. Construa e inicialize o container do Docker a partir do arquivo `docker-compose.yml`

```bash
docker-compose up -d
```

4. Instale as dependências do front-end

```bash
npm install
```

5. Inicialize o back-end e front-end

```bash
go run main.go && npm run start
```

> Nota: o back-end será inicializado na **porta 3333** enquanto que o front-end será renderizado automáticamente numa nova aba na **porta 3000**

## Modelagem do banco de dados

| Campo    | Tipo   | Tag JSON     |
|----------|--------|--------------|
| Id       | int    | json:"id"    |
| Name     | string | json:"name"  |
| History  | string | json:"history"|

## Termos de uso

Este projeto é de livre uso para outros sem nenhuma restrição para cópias ou forks 👍🏻

### Autor: Tom Alexander

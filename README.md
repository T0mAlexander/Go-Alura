# Gin API REST

## Descrição

Projeto exemplar com backend e API REST construída com Golang, Gin, GORM, Docker e PostgreSQL

## Biblioteca e dependências

**GORM:** biblioteca de mapeamento objeto-relacional (ORM) para Go que facilita a interação com bancos de dados relacionais, fornecendo uma camada de abstração para criar, recuperar, atualizar e excluir registros de forma simplificada e eficiente.

**GORM/Postgres:** driver que permite a integração entre o GORM e o banco de dados PostgreSQL, fornecendo funcionalidades específicas para conectar, executar consultas e manipular dados.

**Gin:** framework web leve e rápido para a linguagem Go. Com uma sintaxe concisa, ele simplifica o desenvolvimento de APIs REST e aplicações web.

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

4. Inicialize o back-end

```bash
go run main.go
```

> Nota: o back-end será inicializado na **porta 3333**

## Modelagem do Banco de Dados

| Campo | Tipo  | Descrição     |
|-------|-------|---------------|
| ID    | uint  | Identificador |
| Name  | string| Nome do estudante |
| CPF   | string| Número de CPF |
| RG    | string| Número de RG |

## Rotas da aplicação

- **GET** "/students": Retorna todos os estudantes cadastrados. Função: `controllers.ShowAllStudents`
- **POST** "/students": Cria um novo estudante. Função: `controllers.NewStudent`
- **PATCH** "/students/:id": Atualiza informações de um estudante específico. Função: `controllers.EditStudent`
- **GET** "/students/:id": Retorna as informações de um estudante específico. Função: `controllers.FindById`
- **DELETE** "/students/:id": Exclui um estudante específico. Função: `controllers.DeleteStudent`

## Termos de uso

Este projeto é de livre uso para outros sem nenhuma restrição para cópias ou forks 👍🏻

### Autor: Tom Alexander

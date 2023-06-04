# Projeto Go-Bank

> ### Resumo
>
> Este é um projeto de exemplar simulando sistema bancário simples em Go

## Estrutura do Projeto

- **main.go**: Arquivo principal que contém a função `main` e demonstra o uso das contas bancárias.
- **go.work**: Arquivo de configuração do Go para a versão 1.20 e configuração do projeto principal.
- **go.mod**: Arquivo do módulo Go para o projeto go-bank.
- **accounts/current.go**: Arquivo contendo a definição e métodos da conta corrente (`CurrentAccount`).
- **accounts/savings.go**: Arquivo contendo a definição e métodos da conta poupança (`SavingsAccount`).
- **customers/customers.go**: Arquivo contendo a definição da estrutura de dados para os titulares de conta (`AccHolder`).

## Descrição

O projeto `go-bank` é um exemplo de sistema bancário simples escrito em Go. Ele implementa duas contas bancárias: conta corrente (`CurrentAccount`) e conta poupança (`SavingsAccount`), e demonstra algumas operações bancárias básicas, como depósito, saque e transferência.

## Dependências

Este projeto não possui dependências externas.

## Uso

Para executar o projeto, execute o seguinte comando:

```shell
go run main.go
```

Isso irá rodar o arquivo `main.go`, que cria instâncias de `SavingsAccount` e `CurrentAccount`, realiza transações de pagamento de contas e exibe os saldos das contas.

Certamente! Aqui está uma versão modificada da seção "Arquivos Principais" da documentação do projeto Go-Bank, substituindo o código pelos conceitos e funcionalidades do projeto:

## Funcionalidades Principais

O projeto Go-Bank é um sistema bancário básico desenvolvido em Go que permite a criação e gerenciamento de contas bancárias. Ele oferece as seguintes funcionalidades principais:

1. **Conta Corrente**: O projeto suporta a criação e manipulação de contas correntes. Cada conta corrente possui informações como titular da conta, número da agência e número da conta.

2. **Conta Poupança**: O sistema também suporta a criação e operações em contas poupanças. As contas poupanças possuem características semelhantes às contas correntes, mas com recursos específicos para poupança.

3. **Depósito**: O sistema permite que os clientes façam depósitos em suas contas correntes ou poupanças. Os valores depositados são adicionados ao saldo da conta correspondente.

4. **Saque**: Os titulares de conta podem efetuar saques de suas contas correntes ou poupanças, desde que haja fundos suficientes. Os saldos são atualizados de acordo após a transação.

5. **Transferência**: Os clientes têm a opção de transferir fundos de uma conta para outra. Tanto contas correntes quanto contas poupanças são compatíveis com esse recurso.

## Fluxo do Projeto

1. O projeto começa criando uma instância de `SavingsAccount` e `CurrentAccount` para representar as contas bancárias.

2. Em seguida, são realizadas operações de depósito nas contas bancárias para adicionar fundos.

3. A função `BillPayment` é chamada para efetuar pagamentos de contas, que é uma operação de saque. O valor da conta é deduzido do saldo correspondente.

4. O saldo atualizado das contas bancárias é exibido no final do programa.

Com essa abordagem, você pode fornecer uma visão geral das funcionalidades e do fluxo do projeto sem entrar em detalhes específicos de implementação.

### Autor: Tom Alexander

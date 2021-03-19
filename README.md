# Backend Assessment

_Autor: Otávio Celani_

## Instruções de uso

### Execução

1. Certifique-se de ter o Docker instalado na máquina

```bash
docker -v
```

2. O comando make levanta o cluster de containers em conjunto à API

```bash
make
```

3. Após o passo anterior, é possível testar a API com o comando

```bash
make test
```

### Arquitetura

#### Clean Architecture

Esse projeto aplica o conceito arquitetural chamado [Clean Architecture](https://medium.com/luizalabs/descomplicando-a-clean-architecture-cf4dfc4a1ac6). Esse conceito é um modelo de [Arquitetura Exagonal](https://medium.com/dev-cave/arquitetura-hexagonal-4668a8ffac57), da qual abstrai separadamente a lógica de negócio da aplicação com o objetivo de tornar os componentes reutilizáveis e extensíveis. Dessa forma, esse projeto possui a capacidade de se integrar facilmente a qualquer outro banco de dados, através da confecção simples de um adapter. Nesse projeto, foi utilizado o banco de dados MongoDB, mas também poderia utilizar qualquer outro banco de dados SQL ou NoSQL.

#### Microsserviços e Proxy Reverso

Foi utilizado o modelo de microsserviços, do qual os componentes da aplicação foram separados em pequenas APIs. Isso permite que o projeto realize incrementos contínuos em suas funcionalidades com o acréscimo de novos serviços, o que o torna extensível. Para tanto, foi utilizado a ferramenta [Traefik](https://traefik.io/) na função de [Proxy Reverso](https://medium.com/devane-io/o-que-%C3%A9-proxy-reverso-e-por-que-voc%C3%AA-usa-um-todos-os-dias-4df32a285e0), do qual realiza a conexão entre as aplicações distribuídas da plataforma. A decisão pelo recurso é uma forma de assegurar a proteção do sistema e também para realizar o balanceamento de cargas.

### Modelo de dados

    DÉBITO AUTOMÁTICO
      ID        primitive.ObjectID
      Name      string
      Amount    string
      Frequency string
      Status    string

### Rotas

#### Autenticação

- Cria uma nova senha
  POST `localhost:80/auth/add/:pass`

- Realiza a autenticação
  GET `localhost:80/auth`

**Obs:** o header se chama `X-Auth` e já consta com o token default `jwt-secret` pré cadastrado.

#### Débito automático

- Registra um novo débito automático
  POST `localhost:80/auto-debit/add`
- Visualiza todos os débitos cadastrados
  GET `localhost:80/auto-debit/all`
- Busca por um débito automático com o ID ou Nome
  GET `localhost:80/auto-debit/:find`
- Realiza uma query para visualizar todos os débitos com um determinado status
  GET `localhost:80/auto-debit?status=:status`
- Aprova uma solicitação de débito
  PUT `localhost:80/auto-debit/:id/approve`
- Rejeita uma solicitação de débito
  PUT `localhost:80/auto-debit/:id/reject`
- Deleta uma solicitação de débito
  DELETE `localhost:80/auto-debit/:id`

---

# Problemática da proposta

Olá! 🖖🏽

Nossa intenção é, através deste (breve) desafio, avaliar a habilidade técnica percebida ao empregar e desenvolver uma solução para o problema aqui descrito.

## Domínio Problema

Uma instituição financeira contratou os serviços da T10 buscando maior **agilidade dos dados** através da metrificação de processos que, até então, não eram _observados_ (apropriadamente). Um dos processos é a solicitação do produto débito automático de empresas parceiras.
A operação é realizada manualmente e vai ser automatizada por este serviço, que vai permitir que outros serviços consumam, de forma livre, de seus eventos operacionais.

# Escopo

## Casos de Uso

- [x] 1. Autenticação e acesso a plataforma

Um usuário autenticado,

- [x] 2. solicita uma ativação de débito automático
- [x] 3. cancela uma solicitação de ativação
- [x] 4. aprova uma solicitação de ativação
- [x] 5. rejeita uma solicitação de ativação
- [x] 6. visualiza uma solicitação

Diagrama do [modelo de eventos](img/model.jpg).

Observações **importantes** sobre o modelo:

- É uma representação do domínio _exclusivamente_.

- Não é mandatório ser modelado usando CQRS nem event-driven.

- Não é mandatório implementar o EmailServer

## Requisitos

Especifica o contexto em que a aplicação será operacionalizada

### Não funcionais

1. 30 empresas parceiras
2. 5000 usuários simultâneos
3. 100 reqs/s

### Funcionais

#### Tecnologias

- implementação: `golang | elixir | python`
- armazenamento: `postgres | mongodb`
- **não-mandatório** broker: `kafka | rabbitmq`

#### Protocolos

- pontos de entrada: `http`
- autenticação: `simple jwt`

#### Padrões

Bonus points:

- arquitetural: `cqrs & hexagonal`
- design: `ddd & solid`
- message bus as stream

### 3rd parties

O uso de bibliotecas externas é **livre**.

### Deployment

A forma como a aplicação será disponibilizada é **livre**. Fica a critério do candidato, por exemplo, usar algum PaaS a fim de reduzir a complexidade bem como utilizar receitas prontas através de ferramentas de automatização e.g. `ansible+dockercompose`.

No entanto, é esperado bom senso na documentação caso sejam usadas soluções @ `localhost`.

# Entrega

A _Release_ 0.1 🚀 consiste na implementação de um servidor web que implementa os casos de uso listados acima respeitando os requisitos funcionais e não funcionais. Fica a critério do desenvolvedor como os testes serão escritos, os scripts de _data migration_, os _schemas_ de entrada e saída da api e todas as outras definições que não foram listadas neste documento.

## Avaliação

Critérios ordenados por ordem de peso decrescente:

1. Correção (_correctness_) da solução

   - a fim de solucionar o [domínio-problema](#domínio-problema)
   - a fim de cumprir os [casos de uso](#casos-de-uso)
   - ao implementar os [requisitos](#requisitos) especificados

2. Testes
3. Organização, documentação e clareza na estruturação do projeto
4. Estilo, legibilidade e simplicidade no código
5. Escolhas e uso de 3rd parties
6. Padrões de segurança

#### Bonus points 🏆

1. Teste de stress
2. Boas práticas na modelagem e armazenamento de dados

## Eliminatórios

1. Copiar ou "se inspirar" em código alheio é _veementemente_ vetado ✋

## Submissão

Ao finalizar a implementação, o diretório da solução pode ser submetido de duas formas:

1. através de um _fork_ e um _pull request_ neste repositório ou
2. por email, compactado, para `it@t10.digital` com o assunto `Backend Assessment`

Feito 🤘

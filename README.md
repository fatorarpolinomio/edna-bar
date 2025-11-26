<img width="100%" height="360px" alt="home_banner" src="https://github.com/user-attachments/assets/eeddc485-bf9a-4fc1-9600-8b438dd0bb98" />

# E.D.N.A - Ecossistema D. Negociação Alcoólica

Sistema de banco de dados de Bar. Gerencie produtos, atendimentos e outras atividades do seu bar com nosso sistema.

> Esse projeto é parte do projeto final para a disciplina *Banco de Dados I*.

## Funcionalidades

1. Gerenciar o estoque, atualizando o status dos itens, seus preços e disponibilidade.
2. Buscar e filtrar itens do estoque por critérios como tipo, data de validade e fornecedor.
3. Registrar, consultar e atualizar transações de venda, gastos e cobranças.
4. Gerenciar o cadastro e o saldo devedor dos clientes.
5. Registrar e consultar gastos operacionais com recursos não comerciais, como utensílios e mobília.
6. Gerar e recuperar relatórios financeiros que apresentem ganhos, gastos e projeções para um período arbitrário.
7. Gerenciar (CRUD) o cadastro de funcionários, atribuindo cargos e outros valores.
8. Consultar e atualizar os dados de pagamento, salários e bonificações de um funcionário específico.
9. Gerar e recuperar relatórios de folha de pagamento para uma categoria de funcionários ou para a equipe inteira.
10. Gerenciar (CRUD) a criação de regras de desconto e combos promocionais.
11. Buscar e filtrar promoções por critérios como popularidade, lucratividade e data de criação.


## Tecnologias Utilizadas

A API _backend_ foi criada utilizando a linguagem **Go** (1.24). O servidor HTTP foi desenvolvido usando inteiramente a biblioteca padrão `net/http`. Enquanto o banco de dados utilizado é um banco **PostgreSQL**. Já no _frontend_, foi utilizado o framework **Vue.js** e **CSS** puro. Por último, a aplicação foi conteinerizada através de **Docker** e exposta por _reverse proxy_ **Nginx**.

## Pré-requisitos
- Docker & Docker Compose
- PostgreSQL _Opcionalmente_ (caso prefira não usar o Docker)
- Node.js +22
- Go +1.24
- Go-migrate para as migrações.

O último requisito pode ser instalado com:
```sh
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.3
```

## Como rodar

Para rodar em modo de desenvolvimento.

Antes de mais nada é preciso ter o Docker instalado com o Docker Compose. Ou você pode configurar a base de dados Postgres manualmente.
Crie um arquivo `.env` com variáveis configuradas semelhantes ao arquivo `.env.example`. Também crie um arquivo `.env` dentro da pasta frontend com a variável `VITE_BACKEND_BASE_URL=http://localhost:8080/v1` para o frontend acessar o back em modo de desenvolvimento.
Os arquivos `.env.example` contém uma explicação de cada variável, **não exponha-as em produção**.

Feito isso, inicie o container da base de dados utilizando o comando `docker compose up -d database` ou com o _Make_, com `make docker-run` (derrube o container com `make docker-down`). 

> O arquivo `Makefile` contém vários comandos simples para rodar o projeto, veja mais em `make help`.

Na primeira vez (ou em modificações ao banco), popule ou atualize a base de dados com o script `migrate.sh`.
```sh
./migrate.sh up
```
Não esqueça de configurar a variável `DB_URL` dentro arquivo com o caminho para sua base de dados.

Para rodar o **backend** use o compile o projeto em go ou use o comando `make run`. Para rodar em modo de _hotreloading_ use:
```sh
make watch
```

Já para o **frontend**, entre na pasta `frontend` e execute os seguintes comandos:

1. Para baixar as dependências:
```sh
npm install
```
2. Para rodar o projeto:
```sh
npm run dev
```

## Fazendo o Deploy

A aplicação foi inteiramente conteinerizada para rodar em produção. Configure as variáveis de ambiente corretamente e **não exponha segredos**. Crie um conjunto de certificados, ou através da um provedor e os armazene na pasta `nginx/conf.d/certs`. Se necessário mude os nomes dos certificados em `nginx/conf.d/site.conf` (ou mantenha os nomes atuais: `origin.pem` e `origin.key`).

Para o deploy, rode o comando:

```sh
docker compose up -d --build
```

Para as migrações, um comando único do docker foi adicionado:
```sh
docker compose --profile migrate run --rm migrate
```
Isso deve executar as migrações na base de dados.

> Atualize o script `migrate.sh` com `DB_URL` antes de rodar esse comando.

## Sobre migrações

> Migrações são scripts SQL que são rodados na base de dados e permitem criar um histórico de alterações e navegar por elas.

Não esqueça de acionar a base de dados antes de rodar migrações e de atualizar o script em `migrate.sh` com o endereço, nome de usuário e senha corretos para acessar o banco.

As migrações vivem em `migrations`. Utilize o script `migrate.sh` para gerenciar migrações. Crie novas migrações com `./migrate.sh create <migration_name>`. Vá para a migração mais recente com `./migrate.sh up`, volte **uma** migração com `./migrate.sh down 1`. Veja mais comandos em `./migrate.sh help`. 

A ferramenta de migrações utilizada é o [go-migrate](https://github.com/golang-migrate/migrate).

## Diagrama Conceitual

Segue um Diagrama Entidade Relacionamento da Aplicação em Banco de Dados.

![Diagrama Entidade-Relacionamento do Sistema EDNA bar](assets/E.D.N.A_conceitual_1.jpeg)

## Licensa

Esse projeto está sob a [GNU General Public License v3.0](https://spdx.org/licenses/GPL-3.0-or-later.html). O uso e distribuição é permitido desde que siga as restrições impostas pela [LICENSA](./LICENSE).

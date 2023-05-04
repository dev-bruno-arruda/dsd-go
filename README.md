Para resolver o problema proposto, vamos criar um CRUD em Go e MySQL para cadastrar usuários, vessels e flags com rotas em REST API, utilizando o framework Fiber para gerenciamento das rotas e Mux para multiplexação, seguindo o padrão de arquitetura DDD (Domain Driven Design) e utilizando Docker com Docker Compose para facilitar a execução do projeto.

projeto/
├── app/
│   ├── api/
│   │   ├── handlers/
│   │   ├── routes/
│   │   ├── middleware/
│   ├── config/
│   ├── domain/
│   │   ├── user/
│   │   ├── vessel/
│   │   ├── flag/
│   ├── infra/
│   │   ├── persistence/
│   ├── usecase/
├── docker-compose.yml
├── Dockerfile
├── README.md

mkdir app/
mkdir app/api
mkdir app/api/handlers/
mkdir app/api/routes/
mkdir app/api/middleware/
mkdir app/config/
mkdir app/domain/
mkdir app/domain/user
mkdir app/domain/vessel
mkdir app/domain/flag
mkdir app/infra/
mkdir app/infra/persistence
mkdir app/usecase/


No código acima, criamos os serviços para usuários, embarcações e bandeiras. Esses serviços encapsulam as regras de negócio e orquestram as operações no banco de dados.

No serviço de usuário, utilizamos a biblioteca `bcrypt` para criptografar a senha antes de salvá-la no banco de dados. Também criamos um método `Authenticate` para autenticar o usuário com base no e-mail e na senha.

No serviço de embarcações, adicionamos um método `validateFlag` para verificar se a bandeira atribuída a uma embarcação existe antes de criar ou atualizar a embarcação.

No serviço de bandeiras, adicionamos uma validação para não permitir a exclusão de uma bandeira que esteja atribuída a uma embarcação. Para fazer isso, primeiro buscamos todas as embarcações que possuem essa bandeira e, se houver alguma, retornamos um erro.

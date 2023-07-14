# Invoice Exchange
![test](https://github.com/tauki/invoiceexchange/actions/workflows/test.yml/badge.svg)

### Requirements
- Go 1.18 or later

### Installation
- Clone the repository
- Run `go mod download` to download the dependencies
- Run `go build` to build the binary
- optionally pass `-o <binary_name>` to specify the name of the binary
- Run `./<binary_name>` to start the server.
- You might need to run `chmod +x <binary_name>` to make the binary executable.
- Alternatively, you can run `go run main.go` to start the server

### Usage
- The server will start on `port 8080` by default, it can be changed by setting the `SERVE_PORT` environment variable.
- If an environment variable `TLS` is set to `true`, the server will start with TLS enabled and will look for the following environment variables:
    - `CERT_PATH`: The path to the certificate file, by default its set to `invoiceexchange.local.pem`
    - `CERT_PRIVATE_KEY`: The path to the key file, by default its set to `invoiceexchange.local-key.pem`
    - The TLS port is set to `443` by default, it can be changed by setting the `TLS_PORT` environment variable.
- By default, the server runs in `dev` mode, it can be changed by setting the `ENVIRONMENT` environment variable to `prod`.
- By default, the server starts looks for `POSTGRES_DSN` environment variable to connect to the database.
- In `local.yaml` the `ENT-DRIVER` environment variable is set to `sqlite` to use sqlite as the database driver. If the `ent-driver` in the yaml or the environment variable `ENT_DRIVER` is set to `pgx`, the server will use postgres as the database driver. The sqlite uses postgres dialect so the database can be easily switched.
- If the `ENT_DRIVER` is set to `sqlite`, the server will not use the `POSTGRES_DSN` environment variable to connect to the database.
- The sqlite database starts with some preloaded data for testing purposes. The data is saved in a file called `local.db` in the root directory of the project. The file can be deleted to start with a fresh database.
- The server uses `ent framework` for the database layer. The ent schema is defined in `ent/schema` directory. To install ent follow the instructions [here](https://entgo.io/docs/getting-started/). The ent schema can be generated by running `go generate tools.go` in the root directory of the project. Alternatively, you can use the `ent` binary to generate the schema. 
- To see the ent schema description run `ent describe ./ent/schema` in the root directory of the project. `go generate tools.go` print the describe after generating the schema.
- Alternatively, a generated ent schema describe has been provided at `tools/dump_ddl/ENT_DESCRIBE`. 

### Running with Postgres
Please find the initial migration dump located in `tools/dump_ddl/initial_dump`
- Create a database in postgres.
- Set the `POSTGRES_DSN` environment variable to your postgres dsn in the format of `user=user password=password host=host port=port dbname=dbname sslmode=disable`
- Alternatively update the `PostgresDSN` variable in `tools/dump_ddl/main.go` and run `go run tools/dump_ddl/main.go` to print the migration dump, it will output the dump in `stdout`.
- Run the migration dump in the database to create the tables.
- If running with `dev` environment, set the `ent-driver` variable to `pgx` to use the postgres driver instead of the sqlite driver.
- Alternatively set the `ENT_DRIVER` environment variable to `pgx` to use the postgres driver.
- Update the `postgres-dsn` in `local.yaml` with the correct values.

### Testing
- Run `go test ./...` to run all unit tests
- Run `go test -tags=e2e ./...` to run test with all end-to-end scenarios
- Alternatively `cd handlers && go test -tags=e2e ./...` to run just the end-to-end scenario tests for the handlers

### Project structure
 - `config`: Contains the configuration files for the server
 - `ent`: Contains the ent schema and generated ent code
 - `eventhandler`: Contains the event handlers for the server. Currently, it only contains the handler for `bid.created` event.
 - `handlers`: Contains the HTTP handlers/ controllers for the server
 - `internal`: Contains the internal packages, services, and types
 - `repos`: Contains the repository layer for the server
 - `router`: Contains the router and middleware for the server. The router is built on top of gorilla mux. The routes can be found at `router/routes.go`. The middlewares can be found inside the `router/middleware` directory.
 - `services`: Contains the services for the server, the services are responsible for the business logic of the server.
 - `tools`: Contains the tools for the server, currently it only contains the `dump_ddl` tool to generate the migration dump for the database and the postman collection json for the API.

### API
- A `postman` collection can be found [here](https://www.postman.com/tauki/workspace/public/collection/3106382-ce688021-79ec-4152-9abb-61cca0886351?action=share&creator=3106382). The collection contains all the endpoints and examples for the requests and responses.
- Alternatively, you can import the collection to postman from the `tools/invoiceexchange.postman_collection.json` file in the root directory of the project. The collection is exported with `Collection v2.1`.

### Improvements
- Improve the balance schema and the relation between balance, issuers, and investors. Currently, the `balance.entity_id` is set to unique which won't allow the issuer or investor to have multiple balance entry, but the edges are `O2M`.
- Need ledger implementation for the available balance as well, currently the ledger is only implemented for the total balance.
- Properly assert param and returned values in test.
- More end to end scenarios and unit test cases should be tested to ensure that the system is working as expected.
- A better implementation for the eventbus.
- Proper setup and teardown for the end-to-end tests.
- Invoice items are currently not being saved, the invoice items should be saved in the database and the invoice should reference the invoice items.

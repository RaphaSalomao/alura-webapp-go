
# Alura-WebApp-Go

This is a simple web application made with Golang's standard library


## Autores

- [@RaphaSalomao](https://www.github.com/RaphaSalomao)


## Deploy

In order to run this project, you need a PostgreSQL instance running on localhost on port 5433, you may do this using docker:
```bash
  docker pull postgres
```
```bash
  docker run -d --name dev-postgres -e POSTGRES_PASSWORD=root -v ${HOME}/postgres-data/:/var/lib/postgresql/data -p 5433:5432 postgres
```
Note that you can change the database connection options if you also change `datasource` const at `/database/conector.go`

After that you need to run product table DDL query on the database and import uuid-ossp functions:
```sql
  CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

  CREATE TABLE product(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL,
    description VARCHAR,
    price DECIMAL NOT NULL,
    quantity INTEGER NOT NULL
  );
```
And finally you may run `go run main.go` in the project root and access the index page at `localhost:8000/`
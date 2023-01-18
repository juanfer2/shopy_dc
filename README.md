# Product dd Services
App for test works with products

## Config
Previusly to setup the project it's necesary isntall
 - Postgres
 - Make

Then next commands

`
make copy-config
`

This command generate diferents files in `src/shared/infrastructure/persistence/postgres/config/database_development.yml` in which you must add your credentials postgrest

```yml
drive: postgres
username: username
password: pasword
host: localhost
port: 5432
database: 'shopy_dc_dev'
```
Now run next command for install go dependencies

`
go mod tidy
`

`
go install
`

Then you can create a database

`
go run ./main.go db:create
`

Finally run this command for fill database with data

`
go run ./main.go db:seed products
`

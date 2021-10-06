1) получить репу Postgres 
- ```docker pull postgres``` 
2) развернуть контейнер с Postgres 
- ```docker run --name=todo-db -e POSTGRES_PASSWORD='password' -p 5432:5432 -d --rm postgres```
3) используется пакет Go - migrate
- ```go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest```
4) создание миграций
- ```migrate create -ext sql -dir ./schema -seq init```
5) запуск миграций
- ```migrate -path ./schema -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' up```
password - пароль к учетной записи в Postgres

___

1) зайти в контейнер Postgres
- ```docker exec -it container_id bin/bash```
2) зайти в psql модуль
- ```psql -U postgres``` postgres - имя учетной записи Postgres
3) для апдейта миграции
- ```update schema_migrations set version='000001', dirty=false;``` где 000001 название миграции, а dirty 
# Makefile для создания миграций

# Переменные, которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down
	
# для удобства добавим команду run, которая будет запускать наше приложение
run:
	go run cmd/app/main.go # Теперь при вызове make run мы запустим наш сервер

gen:
	mkdir -p terminal/web/tasks
	mkdir -p terminal/web/users
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./terminal/web/tasks/api.gen.go
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./terminal/web/users/api.gen.go

lint:
	golangci-lint run 



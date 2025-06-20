# Makefile для создания миграций

# Переменные которые будут использоваться в наших командах (Таргетах)
DB_DSN := postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
migrate-new-tasks:
	migrate create -ext sql -dir ./migrations tasks

migrate-new-users:
	migrate create -ext sql -dir ./migrations users

migrate-new-tasks-add-user-id:
	migrate create -ext sql -dir ./migrations tasks_add_user_id

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# для удобства добавим команду run, которая будет запускать наше приложение
run:
	go run cmd/main.go # Теперь при вызове make run мы запустим наш сервер

gen-tasks:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

gen-users:
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

gen: gen-tasks gen-users

lint:
	golangci-lint run --color=always
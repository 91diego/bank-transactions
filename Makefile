APP_BINARY=bankApp

## starts all containers in the background without forcing build
up:
	@echo "Starting Docker Images"
	docker-compose up -d
	@echo "Docker Images started!!"

up_build:
	@echo "Stoping Docker Images (if running)..."
	docker compose down
	@echo "Building (when required) and starting docker images ..."
	docker compose up --build -d
	@echo "Docker compose built and started!"

# down: stop docker compose
down:
	@echo "Stopping Docker compose ..."
	docker compose down
	@echo "Done!"

# down: stop docker and remove orphans from docker compose
down_remove_orphans:
	@echo "Stopping Docker compose ..."
	docker-compose down --remove-orphans
	@echo "Done!"

## build_broker: builds the app binary as a linux executable
build_app:
	@echo "Building app binary..."
	cd . && env GOOS=linux CGO_ENABLED=0 go build -o ${APP_BINARY} .
	@echo "Done!"

migration_up:
	@echo "Starting migrations"
	docker run --network host bank-transactions-migrator-bank -path=/migrations/ -database "postgresql://bank:bank@localhost:7567/bank?sslmode=disable" up
	@echo "DONE!!!"

migration_down:
	@echo "Starting migrations"
	docker run --network host bank-transactions-migrator-bank -path=/migrations/ -database "postgresql://bank:bank@localhost:7567/bank?sslmode=disable" down $(filter-out $@,$(MAKECMDGOALS))
	@echo "DONE!!!"

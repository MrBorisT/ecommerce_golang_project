
build-all:
	cd checkout && GOOS=linux make build
	cd loms && GOOS=linux make build
	cd notifications && GOOS=linux make build

run-all: build-all
	sudo docker compose up --force-recreate --build

precommit:
	cd checkout && make precommit
	cd loms && make precommit
	cd notifications && make precommit

down:
	docker-compose down

# db stuff
up-db:
	docker-compose up -d postgres-checkout
	docker-compose up -d postgres-loms

stop-db:
	docker-compose stop postgres-checkout
	docker-compose stop postgres-loms

start-db:
	docker-compose start postgres-checkout
	docker-compose start postgres-loms

#	kafka
up-kafka:
	docker-compose up -d zookeeper kafka1 kafka2 kafka3
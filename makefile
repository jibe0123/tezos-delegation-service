.PHONY: all build run clean full-clean

all: build run

build:
	docker-compose build --no-cache

run: clean
	docker-compose up --build --force-recreate

clean:
	docker-compose down --volumes --rmi all
	-docker volume rm tezos-delegation_db-data || true

full-clean: clean
	-docker rmi -f $(docker images -q) || true
	-docker volume prune -f
	-docker system prune -af --volumes

rebuild:
	docker-compose build --no-cache
	docker-compose up --build --force-recreate

clean-images:
	-docker rmi -f $(docker images -q) || true
	-docker volume prune -f
	-docker system prune -af --volumes
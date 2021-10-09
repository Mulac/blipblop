.PHONY = up down build deploy

up:
	docker-compose -p blipblop -f cmd/docker-compose.yaml up -d --build

down:
	docker-compose -p blipblop -f cmd/docker-compose.yaml down

build:
	docker-compose -p blipblop -f cmd/docker-compose.yaml build

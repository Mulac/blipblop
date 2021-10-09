.PHONY = backend

up:
	docker-compose -p blipblop -f cmd/docker-compose.yaml up -d

down:
	docker-compose -p blipblop -f cmd/docker-compose.yaml down

build:
	docker-compose -p blipblop -f cmd/docker-compose.yaml build

backend:
	docker build -f cmd/backend/Dockerfile -t blipblop/backend:latest .
build:
	docker build . -t fiber-api:latest

compose:
	make build
	docker-compose up

scale:
	docker-compose scale book-api-go-fiber=3
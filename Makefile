build:
	docker build . -t fiber-api:latest

compose:
	docker-compose up --scale book-api-go-fiber=3 --scale book_db_go=1
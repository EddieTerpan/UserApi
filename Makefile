start-cli:
	docker-compose build && docker-compose up
start:
	docker-compose build && docker-compose up -d
stop:
	docker-compose down
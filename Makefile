start-dev:
	docker-compose build && docker-compose up
start:
	docker-compose up -d
stop:
	docker-compose down
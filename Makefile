.PHONY: up
up:
	docker-compose up -d

.PHONEY: down
down:
	docker-compose down
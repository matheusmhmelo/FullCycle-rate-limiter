build:
	docker build -t ratelimiter .

run:
	docker-compose build && docker-compose up

.PHONY: build run
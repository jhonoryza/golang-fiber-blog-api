dev:
	/Users/fajar/go/bin/air .
build:
	docker build -t app .
run:
	docker compose up -d
update:
	git pull origin main && make build && make run
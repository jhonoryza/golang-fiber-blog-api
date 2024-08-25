dev:
	go run cmd/main.go
build:
	docker build -t jhonoryza/fiber_blog .
run:
	docker compose up -d
update:
	git pull origin main && make build && make run
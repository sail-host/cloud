.PHONY: run

run:
	cd web && bun run dev & ~/go/bin/air && fg

dev: build-web dev-api

build-web:
	cd web && bun run build

dev-api:
	@go run cmd/main.go app --dev

build:
	@echo "Building web..."
	cd web && bun run build
	@echo "Building server..."
	go build -o tmp/main cmd/web/main.go
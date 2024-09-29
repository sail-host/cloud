.PHONY: run

run:
	cd web && npm run dev & ~/go/bin/air && fg

dev:
	cd web && npm run build && cd .. && go run cmd/main.go app --dev

dev-api:
	@go run cmd/main.go app --dev

build:
	@echo "Building web..."
	cd web && npm run build
	@echo "Building server..."
	go build -o tmp/main cmd/web/main.go
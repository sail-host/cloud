.PHONY: run

run:
	cd web && npm run dev & ~/go/bin/air && fg

dev:
	go run cmd/main.go app --dev

build:
	@echo "Building web..."
	cd web && npm run build
	@echo "Building server..."
	go build -o tmp/main cmd/web/main.go
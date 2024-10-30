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

prod:
	@echo "Build Mac AMD64"
	GOOS=darwin GOARCH=amd64 go build -o sailhost-darwin-amd64 cmd/main.go
	tar -czvf "sailhost-darwin-amd64.tar.gz" sailhost-darwin-amd64

	@echo "Build Mac ARM64"
	GOOS=darwin GOARCH=arm64 go build -o sailhost-darwin-arm64 cmd/main.go
	tar -czvf "sailhost-darwin-arm64.tar.gz" sailhost-darwin-arm64

	@echo "Build Linux AMD64"
	GOOS=linux GOARCH=amd64 go build -o sailhost-linux-amd64 cmd/main.go
	tar -czvf "sailhost-linux-amd64.tar.gz" sailhost-linux-amd64

	@echo "Build Linux ARM64"
	GOOS=linux GOARCH=arm64 go build -o sailhost-linux-arm64 cmd/main.go
	tar -czvf "sailhost-linux-arm64.tar.gz" sailhost-linux-arm64
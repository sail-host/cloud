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
	GOOS=darwin GOARCH=amd64 go build -o sailhost cmd/main.go
	tar -czvf "sailhost-mac-amd64-${APP_VERSION}.tar.gz" sailhost 

	@echo "Build Mac ARM64"
	GOOS=darwin GOARCH=arm64 go build -o sailhost cmd/main.go
	tar -czvf "sailhost-mac-arm64-${APP_VERSION}.tar.gz" sailhost

	@echo "Build Linux AMD64"
	GOOS=linux GOARCH=amd64 go build -o sailhost cmd/main.go
	tar -czvf "sailhost-linux-amd64-${APP_VERSION}.tar.gz" sailhost

	@echo "Build Linux ARM64"
	GOOS=linux GOARCH=arm64 go build -o sailhost cmd/main.go
	tar -czvf "sailhost-linux-arm64-${APP_VERSION}.tar.gz" sailhost

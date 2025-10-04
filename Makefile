run: 
	@APP_ENV=development go run ./cmd/main.go

build:
	@go build -o bin/tt-shopee ./cmd/main.go

prod:
	@APP_ENV=production ./bin/tt-shopee
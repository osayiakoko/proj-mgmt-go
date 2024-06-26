run: build
	@./bin/project-mgmt-sys


build: 
	@wgo build -o bin/project-mgmt-sys cmd/api/main.go

runserver: 
	@wgo run ./cmd/api

test:
	@go test -v ./...
MNAME = migration

run: build
	@./bin/project-mgmt-sys


build: 
	@wgo build -o bin/project-mgmt-sys cmd/api/main.go

runserver: 
	@wgo run ./cmd/api

genent:
	@go generate ./ent

test:
	@go test -v ./...

migdiff:
	@atlas migrate diff $(MNAME) \
		--dir file://migrate/migrations \
		--to file://migrate/schema.sql \
		--dev-url "docker://postgres/14.2-alpine/test?search_path=public" \
		--format '{{ sql . "  " }}'

migapply:
	@atlas migrate apply \
		--dir file://migrate/migrations \
		--url "postgres://projmgmt:projmgmt@localhost:5436/projmgmt?sslmode=disable"


migstatus:
	@atlas migrate status \
		--dir file://migrate/migrations \
		--url "postgres://projmgmt:projmgmt@localhost:5436/projmgmt?sslmode=disable"

mighash:
	@atlas migrate hash --dir file://migrate/migrations

miglint:
	@atlas migrate lint \
		--dir file://migrate/migrations \
		--dev-url "docker://postgres/14.2-alpine/test?search_path=public" \
		--latest 1

migdown:
	@atlas migrate down \
		--dir file://migrate/migrations \
		--url "postgres://projmgmt:projmgmt@localhost:5436/projmgmt?sslmode=disable" \
		--dev-url "docker://postgres/14.2-alpine/test?search_path=public"

migdownver:
	@atlas migrate down \
		--dir file://migrate/migrations \
		--url "postgres://projmgmt:projmgmt@localhost:5436/projmgmt?sslmode=disable" \
		--dev-url "docker://postgres/14.2-alpine/test?search_path=public" \
		--to-version $(MVER)

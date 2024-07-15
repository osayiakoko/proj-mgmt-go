## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

MNAME = migration

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]


## run: run the application binary
.PHONY: run
run: build
	@./bin/project-mgmt-sys

## build: build the applicaton binary
.PHONY: build
build: 
	@go build -o bin/project-mgmt-sys cmd/api/main.go

## runserver: run the cmd/api application in watch mode
.PHONY: runserver
runserver: 
	@wgo run ./cmd/api

.PHONY: genent
genent:
	@go generate ./ent

## test: runs test
.PHONY: test
test:
	@go test -v ./...

## migdiff: create db migration
.PHONY: migdiff
migdiff:
	@atlas migrate diff $(MNAME) \
		--dir file://migrate/migrations \
		--to file://migrate/schema.sql \
		--dev-url "docker://postgres/14.2-alpine/test?search_path=public" \
		--format '{{ sql . "  " }}'

## migapply: apply database migration
.PHONY: migapply
migapply: confirm
	@atlas migrate apply \
		--dir file://migrate/migrations \
		--url "postgres://projmgmt:projmgmt@localhost:5436/projmgmt?sslmode=disable"

## migapplyprod: apply database migration for prod environment
.PHONY: migapplyprod
migapplyprod:
	@atlas migrate apply \
		--dir file://migrate/migrations \


## migstatus: checks db migration status
.PHONY: migstatus
migstatus:
	@atlas migrate status \
		--dir file://migrate/migrations \
		--url "postgres://projmgmt:projmgmt@localhost:5436/projmgmt?sslmode=disable"

## mighash: hashes database migration file
.PHONY: mighash
mighash:
	@atlas migrate hash --dir file://migrate/migrations

## miglint: lints pending db migration
.PHONY: miglint
miglint:
	@atlas migrate lint \
		--dir file://migrate/migrations \
		--dev-url "docker://postgres/14.2-alpine/test?search_path=public" \
		--latest 1

## migdown: 
.PHONY: migdown
migdown:
	@atlas migrate down \
		--dir file://migrate/migrations \
		--url "postgres://projmgmt:projmgmt@localhost:5436/projmgmt?sslmode=disable" \
		--dev-url "docker://postgres/14.2-alpine/test?search_path=public"

## migdownver: 
.PHONY: migdownver
migdownver:
	@atlas migrate down \
		--dir file://migrate/migrations \
		--url "postgres://projmgmt:projmgmt@localhost:5436/projmgmt?sslmode=disable" \
		--dev-url "docker://postgres/14.2-alpine/test?search_path=public" \
		--to-version $(MVER)

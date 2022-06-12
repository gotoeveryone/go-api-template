# Golang API Template

## Requirements

- Docker

## Run

```console
$ cp .env.example .env
$ docker compose up
```

## Format check

```console
$ docker compose exec app go vet -v ./...
```

## Build

```console
$ docker compose exec app go build src/main.go
```

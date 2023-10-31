# Kredit Plus

## Clone
```bash
$ git clone https://github.com/muhammadaskar/kredit-plus.git
```

## Install Packages
```bash
$ go mod download
```

## Copy Environtment
```bash
$ cp .env.example .env
```

## Run Migration
```bash
$ go run cmd/migration-runner/main.go
```

## Run Seeder
```bash
$ go run cmd/seeder-runner/main.go
```

## Build Using Docker Compose
```bash
$ docker compose up -d
```

## Build Using Dockerfile
```bash
$ docker build --platform=linux/amd64 --network=host -t kredit-plus-image:v1.0 .
```
```bash
$ docker run -d -p 5000:5000 --name kredit-plus-container kredit-plus-image:v1.0
```

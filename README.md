[![Go Reference](https://pkg.go.dev/badge/github.com/201R/remote.svg)]() [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# Golang Restapi Boilerplate

## Description

Golang REST API boilerplate for typical project, including interactive API documentation.

## Table of Contents

-   [Interactive API documentation](#interactive-api-documentation)
-   [Features](#features)
-   [Quick run](#quick-run)
-   [How to use it](#how-to-use-it)
-   [Comfortable development](#comfortable-development)
-   [Links](#links)
-   [Database](#database)
-   [Tests](#tests)
-   [License](#license)

## Interactive API documentation

[![API docs](.github/swagger.png)](https://github.com/201R/golang_restapi_boilerplate)

## Features

-   [x] Database ([entgo](https://github.com/ent/ent)).
-   [x] Migration ([atlas](https://github.com/ariga/atlas)).
-   [x] Config
-   [ ] Admin and User roles. ([Casbin Rbac](https://github.com/casbin/casbin))
-   [x] Mailing
-   [ ] Sign in and sign up via email.
-   [ ] Redis
-   [ ] Social sign in (Apple, Facebook, Google, Twitter).
-   [ ] I18N
-   [ ] File uploads. Support local and Amazon S3 drivers.
-   [x] Swagger.
-   [ ] nginx.
-   [x] Adminer([Adminer](https://www.adminer.org/))
-   [ ] Docker.
-   [ ] CI (Github Actions).
-   [ ] Monitoring ([Grafana](https://grafana.com/docs/loki/latest/api/))
-   [ ] Unit Test([Testify](https://github.com/stretchr/testify))

## Quick run

```bash
git clone --depth 1 https://github.com/201R/golang_restapi_boilerplate.git my-app
cd my-app/
cp .ini.exemple .ini
docker-compose up -d
```

For check status run

```bash
docker-compose logs
```

## How to use it

## Comfortable development

-   Install Gin Hot reload

Assuming you have a working Go environment and `GOPATH/bin` is in your
`PATH`, `gin` is a breeze to install:

```bash
go get github.com/codegangsta/gin
```

-   Pull code

```bash
git clone --depth 1 https://github.com/201R/golang_restapi_boilerplate.git my-app
cd my-app/
cp .ini.exemple .ini
```

-   Update .ini file

change `Host=postgres` to `Host=localhost`

<!-- Change `MAIL_HOST=maildev` to `MAIL_HOST=localhost` -->

-   Run additional container

```bash
docker-compose up -d postgres adminer redis
```

-   Start App

```bash
go mod download

go run -mod=mod ariga.io/atlas/cmd/atlas@master migrate apply \
  --dir "file://migrations"
  --url "postgres://postgres:password@localhost:6001/db?sslmode=disable"

gin -a 3000 -p 8090 -i
```

## Links

-   Swagger: http://localhost:8090/swagger/index.html
-   Adminer: http://localhost:8080/

## Database

## License

This project is licensed under the terms of the MIT license.

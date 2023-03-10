# Go-Scrapper

Go-Scrapper is a HTTP Restful API for crawling web application.

## Installation

Clone from Github repository

```bash
git clone git@github.com:anacondaf/go-scrapper.git
```

## Setup

```bash
# install all dependencies
go mod tidy
```

## Usage
To run and use this repository, follow these steps:

* Running Postgres server
* Create app.env file and copy all contents from .env.example
* Run by:

```bash
# Using golang air
air

# Using golang run
go run main.go

```

## Features
- [x] Written in Golang
- [x] HTTP Framework - Go Fiber
- [x] Supports PSQL, MongoDB
- [x] Crawler Framework - Go Colly
- [x] OpenAPI - Supports Client Service Generation

<details>
    <summary>Click to See More!</summary>

- [x] Response Caching - Go Redis
- [x] Scheduler Library - Cdule
- [x] Cloud storage - AWS S3
- [x] API Versioning
</details>

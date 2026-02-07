# URL Shortener (In-Memory)

## Getting started

1. `git clone git@github.com:reneroboter/urlshortener.git`
2. `cd urlshortener`
3. `docker build --tag urlshortener .`
4. `docker run --publish 8888:8888 urlshortener`

## Run tests locally
1. `go build -v ./...`
2. `go test -v ./...`

## API Specification

### GET [/{hashedUrl}]
+ Parameters
  + code: `8ece61d2d42e578e86d9f95ad063cf36eb8e774d` (string, required)
+ Response 301 (text/html)
  + Headers
    Location: http://www.example.com
  + Body:  <a href="http://www.example.com">Moved Permanently</a>.
+ Response 400 (text/plain)
  + Body: invalid code
+ Response 404 (text/plain)
  + Body: not_found

#### Examples

1. `curl -i http://localhost:8888/8ece61d2d42e578e86d9f95ad063cf36eb8e774d`
2. `curl -i http://localhost:8888/asd`

### POST [/shorten]
+ Request (application/json)
    + Body: ```{ "url": "http://www.google.de" }```
+ Response 201 (application/json)
    + Body:  ```{ "code": "..." }```
+ Response 400 (text/plain)
    + Body: invalid URL format
+ Response 400 (text/plain)
    + Body: json decoding error
+ Response 409 (text/plain)
    + Body: url already exists

#### Examples

1. ```bash
   curl -i -X POST http://localhost:8888/shorten \
    -H "Content-Type: application/json" \
    -d '{"url":"http://www.google.de"}'
    ```
2. ```bash
   curl -i -X POST http://localhost:8888/shorten \
    -H "Content-Type: application/json" \
    -d '{"url":"www.google.de"}'
    ```



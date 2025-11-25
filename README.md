# URL Shortener (In-Memory)

## Description

I've asked ChatGPT for a mini project idea in Go. Creating an URL Shortener was one of its recommendations, and I thought it would be fun to try it out.
And I'm happy to report that it worked out pretty well.

## What you’ll learn

- HTTP server (net/http)
- Structs + maps
- JSON encoding
- Routing basics

## Tasks

- POST `/shorten` → returns a short ID
- GET `/{id}` → redirects to original URL
- Store everything in a map[string]string

## Getting started

1. `git clone git@github.com:reneroboter/urlshortener.git`
1. `cd urlshortener`
1. `go run main.go`

## How to use?

### Create record

1. `http POST localhost:8888/shorten url=http://www.google.de` -> Returns a short ID (Hashed URL)
1. `http POST localhost:8888/shorten url=www.google.de` -> Returns invalid URL error
1. `http POST localhost:8888/shorten url=http://www.google.de`-> (Second post) Returns bad request error

### Fetch record

1. `http localhost:8888/asd` -> Returns bad request invalid id
1. `http localhost:8888/eb43b895f40fbc0f0bdda29d3d52e58a53e2b4b8` -> Returns redirect to target url
1. `http localhost:8888/129c0d99c6fca772c7a007844b8b71a9097d9915` -> Returns not found error

## Questions 

- Memory management -> What is if the process reached his memory_limit in Go?
- What are nil errors?

## Ideas
- analyze execution time, memory usage and CPU usage

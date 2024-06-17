##  Overview

This project is a hands-on exploration of building a command-line interface (CLI) application using the Go programming language. It's designed as a learning exercise, focusing on best practices and core concepts.

## Key Features & Learning Objectives

* **CLI Structure:** Learn how to design and implement a well-structured CLI application with Go, including easy addition of new commands.
* **Separation of Concerns:** Understand how to organize your Go code into packages for better maintainability. This project has two extra packages for handling API calls and cache.
* **API Interactions:** Gain experience making API calls in Go and handling responses.
* **Caching (In-Memory):** Implement an in-memory cache to reduce API call frequency and improve response times.
* **Goroutines for Cache Eviction:** Explore the use of goroutines for background tasks like cache eviction.
* **Unit Testing:**  Get familiar with Go's built-in testing framework by writing unit tests. 

## Getting Started

### Instalation

> This project has no dependencies. Just clone it and run:
>```console
>$ go build -o pokedex
>```

###  Usage

> See a list of available commands:
>```console
>$ ./pokedex help
>```

### Tests

> Run the test suite using the command below:
>```console
>$ go test
>```





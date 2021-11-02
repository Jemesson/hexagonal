# Hexagonal architecture (Ports and Adapters) [![Go](https://github.com/Jemesson/hexagonal/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/Jemesson/hexagonal/actions/workflows/go.yml)

The application is based from a [Go article](https://medium.com/@matiasvarela?p=cfd4e436faa3).
Changed the app with some refactorings, and to make the understanding easier of each component.

![alt text](https://reflectoring.io/assets/img/posts/spring-hexagonal/hexagonal-architecture.png "Ports and Adapters")

## Topics:
1. [Summary](#summary)
   1. [What is a port?](#ports)
   2. [What is an adapter?](#adapters)
2. [Setup](#setup)
3. [Stack](#stack)
4. [References](#references)

## Summary
The Ports & Adapters Architecture (aka Hexagonal Architecture) was thought of by Alistair Cockburn and written down on 
his blog in 2005. This is how he defines its goal in one sentence:
> Allows an application to equally be driven by users, programs, automated test or batch scripts, and to be developed and tested in isolation from its eventual run-time devices and databases.

The main idea in this architecture is to separate the concerns, each component has a well-defined purpose.
Focus on the core business and the external services will plug and play with the application through ports and adapters.
> P.S. The hexagonal architecture is agnostic of the programming language.

### Ports
A port is a consumer agnostic entry and exit point to/from the core business.
In many languages, it will be an interface. For example, it can be an interface to perform searches in a search engine.
### Adapters
Adapters can be external services that wants to communicate with the application business logic.

> Example 1: An input adapter could be a web interface. When a user clicks a button, the web adapter calls a certain
> input port to communicate with the core.

> Example 2: An output adapter adapters are called by our core business, for instance, persisting data into a  database.

## Setup

Run the app:
> go build hexagonal/cmd/serve

Test:
> go test hexagonal/tests

## Stack
* [Go](https://golang.org/)
* [Gin](https://github.com/gin-gonic/gin)
* [Testify](https://github.com/stretchr/testify)
* [GoMock](https://github.com/golang/mock)
* [UUID](https://github.com/google/uuid)

## References
1. [Alistair Cockburn](https://alistair.cockburn.us/hexagonal-architecture/)
2. [Spring with Hexagonal](https://reflectoring.io/spring-hexagonal/)
3. [Go article](https://medium.com/@matiasvarela?p=cfd4e436faa3)
4. [FullCycle](https://fullcycle.com.br/)
5. [Go By example](https://gobyexample.com/)

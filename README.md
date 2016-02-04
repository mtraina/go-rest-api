# Simple Rest API written in Go

Simple API for executing CRUD operations on Todo.

Extension of the work made by by https://github.com/corylanou with the following tutorial:
https://thenewstack.io/make-a-restful-json-api-go/

Added unit test using [Ginkgo](http://onsi.github.io/ginkgo/) and [Gomega](http://onsi.github.io/gomega/)


## How to use it
Checkout the code and build it with

```shell
go build -o $GOPATH/bin/go-rest-api com/mtraina/go.rest.api/*.go

Then run it with

```shell
$GOPATH/bin/go-rest-api

## Commands

### Get all the todos
curl localhost:8080/todos

### Get single todo
curl localhost:8080/todos/{todoId}

### Create a todo
curl -H "Content-Type: application/json;charset=UTF-8" -XPOST localhost:8080/todos -d '
{
    "name": "Testing Go",
    "completed":true,
    "due": "2016-02-03T00:00:00Z"
}'
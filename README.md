# Simple Rest API written in Go

Simple API for executing CRUD operations on Todo.

Extension of https://github.com/corylanou/tns-restful-json-api made by https://github.com/corylanou

Added unit test using [Ginkgo](http://onsi.github.io/ginkgo/) and [Gomega](http://onsi.github.io/gomega/)


## How to use it
Checkout the code and build it with

```shell
go build -o $GOPATH/bin/go-rest-api com/mtraina/go.rest.api/*.go
```

Then run it with

```shell
$GOPATH/bin/go-rest-api
```

## Commands

#### Get all the todos
```shell
curl localhost:8080/todos
```

#### Get single todo
```shell
curl localhost:8080/todos/{todoId}
```

#### Create a todo
```shell
curl -H "Content-Type: application/json;charset=UTF-8" -XPOST localhost:8080/todos -d '
{
    "name": "Testing Go",
    "completed":true,
    "due": "2016-02-03T00:00:00Z"
}'
```
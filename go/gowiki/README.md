# Writing Web Applications

> <https://go.dev/doc/articles/wiki/>

## Introduction

Covered in this tutorial:

1. Creating a data structure with load and save methods
2. Using the net/http package to build web applications
3. Using the html/template package to process HTML templates
4. Using the regexp package to validate user input
5. Using closures

An http.ResponseWriter value assembles the HTTP server's response; by writing to it, we send data to the HTTP client.
An http.Request is a data structure that represents the client HTTP request. r.URL.Path is the path component of the request URL

```sh
go run .

http://localhost:8080/monkeys
http://localhost:8080/view/test
http://localhost:8080/edit/test
http://localhost:8080/edit/cat
```

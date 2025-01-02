## Tutorial: Developing a RESTful API with Go and Gin

> <https://go.dev/doc/tutorial/web-service-gin>
> Introduces the basics of writing a RESTful web service API with Go and the Gin Web Framework.

Here are the endpoints you’ll create in this tutorial.

1. /albums

GET – Get a list of all albums, returned as JSON.
POST – Add a new album from request data sent as JSON.

2. /albums/:id

GET – Get an album by its ID, returning the album data as JSON.

```sh
curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"

[
    {
        "id": "1",
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
    },
    {
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
    },
    {
        "id": "3",
        "title": "Sarah Vaughan and Clifford Brown",
        "artist": "Sarah Vaughan",
        "price": 39.99
    }
]
```

```sh
curl http://localhost:8080/albums \                                                                                                                       at  github.com:younger-1/notes
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Thu, 02 Jan 2025 03:15:55 GMT
Content-Length: 116

{
    "id": "4",
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
}
```

```sh
curl http://localhost:8080/albums/2

{
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
}
```

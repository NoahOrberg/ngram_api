# ngram api

return N-gram text

## Usage
```
$ go run main.go
```

API server start at `localhost:8000`

## FOR EXAMPLE:
### request
```
$ echo '{"text":"hoge"}' | http localhost:8000/ngram
```
### response
```
HTTP/1.1 200 OK
Content-Length: 37
Content-Type: application/json; charset=utf-8
Date: Tue, 22 Aug 2017 15:03:16 GMT

{
    "response": "hoge\nogeh\ngeho\nehog"
}
```

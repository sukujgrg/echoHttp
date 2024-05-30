Echo over HTTP
===

REST API Server to **GET** what you **POST**.

## Run
```
% go version
go version go1.22.3 darwin/arm64
% go run main.go 80&
```

## POST

```
% curl 0.0.0.0 -d '{"title": "Hello", "text": "Sed congue quam sed dapibus fringilla. Phasellus aliquam porta libero semper elementum. Integer convallis metus turpis, in vehicula purus pellentesque sit amet. Curabitur ante eros, dictum quis nunc quis, rutrum ultricies ligula. Ut non risus placerat, vehicula augue nec, malesuada odio. Duis ut leo nisl. Etiam vestibulum lacus lectus. Quisque dapibus faucibus sapien, scelerisque tempus mauris volutpat vel. Donec mollis sapien laoreet lectus volutpat, ac tempus arcu cursus. Etiam eu magna faucibus, semper eros sit amet, scelerisque mauris. Duis risus ipsum, sagittis quis nisi eget, ultricies congue tortor. Suspendisse eget dui ultricies, semper arcu a, laoreet ex. Sed non orci urna. Curabitur ac aliquam massa."}'
```

## GET
```
% curl -s 0.0.0.0 | jq
{
  "Text": "Sed congue quam sed dapibus fringilla. Phasellus aliquam porta libero semper elementum. Integer convallis metus turpis, in vehicula purus pellentesque sit amet. Curabitur ante eros, dictum quis nunc quis, rutrum ultricies ligula. Ut non risus placerat, vehicula augue nec, malesuada odio. Duis ut leo nisl. Etiam vestibulum lacus lectus. Quisque dapibus faucibus sapien, scelerisque tempus mauris volutpat vel. Donec mollis sapien laoreet lectus volutpat, ac tempus arcu cursus. Etiam eu magna faucibus, semper eros sit amet, scelerisque mauris. Duis risus ipsum, sagittis quis nisi eget, ultricies congue tortor. Suspendisse eget dui ultricies, semper arcu a, laoreet ex. Sed non orci urna. Curabitur ac aliquam massa.",
  "Title": "Hello"
}
```

## GET with headers
```
% curl -s 0.0.0.0 -v
*   Trying 0.0.0.0:80...
* Connected to 0.0.0.0 (127.0.0.1) port 80
> GET / HTTP/1.1
> Host: 0.0.0.0
> User-Agent: curl/8.4.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Sat, 01 Jun 2024 21:22:12 GMT
< Content-Length: 746
<
* Connection #0 to host 0.0.0.0 left intact
{"Text":"Sed congue quam sed dapibus fringilla. Phasellus aliquam porta libero semper elementum. Integer convallis metus turpis, in vehicula purus pellentesque sit amet. Curabitur ante eros, dictum quis nunc quis, rutrum ultricies ligula. Ut non risus placerat, vehicula augue nec, malesuada odio. Duis ut leo nisl. Etiam vestibulum lacus lectus. Quisque dapibus faucibus sapien, scelerisque tempus mauris volutpat vel. Donec mollis sapien laoreet lectus volutpat, ac tempus arcu cursus. Etiam eu magna faucibus, semper eros sit amet, scelerisque mauris. Duis risus ipsum, sagittis quis nisi eget, ultricies congue tortor. Suspendisse eget dui ultricies, semper arcu a, laoreet ex. Sed non orci urna. Curabitur ac aliquam massa.","Title":"Hello"}
```

## Bad POST request data
```
[suku] ~/GitHub/echoHttp [main]  % curl -sv 0.0.0.0 -d "Bad Request"
*   Trying 0.0.0.0:80...
* Connected to 0.0.0.0 (127.0.0.1) port 80
> POST / HTTP/1.1
> Host: 0.0.0.0
> User-Agent: curl/8.4.0
> Accept: */*
> Content-Length: 11
> Content-Type: application/x-www-form-urlencoded
>
< HTTP/1.1 400 Bad Request
< Date: Sat, 01 Jun 2024 21:24:52 GMT
< Content-Length: 0
<
* Connection #0 to host 0.0.0.0 left intact
```

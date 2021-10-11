# HTTP2
The wonders of HTTP/2 in GO

## Generate HTTP/2 Certs for.
HTTP/2 enforces TLS. [There's great documentation here. ](https://pkg.go.dev/golang.org/x/net/http2).
Generate certs using the below command 

```
openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
```
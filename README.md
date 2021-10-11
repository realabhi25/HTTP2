# HTTP2
The wonders of HTTP/2 in GO

## Generate HTTP/2 Server Certs for.
HTTP/2 enforces TLS. There's great documentation  [here. ](https://pkg.go.dev/golang.org/x/net/http2)
Generate certs using the below command 

```
openssl req -x509 -out localhost.crt -keyout localhost.key \
  -newkey rsa:2048 -nodes -sha256 \
  -subj '/CN=localhost' -extensions EXT -config <( \
   printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:localhost\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")
```

## Run server
```
go run server.go
2021/10/11 07:43:17 Serving on https://0.0.0.0:8000
```

## Run Client
```
go run client.go
Got response 200: HTTP/2.0 Hello
```

## Server Push with HTTP2
A Key feature of HTTP2 is [Server Push](https://go.dev/blog/h2push) using a synthentic request. Details [here](https://pkg.go.dev/net/http#Pusher)
To consume push data from client, Open browser and connect to "https://localhost:8000". This should present the certificate challenge. Once accepted, you should see in server log
```
2021/10/11 08:50:17 Received connection: HTTP/2.0
2021/10/11 08:50:17 Handling 2nd
```
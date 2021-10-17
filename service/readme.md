# service

## command

0. init

```bash
go mod init github.com/openHacking/service
```

1. install module

> get dependencies for code in the current directory

```bash
go get .
```

2. run server

```bash
go run .
```

3. check the server api

```bash
curl http://localhost:8080/albums
```

## reference

https://golang.org/doc/tutorial/web-service-gin

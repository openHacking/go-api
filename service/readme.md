# service

## command

1. init

```bash
go mod init github.com/openHacking/service
```

2. install module

> get dependencies for code in the current directory

```bash
go get .
```

3. run server

```bash
go run .
```

4. check the server api

```bash
curl http://localhost:8080/albums
```

## reference

https://golang.org/doc/tutorial/web-service-gin

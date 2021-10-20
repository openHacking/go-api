# use local search module template

## reference

https://golang.org/doc/tutorial/create-module

## scripts

1. init

```bash
go mod init github.com/openHacking/service
```

2. redirect Go tools from its module path (where the module isn't) to the local directory (where it is).

```bash
go mod edit -replace github.com/openHacking/search=../search
```

3. install

```bash
go get .
```

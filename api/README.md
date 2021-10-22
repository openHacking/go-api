# api

go build for linux

1. build
   windows,run cmd.exe as administrator
   windows

```bash
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build
```

mac

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

2. upload file `api` to linux server

3. cd to folder which has `api` file,run

fix `Permission denied`

```bash
chmod 777 api
```

keep running

```bash
nohup ./api &
```

check api server

```bash
ps aux|grep api
```

check port

```bash
sudo netstat -tulpn | grep LISTEN
# or
sudo netstat -tnlp | grep :8080
```

kill

```bash
kill -9 [pid]
```

## reference

https://blog.csdn.net/panshiqu/article/details/53788067
https://stackoverflow.com/questions/20829155/how-to-cross-compile-from-windows-to-linux

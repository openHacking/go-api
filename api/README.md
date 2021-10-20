# api

windows go build for linux

1. Run cmd.exe as administrator

```bash
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build
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
```

kill
```bash
kill -9 [pid]
```
## reference

https://stackoverflow.com/questions/20829155/how-to-cross-compile-from-windows-to-linux
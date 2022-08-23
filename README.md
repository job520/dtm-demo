# dtm-demo

### 下载地址
`https://github.com/dtm-labs/dtm/releases`

### 启动命令
```
./dtm
./dtm -c conf.yml  # 指定配置文件
```

### 管理后台
`http://localhost:36789`

### 基本用法
[点击下载demo](https://github.com/job520/dtm-demo)

1. `http`协议：  
        1. 运行`dtm`服务：  
```
cd dtm-demo/bin
./dtm.exe
```
        2. 运行`http`服务：  
```
cd dtm-demo/http/server
go run main.go
```
        3. 运行客户端：  
```
cd dtm-demo/http/client
go run main.go
```

2. `grpc`协议：  
        1. 运行`dtm`服务：  
```
cd dtm-demo/bin
./dtm.exe
```
        2. 运行`grpc`服务：  
```
cd dtm-demo/grpc/server
go run main.go
```
        3. 运行客户端：  
```
cd dtm-demo/grpc/client
go run main.go
```

3. `go-zero-rpc`协议：  
        1. 运行`dtm`服务，指定配置文件：  
```
cd dtm-demo/bin
./dtm.exe -c conf.yml
```
        2. 运行`go-zero-rpc`服务：  
```
cd dtm-demo/go-zero/rpc
go run test.go -f etc/test.yaml
```
        3. 运行客户端：  
```
cd dtm-demo/go-zero/client
go run main.go
```

# Live

## 本机测试
```bash
go build -o main ./cmd/main.go
./main
```
浏览器打开：http://localhost:9000/swagger/index.html


## 生成`swagger`文档

```bash
# 安装swagger工具 请具体参考：https://github.com/swaggo/swag
go install github.com/swaggo/swag/cmd/swag@latest

# 动态生成文档
swag init -g pkg/handler/handler.go
```


## 函数使用
请参考：https://docs.serverless-devs.com/serverless-devs/quick_start

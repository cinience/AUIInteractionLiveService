# Live
* 本项目代码仅作范例演示使用。

## 构建执行文件
```bash
go build -o main ./cmd/main.go
```

## 本地运行
```bash
export ADMIN_PASSWORD=your-secret
./main
```
所有API及对应文档:
浏览器打开：http://localhost:7001/swagger/index.html
用户名为 admin，密码为您自行设置

* 注意：目前数据库表为自动创建，线上请关注数据库变更，及相关索引请请自行构建。


## 生成`swagger`文档

```bash
# 安装swagger工具 请具体参考：https://github.com/swaggo/swag
go install github.com/swaggo/swag/cmd/swag@latest

# 动态生成文档
swag init -g pkg/handler/*.go
```

## 函数使用
请参考：https://docs.serverless-devs.com/serverless-devs/quick_start

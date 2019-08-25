# wework

企业微信 发送信息/ 适配 open-falcon / wework

# 编译

```bash
go build -o wework cmd/main.go
```

# 或下载编译的包

```bash
wget -c https://github.com/llych/wework/releases/download/v1.0/wework_linux_amd64.tar.gz
tar zxvf wework_linux_amd64.tar.gz
./wework
```

# 运行

## 配置文件

```bash
config/app.yaml

wework:
  CorpID: 企业id
  Secret: 应用的AgentId
  AgentId: 应用的AgentId

./wework
2019-08-25T14:12:41+08:00 INF init config
2019-08-25T14:12:41+08:00 INF app start :9000

```

## 测试

```bash
curl --data "tos=xxx,xx&content=测试内容" http://127.0.0.1:9000/api/v1/wework

// {"msg":"ok"}
```

## 参数说明

```
POST /api/v1/wework
也支持json 请求
{
    tos: "用户1,用户2",
    content: "消息内容"
}
```

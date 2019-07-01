# chatroom
golang学习项目：基于tcp实现聊天室

### 功能
- 注册
- 登录
- 上线、离线通知
- 群消息
- 私聊

### 模块
- client
- server
- common

### 数据库
- redis

```
go get github.com/garyburd/redigo/redis
```

### 启动
- 运行客户端：
```
go run chatroom/client/main
```
- 运行服务端：
```
go run chatroom/server/main
```

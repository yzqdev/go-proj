# GoFrame websocket 的实现

## 实现功能

- 全局广播
- 单个房间推送
- 单个用户推送
- 单个客户端推送

## 客户端发送数据格式

```json
{
  "e": "event_name",
  "d": {
    "key": "value"
  }
}
```
```json
{ "e": "join", "d": { "name": "我的名字" } }
```

```json
{ "e": "login", "d": { "user_id": "1256" ,"client": "lllll"} }
```
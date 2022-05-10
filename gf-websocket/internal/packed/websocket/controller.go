package websocket

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gookit/color"
)

// LoginController 用户登录
func LoginController(client *Client, req *WReq) {
	fmt.Println("login controller")
	fmt.Println(req.Data)
	userId := gconv.Uint64(req.Data["user_id"])
	fmt.Println(userId)
	// 用户登录
	login := &LoginUser{
		UserId: userId,
		Client: client,
	}
	color.Cyanln(login)
	clientManager.Login <- login
	color.Redln(clientManager.Login)
	fmt.Println("登录成功")
	client.SendMsg(&WResponse{
		Event: Login,
		Data:  "success",
	})
}

func IsAppController(client *Client) {
	client.isApp = true
}

// JoinController 加入
func JoinController(client *Client, req *WReq) {
	name := gconv.String(req.Data["name"])

	if !client.tags.Contains(name) {
		client.tags.Append(name)
	}
	client.SendMsg(&WResponse{
		Event:   Join,
		Data:    client.tags.Slice(),
		Success: true,
	})
}

// QuitController 退出
func QuitController(client *Client, req *WReq) {
	name := gconv.String(req.Data["name"])
	if client.tags.Contains(name) {
		client.tags.RemoveValue(name)
	}
	client.SendMsg(&WResponse{
		Event: Quit,
		Data:  client.tags.Slice(),
	})
}
func PingController(client *Client) {
	currentTime := uint64(gtime.Now().Unix())
	client.Heartbeat(currentTime)
}

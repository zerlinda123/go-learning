package process

import (
	"fmt"
	"go-learning/chatroom/common/message"
	"go-learning/chatroom/imclient/model"
)

var onlineUsers = make(map[int]*message.User, 10)
var CurUser model.CurUser

//打印在线用户列表
func outputOnlineUser() {
	fmt.Printf("------------当前在线用户列表------------\n")
	for id := range onlineUsers {
		fmt.Printf("ID:%d\n", id)
	}
	fmt.Printf("------------end------------\n")
}

//更新在线状态
func updataUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	//打印在线列表
	outputOnlineUser()
}

package logic

import "errors"

var (
	userNamAlreadyExit = errors.New("用户名已经存在")
	userCreateFailed   = errors.New("用户失败")
	userLoginFailed    = errors.New("用户登入失败")
	friendAlreadyExit  = errors.New("该联系人已经存在")
	friendAlNotExit    = errors.New("该联系人不存在")
	sendMessageField   = errors.New("消息发送失败")
)

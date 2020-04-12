package error

import "errors"

var (
	UserNotLogin       = errors.New("用户未登入")
	UserNamAlreadyExit = errors.New("用户名已经存在")
	UserCreateFailed   = errors.New("用户失败")
	UserLoginFailed    = errors.New("用户登入失败")
	FriendAlreadyExit  = errors.New("该联系人已经存在")
	FriendAlNotExit    = errors.New("该联系人不存在")
	SendMessageField   = errors.New("消息发送失败")
	MemorandumNotExit  = errors.New("该备忘录不存在")
)

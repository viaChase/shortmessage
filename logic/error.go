package logic

import "errors"

var (
	userNamAlreadyExit = errors.New("用户名已经存在")
	userCreateFailed   = errors.New("用户失败")
	userLoginFailed    = errors.New("用户登入失败")
)

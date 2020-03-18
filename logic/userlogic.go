package logic

import (
	"fmt"
	"shortmessage/common"
	"shortmessage/model"
)

type (
	RegisterRequest struct {
		UserName string `json:"user_name"`
		PassWard string `json:"pass_ward"`
	}

	RegisterResponse struct {
		UserId int64 `json:"user_id"`
	}

	LoginRequest struct {
		UserName string `json:"user_name"`
		PassWard string `json:"pass_ward"`
	}

	LoginResponse struct {
		LastReadMessageTime int64 `json:"lastReadMessageTime"`
	}
)

func (sml *ShortMessageLogic) Register(req *RegisterRequest) (*RegisterResponse, error) {
	//先判断用户名是否存在
	userNameCount, err := sml.userModel.CountByName(req.UserName)
	if err != nil {
		return nil, err
	}

	if userNameCount >= 1 {
		return nil, userNamAlreadyExit
	}

	//返回用户id
	userId, err := sml.userModel.Insert(&model.User{UserName: req.UserName, PassWord: req.PassWard})
	if err != nil {
		return nil, userCreateFailed
	}

	return &RegisterResponse{UserId: userId}, nil
}

func (sml *ShortMessageLogic) Login(req *LoginRequest) (int64, string, *LoginResponse, error) {

	//通过用户名 查找
	user, err := sml.userModel.FindByName(req.UserName)
	if err != nil {
		return 0, "", nil, userLoginFailed
	}

	if user.PassWord != req.PassWard {
		return 0, "", nil, userCreateFailed
	}

	//header  userId = 112 ; jwt = "xxxxxx"
	//这个就是 jwt 的思想

	//先判断用户名密码是否一致，如果不一致直接返回错误
	//如果一致  ‘userId-服务器的盐’  这个字符串md5 取特征值，
	//每次用户访问 都带上userId 和 生成的md5，
	//我们只需要把 用户给的 userId 加上我们自己才知道的盐 生成字符串 取 MD5 和用户 带上的md5 判断是否一致
	//就知道 用户是不是真的登入了

	jwt := common.Md5(fmt.Sprintf("%v-%v", user.ID, sml.salt))

	return user.ID, jwt, &LoginResponse{LastReadMessageTime: user.LastReadMessageTime}, nil
}

package logic

import "shortmessage/model"

type (
	ShortMessageLogic struct {
		salt      string
		userModel *model.UserModel
	}
)

func NewShortMessageLogic(userModel *model.UserModel, salt string) *ShortMessageLogic {
	return &ShortMessageLogic{
		salt:      salt,
		userModel: userModel,
	}
}

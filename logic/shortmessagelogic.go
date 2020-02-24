package logic

import "shortmessage/model"

type (
	ShortMessageLogic struct {
		userModel *model.UserModel
	}
)

func NewShortMessageLogic(userModel *model.UserModel) *ShortMessageLogic {
	return &ShortMessageLogic{
		userModel: userModel,
	}
}

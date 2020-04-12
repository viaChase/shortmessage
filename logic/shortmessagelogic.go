package logic

import "shortmessage/model"

type (
	ShortMessageLogic struct {
		salt            string
		userModel       *model.UserModel
		mailListModel   *model.MailListModel
		messageModel    *model.MessageModel
		memorandumModel *model.MemorandumModel
	}
)

func NewShortMessageLogic(
	userModel *model.UserModel,
	mailListModel *model.MailListModel,
	messageModel *model.MessageModel,
	memorandumModel *model.MemorandumModel,
	salt string) *ShortMessageLogic {
	return &ShortMessageLogic{
		salt:            salt,
		userModel:       userModel,
		mailListModel:   mailListModel,
		messageModel:    messageModel,
		memorandumModel: memorandumModel,
	}
}

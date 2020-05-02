package logic

import "shortmessage/model"

type (
	ShortMessageLogic struct {
		salt            string
		userModel       *model.UserModel
		mailListModel   *model.MailListModel
		messageModel    *model.MessageModel
		memorandumModel *model.MemorandumModel
		backupModel     *model.BackUpDataModel
	}
)

func NewShortMessageLogic(
	userModel *model.UserModel,
	mailListModel *model.MailListModel,
	messageModel *model.MessageModel,
	memorandumModel *model.MemorandumModel,
	backupDataModel *model.BackUpDataModel,
	salt string) *ShortMessageLogic {
	return &ShortMessageLogic{
		salt:            salt,
		userModel:       userModel,
		mailListModel:   mailListModel,
		messageModel:    messageModel,
		memorandumModel: memorandumModel,
		backupModel:     backupDataModel,
	}
}

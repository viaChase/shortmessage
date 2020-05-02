package model

import (
	"github.com/go-xorm/xorm"
	"time"
)

const (
	MessageAllType = 0
	MessageNewType = 1
)

type (
	Message struct {
		ID         int64     //自增id
		SendId     int64     `xorm:"INTEGER"`      //发送id
		UserId     int64     `xorm:"INTEGER"`      //接受者id
		Content    string    `xorm:"varchar(400)"` //消息内容
		IsRead     int64     `xorm:"INTEGER"`      //是否已读
		CreateTime time.Time `xorm:"created"`      //创建时间
	}

	MessageModel struct {
		x *xorm.Engine
	}
)

/*


req server :有没有关于B的消息，我上次拿到的消息是在 5.05
	sever do : 数据库查找 有没有 5.05以后关于B的消息
server resp: no

req server :有没有关于B的消息，我上次拿到的消息是在 5.05
	sever do : 数据库查找 有没有 5.05以后关于B的消息
server resp: no

 B  5.10 "xxx" A
req server :有没有关于B的消息，我上次拿到的消息是在 5.05
	sever do : 数据库查找 有没有 5.05以后关于B的消息
server resp :yes  xxxx to B

req server :有没有关于B的消息，我上次拿到的消息是在 5.10
	sever do : 数据库查找 有没有 5.10以后关于B的消息
server resp: no


*/

func NewMessageModel(x *xorm.Engine) (*MessageModel, error) {
	if err := x.Sync(&Message{}); err != nil {
		return nil, err

	}

	return &MessageModel{x: x}, nil
}

func (mm *MessageModel) Insert(data *Message) (int64, error) {
	return mm.x.Insert(data)
}

//查找所有未读消息
func (mm *MessageModel) FindNewMessage(userId int64) ([]*Message, error) {
	var (
		messages []*Message
	)

	if err := mm.x.Where(" user_id = ? and is_read = 0 ", userId).Find(&messages); err != nil {
		return nil, err
	}

	return messages, nil
}

//查找所有已读消息
func (mm *MessageModel) FindAllMessage(userId int64, searchData string) ([]*Message, error) {
	var (
		messages []*Message
	)

	//select xxx from xxx where userid = ? and searchData like %?%
	if searchData == "" {
		if err := mm.x.Where(" user_id = ? and is_read = 1 ", userId).Find(&messages); err != nil {
			return nil, err
		}
	} else {
		if err := mm.x.Where(" user_id = ?  and is_read = 1 and content like ? ", userId, "%"+searchData+"%").Find(&messages); err != nil {
			return nil, err
		}
	}

	return messages, nil
}

//删除所有已读消息
func (mm *MessageModel) DeleteReadMessageByUserId(userId int64) (int64, error) {
	return mm.x.Where(" user_id = ? and is_read = 1 ", userId).Delete(&Message{})
}

func (mm *MessageModel) SetMessageRead(messageId, userId int64) error {
	_, err := mm.x.Exec("update message set is_read = 1 where i_d = ? and user_id = ? ", messageId, userId)
	return err
}

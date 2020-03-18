package model

import (
	"github.com/go-xorm/xorm"
	"time"
)

type (
	Message struct {
		ID         int64     //自增id
		SendId     int64     `xorm:"INTEGER"`      //发送id
		UserId     int64     `xorm:"INTEGER"`      //接受者id
		Content    string    `xorm:"varchar(400)"` //消息内容
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

func (mm *MessageModel) FindNewMessage(userId, lastTime int64) ([]*Message, error) {
	var (
		messages []*Message
	)

	if err := mm.x.Where("CreateTime > ? and UserId = ? ", time.Unix(lastTime, 0), userId).
		Find(messages); err != nil {
		return nil, err
	}

	return messages, nil
}

package model

import (
	"github.com/go-xorm/xorm"
	"time"
)

type (
	MailList struct {
		ID         int64
		UserId     int64     `xorm:"INTEGER"`
		FriendId   int64     `xorm:"INTEGER"`
		FriendName string    `xorm:"varchar(400)"`
		CreateTime time.Time `xorm:"created"`
		UpdateTime time.Time `xorm:"updated"`
	}

	MailListModel struct {
		x *xorm.Engine
	}
)

func NewMailListModel(x *xorm.Engine) (*MailListModel, error) {
	if err := x.Sync(&MailList{}); err != nil {
		return nil, err

	}

	return &MailListModel{x: x}, nil
}

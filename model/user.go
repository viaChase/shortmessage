package model

import (
	"github.com/go-xorm/builder"
	"github.com/go-xorm/xorm"
	"time"
)

type (
	User struct {
		ID                  int64
		UserName            string    `xorm:"varchar(400)"`
		PassWord            string    `xorm:"varchar(400)"`
		LastReadMessageTime int64     `xorm:"INTEGER"` //上次读取消息的时间戳
		CreateTime          time.Time `xorm:"created"`
		UpdateTime          time.Time `xorm:"updated"`
	}

	UserModel struct {
		x *xorm.Engine
	}
)

func NewUserModel(x *xorm.Engine) (*UserModel, error) {
	if err := x.Sync(&User{}); err != nil {
		return nil, err
	}

	return &UserModel{x: x}, nil
}

func (um *UserModel) Insert(data *User) (int64, error) {
	//um.x.Exec("insert into user (...) (?,?)")
	return um.x.Insert(data)
}

func (um *UserModel) FindById(userId int64) (*User, error) {
	var (
		eq = builder.Eq{
			"i_d": userId,
		}

		data User
	)

	if err := um.x.Find(&data, eq); err != nil {
		return nil, err
	}

	//um.x.Query("select * from user where username = ? and password = ? ", userName,password)
	return &data, nil
}

func (um *UserModel) FindByName(userName string) (*User, error) {
	var (
		eq = builder.Eq{
			"UserName": userName,
		}

		data User
	)

	if err := um.x.Find(&data, eq); err != nil {
		return nil, err
	}

	//um.x.Query("select * from user where username = ? and password = ? ", userName,password)
	return &data, nil
}

func (um *UserModel) CountByName(userName string) (int64, error) {
	var (
		eq = builder.Eq{
			"UserName": userName,
		}
	)

	//select count(1) from table where username = username
	return um.x.Count(eq)
}

func (um *UserModel) Update(data *User) error {
	var (
		eq = builder.Eq{
			"ID": data.ID,
		}
	)

	if _, err := um.x.Update(data, eq); err != nil {
		return err
	}

	return nil
}

func (um *UserModel) UpdateLastReadTime(userId, lastReadTime int64) error {
	_, err := um.x.Exec("update user set last_read_message_time = ? where i_d = ? ", lastReadTime, userId)
	return err
}

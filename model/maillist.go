package model

import (
	"github.com/go-xorm/builder"
	"github.com/go-xorm/xorm"
	"time"
)

type (
	MailList struct {
		ID         int64     //自增id
		UserId     int64     `xorm:"INTEGER"`      //用户id
		FriendId   int64     `xorm:"INTEGER"`      //好友手机号
		FriendName string    `xorm:"varchar(400)"` //好友备注
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

//  select * from mailList where FriendId = xxx and userId = xxx
func (mlm *MailListModel) FindByUserIdAndFriendId(friendId, userId int64) (*MailList, error) {

	/*
		go 定义一个类型的话 都会给它一个默认值

		如果定义了 指针 map slice channel  默认值 就是 nil

		如果是定义具体类型的话，go 会给它初始化

	*/

	var (
		ml MailList
		eq = builder.Eq{
			"FriendId": friendId,
			"UserId":   userId,
		}
	)

	if err := mlm.x.Find(&ml, eq); err != nil {
		return nil, err
	}

	return &ml, nil
}

func (mlm *MailListModel) CountByUserIdAndFriendId(friendId, userId int64) (int64, error) {
	var (
		eq = builder.Eq{
			"FriendId": friendId,
			"UserId":   userId,
		}
	)

	return mlm.x.Count(eq)
}

func (mlm *MailListModel) Insert(data *MailList) (int64, error) {
	return mlm.x.Insert(data)
}

func (mlm *MailListModel) FindByUserId(nowPage, pageSize, userId int64, searchData string) ([]*MailList, error) {
	var (
		data []*MailList
		eq   = builder.Eq{
			"UserId": userId,
		}
		like = builder.Like{
			"FriendName", "%" + searchData + "%",
		}
	)

	// select * from table where userId = 111 limit 0, 10 反射
	if searchData == "" {
		if err := mlm.x.Limit(int(pageSize), int(pageSize)*int(nowPage-1)).Find(&data, eq); err != nil {
			return nil, err
		}
	} else {
		if err := mlm.x.Limit(int(pageSize), int(pageSize)*int(nowPage-1)).Find(&data, eq, like); err != nil {
			return nil, err
		}
	}

	return data, nil
}

//查找这个人有多少条记录
func (mlm *MailListModel) CountByUserId(userId int64, searchData string) (int64, error) {
	var (
		eq = builder.Eq{
			"UserId": userId,
		}

		like = builder.Like{
			"FriendName", "%" + searchData + "%",
		}
	)

	if searchData == "" {
		return mlm.x.Count(eq)
	} else {
		return mlm.x.Count(eq, like) //select * from table where userId = 111 and FriendName like "%雨%"
	}
}

func (mlm *MailListModel) DeleteByUserIdAndFriendId(userId, friendId int64) (int64, error) {
	var (
		eq = builder.Eq{
			"FriendId": friendId,
			"UserId":   userId,
		}
	)

	return mlm.x.Delete(eq)
}

func (mlm *MailListModel) UpdateByUserIdAndFriendId(data *MailList) (int64, error) {
	// update table set fiendName = "xxx" where Id = 1
	return mlm.x.Id(data.ID).Update(data)
}

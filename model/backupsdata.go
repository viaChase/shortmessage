package model

import (
	"github.com/go-xorm/xorm"
	"time"
)

const (
	MailListType   = 1
	MemorandumType = 2
	MessageType    = 3
)

type (
	BackUpData struct {
		ID         int64     //自增id
		UserId     int64     `xorm:"INTEGER"` //用户id
		Data       string    `xorm:"text"`    //备份数据
		DataType   int64     `xorm:"INTEGER"` //数据类型 1 通讯录 2备忘录 3短信
		CreateTime time.Time `xorm:"created"`
	}

	BackUpDataModel struct {
		x *xorm.Engine
	}
)

func NewBackUpDataModel(x *xorm.Engine) (*BackUpDataModel, error) {
	if err := x.Sync(&BackUpData{}); err != nil {
		return nil, err

	}

	return &BackUpDataModel{x: x}, nil
}

func (bdm *BackUpDataModel) Insert(data *BackUpData) (int64, error) {
	return bdm.x.Insert(data)
}

func (bdm *BackUpDataModel) Delete(id, userId int64) (int64, error) {
	return bdm.x.Where("i_d = ? and user_id = ?", id, userId).Delete(&BackUpData{})
}

func (bdm *BackUpDataModel) Find(userId int64, pageNum, pageSize int) ([]*BackUpData, error) {
	var data []*BackUpData
	if err := bdm.x.Where("user_id = ？ ", userId).Limit(pageSize, (pageNum-1)*pageSize).Find(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (bdm *BackUpDataModel) Count(userId int64) (int64, error) {
	return bdm.x.Where("user_id = ？ ", userId).Count(&BackUpData{})
}

func (bdm *BackUpDataModel) FindById(userId, id int64) (*BackUpData, error) {
	var (
		data BackUpData
	)

	if find, err := bdm.x.Where("i_d = ? and user_id = ? ", id, userId).Get(&data); err != nil || find == false {
		return nil, ErrNotFind
	}

	return &data, nil
}

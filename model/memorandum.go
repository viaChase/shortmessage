package model

import (
	"github.com/go-xorm/xorm"
	"time"
)

type (
	Memorandum struct {
		ID         int64     //自增id
		UserId     int64     `xorm:"INTEGER"`      //用户id
		Title      string    `xorm:"varchar(100)"` //备忘录标题
		Content    string    `xorm:"varchar(800)"` //备忘录内容
		CreateTime time.Time `xorm:"created"`      //创建时间
	}

	MemorandumModel struct {
		x *xorm.Engine
	}
)

func NewMemorandumModel(x *xorm.Engine) (*MemorandumModel, error) {
	if err := x.Sync(&Memorandum{}); err != nil {
		return nil, err

	}

	return &MemorandumModel{x: x}, nil
}

func (mdm *MemorandumModel) Insert(data *Memorandum) (int64, error) {
	return mdm.x.Insert(data)
}

func (mdm *MemorandumModel) Update(data *Memorandum) (int64, error) {
	return mdm.x.ID(data.ID).Update(data)

}

func (mdm *MemorandumModel) Delete(id, userId int64) error {
	var query = "delete from memorandum where i_d = ?  and user_id = ? "
	_, err := mdm.x.Exec(query, id, userId)
	return err
}

func (mdm *MemorandumModel) Find(userId int64, pageNum, pageSize int, searchData string) ([]*Memorandum, error) {
	var data []*Memorandum
	if searchData == "" {
		if err := mdm.x.Where("user_id = ? ", userId).Limit(pageSize, (pageNum-1)*pageSize).Find(&data); err != nil {
			return nil, err
		}

	} else {
		if err := mdm.x.Where("user_id = ? and ( content like ? or title like ? ) ", userId, "%"+searchData+"%", "%"+searchData+"%").Limit(pageSize, (pageNum-1)*pageSize).Find(&data); err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (mdm *MemorandumModel) FindOne(id, userId int64) (*Memorandum, error) {

	var (
		data Memorandum
	)

	if find, err := mdm.x.Where(" i_d = ? and user_id = ? ", id, userId).Get(&data); err != nil || find == false {
		return nil, ErrNotFind
	}

	return &data, nil

}

func (mdm *MemorandumModel) Count(userId int64, searchData string) (int64, error) {

	if searchData == "" {
		return mdm.x.Where("user_id = ? ", userId).Count(&Memorandum{})

	} else {
		return mdm.x.Where("user_id = ? and ( content like ? or title like ? ) ", userId, "%"+searchData+"%", "%"+searchData+"%").Count(&Memorandum{})
	}
}

func (mdm *MemorandumModel) DeleteByUserId(userId int64) (int64, error) {
	return mdm.x.Where(" user_id = ?", userId).Delete(&Memorandum{})
}

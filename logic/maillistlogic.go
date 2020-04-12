package logic

import (
	"fmt"
	"log"
	er "shortmessage/error"
	"shortmessage/model"
)

type (
	AddContactsRequest struct {
		ContactsId   int64  `json:"contactsId"`   //用户id 类似于手机号
		ContactsName string `json:"contactsName"` //备注
	}

	ContactsListRequest struct {
		Search   string `json:"search"`   //需要搜索的关键词
		PageSize int64  `json:"pageSize"` //一页多少数据
		NowPage  int64  `json:"nowPage"`  //当前第几页
	}

	ContactsListResponse struct {
		Count        int64               `json:"count"`        //总条数
		ContactsList []*ContactsListItem `json:"contactsList"` //数据
	}

	ContactsListItem struct {
		FriendName string `json:"friendName"` //好友备注
		FriendId   int64  `json:"friendId"`   //好友id 其实就是手机号
	}

	DeleteContactRequest struct {
		FriendId int64 `json:"friendId"` //好友id 其实就是手机号
	}

	UpdateContactsRequest struct {
		ContactsId   int64  `json:"contactsId"`   //用户id 类似于手机号
		ContactsName string `json:"contactsName"` //备注
	}
)

// 增加联系人
func (sml *ShortMessageLogic) AddContacts(req *AddContactsRequest, userId int64) error {
	fmt.Println(req.ContactsName)
	//判断用户是否已经添加了该联系人
	count, _ := sml.mailListModel.CountByUserIdAndFriendId(req.ContactsId, userId)
	if count > 0 {
		return er.FriendAlreadyExit
	}

	//插入数据
	_, _ = sml.mailListModel.Insert(&model.MailList{
		UserId:     userId,
		FriendId:   req.ContactsId,
		FriendName: req.ContactsName,
	})

	return nil
}

//获得通讯录列表
func (sml *ShortMessageLogic) ContactsList(req *ContactsListRequest, userId int64) (*ContactsListResponse, error) {

	data, err := sml.mailListModel.FindByUserId(req.NowPage, req.PageSize, userId, req.Search)
	if err != nil {
		return nil, err
	}

	count, err := sml.mailListModel.CountByUserId(userId, req.Search)
	if err != nil {
		return nil, err
	}

	var contactsList = make([]*ContactsListItem, 0, len(data))

	for i := 0; i < len(data); i++ {
		contactsList = append(contactsList, &ContactsListItem{
			FriendName: data[i].FriendName,
			FriendId:   data[i].FriendId,
		})
	}

	return &ContactsListResponse{
		ContactsList: contactsList,
		Count:        count,
	}, nil
}

//删除联系人
func (sml *ShortMessageLogic) DelContact(req *DeleteContactRequest, userId int64) error {
	if _, err := sml.mailListModel.DeleteByUserIdAndFriendId(userId, req.FriendId); err != nil {
		log.Println(err)
	}
	return nil
}

//修改备注
func (sml *ShortMessageLogic) UpdateContact(req *UpdateContactsRequest, userId int64) error {
	data, err := sml.mailListModel.FindByUserIdAndFriendId(req.ContactsId, userId)
	if err != nil {
		return er.FriendAlNotExit
	}

	data.FriendName = req.ContactsName

	if _, err = sml.mailListModel.UpdateByUserIdAndFriendId(data); err != nil {
		log.Println(err)
	}

	return nil
}

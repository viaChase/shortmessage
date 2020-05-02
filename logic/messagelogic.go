package logic

import (
	"fmt"
	er "shortmessage/error"
	"shortmessage/model"
	"time"
)

type (
	SendMessageRequest struct {
		TargetUserId int64  `json:"targetUserId"`
		Content      string `json:"content"`
	}

	MessageListItem struct {
		SendId         int64  `json:"sendId"`
		SendFriendName string `json:"sendFriendName"`
		Content        string `json:"content"`
		SendTime       int64  `json:"sendTime"`
	}

	MessagePeopleViewRequest struct {
		DataType   int64  `json:"data_type"`
		SearchData string `json:"searchData"`
	}

	MessagePeopleViewResponse struct {
		UserMessageList []*UserMessageListItem `json:"userMessageList"`
	}

	UserMessageListItem struct {
		UserId      int64          `json:"userId"`
		UserName    string         `json:"userName"`
		MessageList []*MessageItem `json:"messageList"`
	}

	MessageItem struct {
		MessageId  int64  `json:"messageId"`
		Content    string `json:"content"`
		CreateTime int64  `json:"createTime"`
		IsNew      bool   `json:"isNew"`
	}

	MessageReadRequest struct {
		MessageId []int64
	}
)

func (sml *ShortMessageLogic) SendMessage(req *SendMessageRequest, userId int64) error {
	//先判断用户是否存在
	targetUser, err := sml.userModel.FindById(req.TargetUserId)
	if err != nil || targetUser == nil {
		return er.FriendAlNotExit
	}

	//如果该用户存在就往message数据库里插入一条数据
	_, err = sml.messageModel.Insert(&model.Message{
		SendId:     userId,
		UserId:     targetUser.ID,
		Content:    req.Content,
		CreateTime: time.Now(),
	})

	if err != nil {
		return er.SendMessageField
	}

	return nil
}

// 获得消息
func (sml *ShortMessageLogic) MessagePeopleView(req *MessagePeopleViewRequest, userId int64) (*MessagePeopleViewResponse, error) {

	var (
		messageList      []*model.Message
		unReadMessageIds []int64
		err              error
	)

	switch req.DataType {
	case model.MessageAllType:
		messageList, err = sml.messageModel.FindAllMessage(userId, req.SearchData)
	case model.MessageNewType:
		messageList, err = sml.messageModel.FindNewMessage(userId)
	}

	if err != nil {
		return nil, err
	}

	if len(messageList) == 0 {
		return &MessagePeopleViewResponse{
			UserMessageList: nil,
		}, nil
	}

	var (
		friendMessageMap = make(map[int64][]*MessageItem)
		resp             = &MessagePeopleViewResponse{}
	)

	for _, message := range messageList {

		friendMessageMap[message.SendId] = append(friendMessageMap[message.SendId], &MessageItem{
			MessageId:  message.ID,
			Content:    message.Content,
			CreateTime: message.CreateTime.Unix(),
			IsNew:      message.IsRead == 0,
		})

		if message.IsRead == 0 {
			unReadMessageIds = append(unReadMessageIds, message.ID)
		}
	}

	for friendId, messages := range friendMessageMap {
		friendName := fmt.Sprintf("%v", friendId)
		if mailList, err := sml.mailListModel.FindByUserIdAndFriendId(friendId, userId); err == nil {
			friendName = mailList.FriendName
		}

		resp.UserMessageList = append(resp.UserMessageList, &UserMessageListItem{
			UserId:      friendId,
			UserName:    friendName,
			MessageList: messages,
		})
	}

	//把未读消息设置为已读消息
	for _, unReadMessageId := range unReadMessageIds {
		_ = sml.messageModel.SetMessageRead(unReadMessageId, userId)
	}

	return resp, nil
}

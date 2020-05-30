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
		IsSelf     bool   `json:"isSelf"`
		IsNew      bool   `json:"isNew"`
	}

	MessageReadRequest struct {
		MessageId []int64 `json:"message_id"`
	}

	MessageDeleteRequest struct {
		MessageId []int64 `json:"message_id"`
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
		SelfId:     0,
		CreateTime: time.Now(),
	})

	if err != nil {
		return er.SendMessageField
	}

	//插入发送记录
	_, err = sml.messageModel.Insert(&model.Message{
		SendId:     userId,
		UserId:     targetUser.ID,
		Content:    req.Content,
		SelfId:     userId,
		IsRead:     0,
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
		messageList []*model.Message
		err         error
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

		//selfId 不等于0 说明，这条消息肯定是 消息记录，如果这条消息记录不是自己的 就直接过滤
		if message.SelfId != 0 && message.SelfId != userId {
			continue
		}

		friendMessageMap[message.SendId] = append(friendMessageMap[message.SendId], &MessageItem{
			MessageId:  message.ID,
			Content:    message.Content,
			CreateTime: message.CreateTime.Unix(),
			IsNew:      message.IsRead == 0,
			IsSelf:     message.SelfId == userId, //是否是发送消息
		})

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

	return resp, nil
}

//把某些消息置为已读
func (sml *ShortMessageLogic) MessageRead(req *MessageReadRequest, userId int64) {
	for _, unReadMessageId := range req.MessageId {
		_ = sml.messageModel.SetMessageRead(unReadMessageId, userId)
	}
}

//删除指定消息
func (sml *ShortMessageLogic) MessageDelete(req *MessageDeleteRequest, userId int64) {
	for _, messageId := range req.MessageId {
		_, _ = sml.messageModel.DeleteMessageById(userId, messageId)
	}
}

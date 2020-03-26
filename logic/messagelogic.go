package logic

import (
	"fmt"
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
		LastTime   int64  `json:"lastTime"`
		SearchData string `json:"searchData"`
	}

	MessagePeopleViewResponse struct {
		LastTime        int64                  `json:"lastTime"`
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
)

func (sml *ShortMessageLogic) SendMessage(req *SendMessageRequest, userId int64) error {
	//先判断用户是否存在
	targetUser, err := sml.userModel.FindById(req.TargetUserId)
	if err != nil || targetUser == nil {
		return friendAlNotExit
	}

	//如果该用户存在就往message数据库里插入一条数据
	_, err = sml.messageModel.Insert(&model.Message{
		SendId:     userId,
		UserId:     targetUser.ID,
		Content:    req.Content,
		CreateTime: time.Now(),
	})

	if err != nil {
		return sendMessageField
	}

	return nil
}

// 获得消息
func (sml *ShortMessageLogic) MessagePeopleView(req *MessagePeopleViewRequest, userId int64) (*MessagePeopleViewResponse, error) {

	var (
		lastTime = req.LastTime
		hasNew   = false
	)

	if req.LastTime == 0 {
		user, err := sml.userModel.FindById(userId)
		if err != nil {
			return nil, err
		}
		lastTime = user.LastReadMessageTime
	}

	messageList, err := sml.messageModel.FindAllMessage(userId, req.LastTime, req.SearchData)
	if err != nil {
		return nil, err
	}

	if len(messageList) == 0 {
		return &MessagePeopleViewResponse{
			LastTime:        lastTime,
			UserMessageList: nil,
		}, nil
	}

	var (
		friendMessageMap = make(map[int64][]*MessageItem)
		resp             = &MessagePeopleViewResponse{}
	)

	for _, message := range messageList {

		if message.CreateTime.Unix() > lastTime {
			lastTime = message.CreateTime.Unix()
			hasNew = true
		}

		friendMessageMap[message.SendId] = append(friendMessageMap[message.SendId], &MessageItem{
			MessageId:  message.ID,
			Content:    message.Content,
			CreateTime: message.CreateTime.Unix(),
			IsNew:      message.CreateTime.Unix() > req.LastTime,
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

	resp.LastTime = lastTime
	if hasNew {
		_ = sml.userModel.UpdateLastReadTime(userId, lastTime)
	}

	return resp, nil
}

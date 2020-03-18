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

	NewMessageRequest struct {
		LastTimeSnap int64 `json:"lastTimeSnap"`
	}

	NewMessageResponse struct {
		LastTimeSnap int64              `json:"lastTimeSnap"`
		MessageList  []*MessageListItem `json:"messageList"`
	}

	MessageListHistoryRequest struct {
		SearchData string `json:"searchData"` //短信内容关键词 如果为"" 就是不过滤
		SendId     int64  `json:"sendId"`     //谁发送的 如果未0就是所有人
		PageSize   int64  `json:"pageSize"`
		PageNum    int64  `json:"pageNum"`
	}

	MessageListHistoryResponse struct {
		Count       int64              `json:"count"` //一共多少条数据
		MessageList []*MessageListItem `json:"messageList"`
	}

	MessageListItem struct {
		SendId         int64  `json:"sendId"`
		SendFriendName string `json:"sendFriendName"`
		Content        string `json:"content"`
		SendTime       int64  `json:"sendTime"`
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

func (sml *ShortMessageLogic) NewMessage(req *NewMessageRequest, userId int64) (*NewMessageResponse, error) {

	messagesList, err := sml.messageModel.FindNewMessage(userId, req.LastTimeSnap)
	if err != nil {
		return nil, err
	}

	if len(messagesList) == 0 {
		return &NewMessageResponse{
			LastTimeSnap: req.LastTimeSnap,
			MessageList:  nil,
		}, nil
	}

	var (
		messageDatas  []*MessageListItem
		lastTimeStnp  int64
		friendNameMap = make(map[int64]string)
	)

	for i := 0; i < len(messagesList); i++ {
		var sendFriendName = ""

		sendFriendName, exit := friendNameMap[messagesList[i].SendId]
		if exit == false {
			if mailList, err := sml.mailListModel.FindByUserIdAndFriendId(messagesList[i].SendId, userId); err == nil {
				sendFriendName = mailList.FriendName
			} else {
				sendFriendName = fmt.Sprintf("%v", messagesList[i].SendId)
			}
			friendNameMap[messagesList[i].SendId] = sendFriendName
		}

		if lastTimeStnp < messagesList[i].CreateTime.Unix() {
			lastTimeStnp = messagesList[i].CreateTime.Unix()
		}

		messageDatas = append(messageDatas, &MessageListItem{
			SendId:         messagesList[i].SendId,
			SendFriendName: sendFriendName,
			Content:        messagesList[i].Content,
			SendTime:       messagesList[i].CreateTime.Unix(),
		})
	}

	if len(messageDatas) != 0 {
		_ = sml.userModel.UpdateLastReadTime(userId, lastTimeStnp)
	}

	return &NewMessageResponse{
		LastTimeSnap: lastTimeStnp,
		MessageList:  messageDatas,
	}, nil
}

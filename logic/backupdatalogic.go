package logic

import (
	"encoding/json"
	"fmt"
	"shortmessage/model"
)

type (
	BackupDataAddRequest struct {
		DataType int64 `json:"data_type"`
	}

	BackupDataDeleteRequest struct {
		Id int64 `json:"id"`
	}

	BackupDataListRequest struct {
		PageSize int64 `json:"page_size"`
		PageNum  int64 `json:"page_num"`
	}

	BackupDataListResponse struct {
		Count int64                 `json:"count"`
		Data  []*BackupDataListItem `json:"data"`
	}

	BackupDataListItem struct {
		BackId     int64 `json:"back_id"`
		CreateTime int64 `json:"create_time"`
		DataType   int64 `json:"data_type"`
	}

	BackupRequest struct {
		Id int64 `json:"id"`
	}
)

//数据备份接口
func (sml *ShortMessageLogic) BackupAdd(req *BackupDataAddRequest, userId int64) error {

	jsonData := ""
	switch req.DataType {
	case model.MailListType:
		dataList, err := sml.mailListModel.FindByUserId(1, 999999, userId, "")
		if err != nil {
			return err
		}

		dataBytes, err := json.Marshal(dataList)
		if err != nil {
			return err
		}

		jsonData = string(dataBytes)

	case model.MemorandumType:
		dataList, err := sml.memorandumModel.Find(userId, 1, 999999, "")
		if err != nil {
			return err
		}

		dataBytes, err := json.Marshal(dataList)
		if err != nil {
			return err
		}

		jsonData = string(dataBytes)

	case model.MessageType:
		dataList, err := sml.messageModel.FindAllMessage(userId, "")
		if err != nil {
			return err
		}

		dataBytes, err := json.Marshal(dataList)
		if err != nil {
			return err
		}

		jsonData = string(dataBytes)
	}

	fmt.Println(jsonData)

	_, err := sml.backupModel.Insert(&model.BackUpData{DataType: req.DataType, UserId: userId, Data: jsonData})
	return err
}

func (sml *ShortMessageLogic) BackupDelete(req *BackupDataDeleteRequest, userId int64) error {
	_, err := sml.backupModel.Delete(req.Id, userId)
	return err
}

func (sml *ShortMessageLogic) BackupList(req *BackupDataListRequest, userId int64) (*BackupDataListResponse, error) {
	count, err := sml.backupModel.Count(userId)
	if err != nil {
		return nil, err
	}

	dataList, err := sml.backupModel.Find(userId, int(req.PageNum), int(req.PageSize))

	resp := &BackupDataListResponse{
		Count: count,
		Data:  nil,
	}

	for _, data := range dataList {
		resp.Data = append(resp.Data, &BackupDataListItem{
			BackId:     data.ID,
			CreateTime: data.CreateTime.Unix(),
			DataType:   data.DataType,
		})
	}

	return resp, nil
}

func (sml *ShortMessageLogic) BackUp(req *BackupRequest, userId int64) error {
	data, err := sml.backupModel.FindById(userId, req.Id)
	if err != nil {
		return err
	}

	switch data.DataType {
	case model.MailListType:
		_, _ = sml.mailListModel.DeleteByUserId(userId)
		var dataList []*model.MailList
		_ = json.Unmarshal([]byte(data.Data), &dataList)
		for _, data := range dataList {
			_, _ = sml.mailListModel.Insert(data)
		}
	case model.MemorandumType:
		_, _ = sml.memorandumModel.DeleteByUserId(userId)
		var dataList []*model.Memorandum
		_ = json.Unmarshal([]byte(data.Data), &dataList)
		for _, data := range dataList {
			_, _ = sml.memorandumModel.Insert(data)
		}

	case model.MessageType:
		_, _ = sml.messageModel.DeleteReadMessageByUserId(userId)
		var dataList []*model.Message
		_ = json.Unmarshal([]byte(data.Data), &dataList)
		for _, data := range dataList {
			_, _ = sml.messageModel.Insert(data)
		}
	}

	return nil
}

package logic

import (
	"shortmessage/model"
	"time"
)

type (
	TestRequest struct {
		Name     string `form:"name"`
		PassWord string `form:"pass_word"`
	}

	TestResponse struct {
		Now int64 `json:"now"`
	}
)

func (sml *ShortMessageLogic) Test(req *TestRequest) (*TestResponse, error) {

	_, err := sml.userModel.Insert(&model.User{
		UserName:   req.Name,
		PassWord:   req.PassWord,
		Token:      "112",
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	})

	if err != nil {
		return nil, err
	}

	return &TestResponse{Now: time.Now().Unix()}, nil
}

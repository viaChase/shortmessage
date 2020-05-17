package logic

import (
	"fmt"
	"log"
	er "shortmessage/error"
	"shortmessage/model"
)

type (
	AddContentRequest struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	DeleteContentRequest struct {
		Id int64 `json:"id"`
	}

	UpdateContentRequest struct {
		ContentId int64  `json:"contentId"`
		Title     string `json:"title"`
		Content   string `json:"content"`
	}

	ContentListRequest struct {
		Search   string `json:"search"`   //需要搜索的关键词
		PageSize int64  `json:"pageSize"` //一页多少数据
		NowPage  int64  `json:"nowPage"`  //当前第几页
	}

	ContentListResponse struct {
		Count       int64              `json:"count"`       //总条数
		ContentList []*ContentListItem `json:"contentlist"` //内容
	}

	ContentListItem struct {
		ContentId int64  `json:"contentId"`
		Title     string `json:"title"`
		Content   string `json:"content"`
	}
)

//增加一条备忘录
func (sml *ShortMessageLogic) AddContent(req *AddContentRequest, userId int64) error {
	fmt.Println(req.Title, req.Content)
	_, _ = sml.memorandumModel.Insert(&model.Memorandum{
		UserId:  userId,
		Title:   req.Title,
		Content: req.Content,
	})
	return nil
}

//删除一条备忘录
func (sml *ShortMessageLogic) DelContent(req *DeleteContentRequest, userId int64) error {
	if err := sml.memorandumModel.Delete(req.Id, userId); err != nil {
		log.Println(err)
	}
	return nil
}

//修改备忘录
func (sml *ShortMessageLogic) Update(req *UpdateContentRequest, userId int64) error {

	//findone 函数
	data, err := sml.memorandumModel.FindOne(req.ContentId, userId)
	if err != nil {
		return er.MemorandumNotExit
	}

	data.Title = req.Title
	data.Content = req.Content

	if _, err = sml.memorandumModel.Update(data); err != nil {
		log.Println(err)
	}

	return nil
}

//获取备忘录列表
func (sml *ShortMessageLogic) ContentList(req *ContentListRequest, userId int64) (*ContentListResponse, error) {

	data, err := sml.memorandumModel.Find(userId, int(req.NowPage), int(req.PageSize), req.Search)
	if err != nil {
		return nil, err
	}

	count, err := sml.memorandumModel.Count(userId, req.Search)
	if err != nil {
		return nil, err
	}

	var contentList = make([]*ContentListItem, 0, len(data))

	for i := 0; i < len(data); i++ {

		contentList = append(contentList, &ContentListItem{
			ContentId: data[i].ID,
			Title:     data[i].Title,
			Content:   data[i].Content,
		})
	}

	return &ContentListResponse{
		ContentList: contentList,
		Count:       count,
	}, nil
}

package logic

import (
	"log"
	"shortmessage/model"
)

//todo 备忘录

type (
	AddContentRequest struct {
		UserId  int64  `json:"userId"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	DeleteContentRequest struct {
		UserId int64  `json:"userId"`
		Title  string `json:"title"`
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
		Title   string `json:"title"`
		Content string `json:"content"`
	}
)

//增加一条备忘录
func (sml *ShortMessageLogic) AddContent(req *AddContentRequest, userId int64) error {
	_, _ = sml.memorandumModel.Insert(&model.Memorandum{
		UserId:  req.UserId,
		Title:   req.Title,
		Content: req.Content,
	})
	return nil
}

//删除一条备忘录
func (sml *ShortMessageLogic) DelContent(req *DeleteContentRequest, userId int64) error {
	if err := sml.memorandumModel.Delete(userId); err != nil {
		log.Println(err)
	}
	return nil
}

//修改备忘录
func (sml *ShortMessageLogic) Update(req *UpdateContentRequest, userId int64) error {

	//findone 函数

	if _, err = sml.memorandumModel.Update(data); err != nil {
		log.Println(err)
	}

	return nil
}

//获取备忘录列表
func (sml *ShortMessageLogic) ContentList(req *ContentListRequest, userId int64) (*ContentListResponse, error) {

	data, err := sml.memorandumModel.Find(userId, req.NowPage, req.PageSize, req.Search)
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
			Title:   data[i].Title,
			Content: data[i].Content,
		})
	}

	return &ContentListResponse{
		ContentList: contentList,
		Count:       count,
	}, nil
}

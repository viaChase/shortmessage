package logic

type (
	LoginRequest struct {
		Name string `json:"name"`
	}

	LoginResponse struct {
	}
)

func (sml *ShortMessageLogic) Login(req *LoginRequest) (*LoginResponse, error) {

	//查询数据库
	//判断用户密码是否正确

	return nil, nil
}

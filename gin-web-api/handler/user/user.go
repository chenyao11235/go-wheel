package user

import "wheel/go-web-gin/model"

// 研发经验： 如果小系体中有json参数要传递，建议针对每个api定义独立的go struct来接收，并将这些结构体放在同一个go文件中，方便后续的修改维护

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"'offset'"`
	Limit    int    `json:"speed_limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
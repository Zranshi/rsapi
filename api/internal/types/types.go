// Code generated by goctl. DO NOT EDIT.
package types

type DefaultResp struct {
	Ok  bool   `json:"ok"`
	Msg string `json:"msg"`
}

type LoginReq struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

type LoginResp struct {
	DefaultResp
}

type PageReq struct {
	Cursor   int64 `path:"cursor"`
	PageSize int64 `path:"pagesize"`
	IsEnd    int64 `query:"is_end"`
}

type RegisterReq struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

type RegisterResp struct {
	DefaultResp
}

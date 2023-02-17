package types

// UserResp is the response of user register/login api
type UserResp struct {
	Response
	Token  string `json:"token"`   // 用户鉴权token
	UserID int64  `json:"user_id"` // 用户id
}

// UserReq is the request of user register/login api
type UserReq struct {
	Username string `form:"username" json:"username" query:"username" validate:"required,alphanum,min=1,max=16"`   // 用户名
	Password string `form:"password" json:"password" query:"password" validate:"required,printascii,min=8,max=16"` // 密码
}

// UserInfoReq is the request of user info api
type UserInfoReq struct {
	UserID int64  `query:"user_id" binding:"required"`
	Token  string `query:"token" binding:"required"`
}

// UserInfoResp is the response of user info api
type UserInfoResp struct {
	Response
	User *User `json:"user,omitempty"` // 用户信息
}

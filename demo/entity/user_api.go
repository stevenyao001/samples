package entity

//创建用户
type CreateUserReq struct {
	Uid          int    `form:"uid" json:"uid" uri:"uid" xml:"uid" binding:"required"`
	Name         string `form:"name" json:"name" uri:"name" xml:"name" binding:"required"`
	Birthday     int    `form:"birthday" json:"birthday" uri:"birthday" xml:"birthday" binding:"required"`
	Gender       int    `form:"gender" json:"gender" uri:"gender" xml:"gender" binding:"required"`
	AvatarStatus int    `form:"avatar_status" json:"avatar_status" uri:"avatar_status" xml:"avatar_status" binding:"required"`
}
type CreateUserResp struct {
	Ts   int      `json:"ts"`
	Info Userinfo `json:"info"`
}

//删除用户
type DeleteUserReq struct {
	Uid int `form:"uid" json:"uid" uri:"uid" xml:"uid" binding:"required"`
}
type DeleteUserResp struct {
	Rows int64 `json:"rows"`
	Ts   int64 `json:"ts"`
}

//修改用户
type UpdateUserReq struct {
	Uid  int    `form:"uid" json:"uid" uri:"uid" xml:"uid" binding:"required"`
	Name string `form:"name" json:"name" uri:"name" xml:"name" binding:"required"`
}
type UpdateUserResp struct {
	Rows int64 `json:"rows"`
	Ts   int64 `json:"ts"`
}

//获取用户信息
type UserinfoReq struct {
	Uid int `form:"uid" json:"uid" uri:"uid" xml:"uid" binding:"required"`
}
type UserinfoResp struct {
	Ts   int      `json:"ts"`
	Info Userinfo `json:"info"`
}

//获取用户信息
type ListReq struct {
	Uid int `form:"uid" json:"uid" uri:"uid" xml:"uid" binding:"required"`
}
type ListResp struct {
	Ts   int        `json:"ts"`
	List []Userinfo `json:"info"`
}

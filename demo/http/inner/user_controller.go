package inner

import (
	"demo/entity"
	"demo/model/pgsql"
	"github.com/gin-gonic/gin"
	"github.com/stevenyao001/edgeCommon/http"
	"github.com/stevenyao001/edgeCommon/logger"
)

type UserController struct {
}

func (u *UserController) Create(c *gin.Context) {
	var req = entity.CreateUserReq{}
	var resp = entity.CreateUserResp{}
	var err error

	if err = c.Bind(&req); err != nil {
		logger.ErrorLog("UserController-Create", "绑定参数错误", "123", err)
		http.RespError(c, http.RespCodeParamErr, nil)
		return
	}

	uid, err := pgsql.NewUserPgs().Create(req.Name, req.Gender, req.Birthday, req.AvatarStatus)
	if err != nil {
		http.RespError(c, http.RespCodeUnknownErr, nil)
		return
	}

	resp.Info = entity.Userinfo{
		Uid:  int(uid),
		Name: req.Name,
	}

	http.RespSuccess(c, resp)
}

//
//func (u *UserController) Delete(c *gin.Context) {
//	var req = entity.DeleteUserReq{}
//	var resp = entity.DeleteUserResp{}
//	var err error
//
//	if err = c.Bind(&req); err != nil {
//		component.ErrorLog("UserController-Delete", "绑定参数错误", "123", err)
//		component.RespError(c, cont.RespCodeParamErr, nil)
//		return
//	}
//
//	rows, err := pgsql.NewUserPgs().Delete(req.Uid)
//	if err != nil {
//		component.RespError(c, cont.RespCodeUnknownErr, nil)
//		return
//	}
//
//	resp.Rows = rows
//	resp.Ts = time.Now().UnixNano() / 1e6
//
//	component.RespSuccess(c, resp)
//}
//
//func (u *UserController) Update(c *gin.Context) {
//	var req = entity.UpdateUserReq{}
//	var resp = entity.UpdateUserResp{}
//	var err error
//
//	if err = c.Bind(&req); err != nil {
//		component.ErrorLog("UserController-Update", "绑定参数错误", "123", err)
//		component.RespError(c, cont.RespCodeParamErr, nil)
//		return
//	}
//
//	rows, err := pgsql.NewUserPgs().Update(req.Uid, req.Name)
//	if err != nil {
//		component.RespError(c, cont.RespCodeUnknownErr, nil)
//		return
//	}
//
//	resp.Rows = rows
//	resp.Ts = time.Now().UnixNano() / 1e6
//
//	component.RespSuccess(c, resp)
//}
//
//func (u *UserController) Userinfo(c *gin.Context) {
//	var req = entity.UserinfoReq{}
//	var resp = entity.UserinfoResp{}
//	var err error
//
//	if err = c.Bind(&req); err != nil {
//		component.ErrorLog("UserController-Userinfo", "绑定参数错误", "123", err)
//		component.RespError(c, cont.RespCodeParamErr, nil)
//		return
//	}
//
//	userStr, err := new(redis.UserRds).GetUser(req.Uid)
//	if err != nil && err != redis2.Nil {
//		component.RespError(c, cont.RespCodeParamErr, nil)
//		return
//	}
//
//	if err == nil {
//		userRes := pgsql.UserPgs{}
//		_ = json.Unmarshal([]byte(userStr), &userRes)
//
//		resp.Info = entity.Userinfo{
//			Uid:  userRes.Uid,
//			Name: userRes.Name,
//		}
//		component.RespSuccess(c, resp)
//
//		fmt.Println("缓存有，直接返回-----------")
//		return
//	}
//
//	res, err := pgsql.NewUserPgs().Userinfo(req.Uid)
//	if err != nil && err != sql.ErrNoRows {
//		component.RespError(c, cont.RespCodeUnknownErr, nil)
//		return
//	}
//
//	if err == sql.ErrNoRows {
//		component.ErrorLog("UserController-Userinfo", "用户不存在", "123", err)
//		component.RespError(c, cont.RespCodeUserNotExists, nil)
//		return
//	}
//	resp.Info = entity.Userinfo{
//		Uid:  res.Uid,
//		Name: res.Name,
//	}
//
//	//设置缓存
//	userBy, _ := json.Marshal(res)
//	_ = new(redis.UserRds).SetUser(req.Uid, string(userBy))
//	fmt.Println("缓存没有，设置缓存-----------")
//
//	component.RespSuccess(c, resp)
//}
//
//func (u *UserController) List(c *gin.Context) {
//	var req = entity.ListReq{}
//	var resp = entity.ListResp{}
//	var err error
//
//	if err = c.Bind(&req); err != nil {
//		component.ErrorLog("UserController-List", "绑定参数错误", "123", err)
//		component.RespError(c, cont.RespCodeParamErr, nil)
//		return
//	}
//
//	res, err := pgsql.NewUserPgs().List(req.Uid)
//	if err != nil {
//		component.RespError(c, cont.RespCodeUnknownErr, nil)
//		return
//	}
//
//	var userList = make([]entity.Userinfo, 0, len(res))
//	for k := range res {
//		tmp := entity.Userinfo{}
//		tmp.Uid = res[k].Uid
//		tmp.Name = res[k].Name
//		userList = append(userList, tmp)
//	}
//	resp.List = userList
//	component.RespSuccess(c, resp)
//}

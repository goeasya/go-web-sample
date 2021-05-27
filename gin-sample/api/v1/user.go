package v1

import (
	"net/http"

	"gin-sample/handler"
	"gin-sample/model/request"
	"gin-sample/util"

	"github.com/gin-gonic/gin"
)

// @Tags user
// @Summary 删除用户,软删除
// @Produce  application/json
// @Param Body body request.UserDelReq true "param for body"
// @Success 200 {string}  string "{"success": true, "code": 200,"data":[],"msg":"success"}"
// @Router /api/v1/user/ [DELETE]
func DeleteUser(c *gin.Context) {
	var arg request.UserDelReq
	if err := c.ShouldBindJSON(&arg); err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := handler.GetUserHandler().DeleteUser(arg.UserId); err != nil {
		util.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.ResponseSuccess(c, http.StatusOK, nil)
}

// @Tags user
// @Summary 修改用户信息
// @Produce  application/json
// @Param Body body request.UserUpdateReq true "param for body"
// @Success 200 {string}  string "{"success": true, "code": 200,"data":[],"msg":"success"}"
// @Router /api/v1/user/ [PUT]
func UpdateUser(c *gin.Context) {
	var args request.UserUpdateReq
	if err := c.ShouldBindJSON(&args); err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := handler.GetUserHandler().UpdateUser(&args); err != nil {
		util.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.ResponseSuccess(c, http.StatusOK, nil)
	return
}

// @Tags user
// @Summary 获取用户细节
// @Produce  application/json
// @Param username query string true "username"
// @Success 200 {string}  string "{"success": true, "code": 200,"data":[],"msg":"success"}"
// @Router /api/v1/user/detail [GET]
func GetUserDetail(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		util.ResponseError(c, http.StatusBadRequest, "invalid username")
		return
	}
	detail, err := handler.GetUserHandler().GetUserDetail(username)
	if err != nil {
		util.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	util.ResponseSuccess(c, http.StatusOK, detail)
	return
}

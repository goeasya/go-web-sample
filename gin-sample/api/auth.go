package api

import (
	"net/http"

	"gin-sample/handler"
	"gin-sample/middleware"
	"gin-sample/model/request"
	"gin-sample/model/response"
	"gin-sample/types"
	"gin-sample/util"

	"github.com/gin-gonic/gin"
)

// @Tags user
// @Summary 用户注册
// @Produce  application/json
// @Param Body body request.UserRegisterReq true "param for body"
// @Success 200 {string}  string "{"success": true, "code": 200,"data":[],"msg":"success"}"
// @Router /user/register [POST]
func Register(c *gin.Context) {
	var args request.UserRegisterReq
	if err := c.ShouldBindJSON(&args); err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := handler.GetUserHandler().RegisterUser(&args); err != nil {
		util.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.ResponseSuccess(c, http.StatusOK, nil)
}

// @Tags user
// @Summary 用户登录
// @Produce  application/json
// @Param Body body request.UserLoginReq true "param for body"
// @Success 200 {object}  response.LoginToken "{"success": true, "code": 200,"data":[],"msg":"success"}"
// @Router /user/login [POST]
func Login(c *gin.Context) {
	var args request.UserLoginReq
	if err := c.ShouldBindJSON(&args); err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := handler.GetUserHandler().Login(args.UserName, args.Password)
	if err != nil {
		if err == types.ErrInvalidUserOrPassword {
			util.ResponseError(c, http.StatusUnauthorized, err.Error())
			return
		}
		util.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := middleware.GenerateToken(user)
	if err != nil {
		util.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	loginToken := &response.LoginToken{Token: token}
	util.ResponseSuccess(c, http.StatusOK, loginToken)
	return
}

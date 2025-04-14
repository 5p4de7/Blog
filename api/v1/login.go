package v1

import (
	"Blog/middleware"
	"Blog/model"
	"Blog/utils/errormsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data model.User
	var Token string
	c.ShouldBindJSON(&data)

	code := model.CheckLogin(data.Username,data.PassWord)
	if code ==errormsg.SUCCSE{
		Token,code = middleware.SetToken(data.Username)

	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errormsg.GetErrMsg(code),
		"token": Token,
	})
}
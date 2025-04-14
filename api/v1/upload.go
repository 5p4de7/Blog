package v1

import (
	"Blog/model"
	"Blog/utils/errormsg"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload1(c *gin.Context){
	file,fileHeader,_ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url,code :=model.UpLoadFile(file,fileSize)
	fmt.Println("Upload result: url =", url, "code =", code)

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message": errormsg.GetErrMsg(code),
		"jj":452,
		"url":url,
	})
}
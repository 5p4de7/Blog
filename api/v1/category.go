package v1

import (
	"Blog/model"
	"Blog/utils/errormsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询分类是否存在
func UserCategory(c *gin.Context) {
	//
}

// 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	//TODO :检查用户名和密码是否为空
	if code == errormsg.SUCCSE {
		model.CreateCategory(&data)
	}
	if code == errormsg.ERROR_CATENAME_USED {
		code = errormsg.ERROR_CATENAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})
}


// 查询分类列表
func GetCategorys(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1 //可以取消gorm的limit
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data,total := model.GetCategorys(pageSize, pageNum)
	code := errormsg.SUCCSE

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total": total,
		"message": errormsg.GetErrMsg(code),
	})
}

// 编辑分类名
// TODO 改造编辑用户接口，使得不用改用户名也成立
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errormsg.SUCCSE {
		model.EditCategory(id, &data)
	}
	if code == errormsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})

}

// 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteCategory(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}


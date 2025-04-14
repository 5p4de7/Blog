package model
import (
	"Blog/utils/errormsg"
	_"log"


	"gorm.io/gorm"
)

type Category struct {
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errormsg.ERROR_CATENAME_USED //1001
	}
	return errormsg.SUCCSE
}

// 新增分类
func CreateCategory(data *Category) int {
	//data.PassWord = ScryptPw(data.PassWord)
	//data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR //500
	}
	return errormsg.SUCCSE
}

// 查询分类列表
func GetCategorys(pageSize int, pageNum int) ([]Category,int64) {
	var cate []Category
	var total int64
	//注意一下下面这一行

	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	err := db.Limit(pageSize).Offset(offset).Find(&cate).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil,0
	}

	return cate,total
}

// 编辑分类
func EditCategory(id int, data *Category) int {
	//注意gin框架的问题
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"]=data.Name

	err :=db.Model(&cate).Where("id=?",id).Updates(maps).Error

	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

// 删除分类
func DeleteCategory(id int) int {
	var cate Category
	err := db.Where("id=?", id).Delete(&cate).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}
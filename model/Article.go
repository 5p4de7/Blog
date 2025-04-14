package model

import (
	"Blog/utils/errormsg"
	_ "log"

	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// 新增文章
func CreateArticle(data *Article) int {
	//data.PassWord = ScryptPw(data.PassWord)
	//data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR //500
	}
	return errormsg.SUCCSE
}

// 查询分类所有文章
func GetCateArticle(id int, pageSize int, pageNum int) ([]Article, int,int64) {
	var cateArtList []Article
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	err := db.Preload("Category").Limit(pageSize).Offset(offset).Where("cid=?", id).Find(&cateArtList).Count(&total).Error
	if err != nil {
		return nil, errormsg.ERROR_CATE_NOT_EXIST,0
	}
	return cateArtList, errormsg.SUCCSE,total

}

// 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id=?", id).First(&art).Error
	if err != nil {
		return art, errormsg.ERROR_ART_NOT_EXIST
	}
	return art, errormsg.SUCCSE
}

// 查询文章列表
func GetArticle(pageSize int, pageNum int) ([]Article, int,int64) {
	var articleList []Article
	var total int64
	//注意一下下面这一行
	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}

	err := db.Preload("Category").Limit(pageSize).Offset(offset).Find(&articleList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errormsg.ERROR,0
	}

	return articleList, errormsg.SUCCSE,total
}

// 编辑文章
func EditArticle(id int, data *Article) int {
	//注意gin框架的问题
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := db.Model(&art).Where("id=?", id).Updates(maps).Error

	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

// 删除文章
func DeleteArticle(id int) int {
	var art Article
	err := db.Where("id=?", id).Delete(&art).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

package model

import (
	"Blog/utils/errormsg"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/scrypt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	PassWord string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int:DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// 查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errormsg.ERROR_USERNAME_USED //1001
	}
	return errormsg.SUCCSE
}

// 新增用户
func CreateUser(data *User) int {
	data.PassWord = ScryptPw(data.PassWord)
	//data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR //500
	}
	return errormsg.SUCCSE
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) ([]User,int64 ){
	var users []User
	var total int64
	//注意一下下面这一行

	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	err := db.Limit(pageSize).Offset(offset).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil,0
	}

	return users,total
}

// 编辑用户
func EditUser(id int, data *User) int {
	//注意gin框架的问题
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id=?", id).Updates(maps).Error

	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

// 删除
func DeleteUser(id int) int {
	var user User
	err := db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

// 密码加密
//func (u *User) BeforeSave() {
//	u.PassWord = ScryptPw(u.PassWord)
//}

// TODO  使用bcrypt创建一个加密方法
func ScryptPw(password string) string {
	const Keylen = 10
	salt := make([]byte, 8)

	salt = []byte{12, 55, 6, 32, 18, 54, 12, 9}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, Keylen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)

	return (fpw)

}

// 登录验证
func CheckLogin(username string ,password string)int{
	var user User

	db.Where("username=?",username).First(&user)
	if user.ID == 0{
		return errormsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.PassWord{
		return errormsg.ERROE_PASSWORD_WRONG
	}
	if user.Role !=1{
		return errormsg.ERROR_USER_NO_RIGHT
	}
	return errormsg.SUCCSE
}
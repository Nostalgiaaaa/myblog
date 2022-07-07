package model

import (
	"encoding/base64"
	"log"
	"myblog/utils/errcode"

	"golang.org/x/crypto/scrypt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null" json:"username"  `
	Password string `gorm:"type:varchar(20);not null " json:"password"  `
	Role     int    `gorm:"type:int " json:"role"`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) *errcode.Error {
	var users User
	db.Select("id").Where("user_name = ?", name).First(&users)
	//大于0  用户已存在 返回状态码
	if users.ID > 0 {
		return errcode.ErrorUserNameUsed
	}
	return errcode.Success
}

// CreateUser 新增用户
func CreateUser(data *User) *errcode.Error {
	//data.Password = ScryptPw(data.Password)
	err := db.Create(data).Error
	if err != nil {
		return errcode.ServerError // 500
	}
	return errcode.Success
}

// GetUsers 查询用户列表
func GetUsers(pageSize, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// EditUser 编辑用户信息
func EditUser(id int, data *User) *errcode.Error {
	var user User
	var maps = make(map[string]any)
	maps["user_name"] = data.UserName
	maps["password"] = data.Password
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errcode.ServerError
	}
	return errcode.Success
}

// DeleteUser 删除用户
func DeleteUser(id int) *errcode.Error {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errcode.ServerError
	}
	return errcode.Success
}

// ScryptPw 加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{21, 213, 65, 81, 43, 56, 48, 73}

	hashPassword, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(hashPassword)
}

// BeforeSave gorm 钩子函数
func (u *User) BeforeSave(*gorm.DB) error {
	u.Password = ScryptPw(u.Password)
	return nil
}

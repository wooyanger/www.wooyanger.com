package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// 定义用户模型
type User struct {
	Id				int64		`xorm:"int(20) notnull autoincr pk"`
	Username		string		`xorm:"varchar(255) notnull unique"`
	PasswordHash	string		`xorm:"varchar(255) notnull"`
	Email			string		`xorm:"varchar(255) notnull"`
	CreateAt		time.Time	`xorm:"datetime notnull"`
}

// 获取用户信息
func (u *User) Get(uid int64) *User {
	newUser := &User{Id: uid}
	ok, err := x.Get(newUser)
	if ok && err == nil {
		return newUser
	}
	return nil
}

// 生成密码
func (u *User) GeneratePasswordHash() {
	hash := md5.New()
	hash.Write([]byte(u.PasswordHash))
	u.PasswordHash = hex.EncodeToString(hash.Sum(nil))
}


// 验证密码是否正确
func (u *User) ValidatePassword(password string) bool {
	newUser := &User{PasswordHash: password}
	newUser.GeneratePasswordHash()
	if u.PasswordHash == newUser.PasswordHash {
		return true
	}
	return false
}

// 用户登录
func UserLogin(username string, password string) (*User, error) {
	var newUser *User
	if len(username) == 0 || len(password) == 0 {
		return nil, fmt.Errorf("username or password is not empty.")
	}
	if strings.Contains(username, "@") {
		newUser = &User{Email: username}
	} else {
		newUser = &User{Username: username}
	}
	hasUser, err := x.Get(newUser)
	if err != nil {
		return nil, fmt.Errorf("get user record: %v", err)
	}

	if hasUser {
		if newUser.ValidatePassword(password) {
			return newUser, nil
		}
		return nil, fmt.Errorf("username or password is invalid.")
	}
	return nil, fmt.Errorf("user is not exist.")
}

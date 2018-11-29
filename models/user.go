package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID				int64		`xorm: "int(12) not null autoincr pk"`
	Username		string		`xorm: "varchar(32) not null unique"`
	PasswordHash	string		`xorm: "varchar(64) not null"`
	Email			string		`xorm: "varchar(64) not null"`
	CreateAt		time.Time	`xorm: "datetime not null"`
}

func (u *User) Get(uid int64) *User {
	newUser := &User{ID: uid}
	ok, err := x.Get(newUser)
	if ok && err == nil {
		return newUser
	}
	return nil
}

func (u *User) GeneratePasswordHash() {
	hash := md5.New()
	hash.Write([]byte(u.PasswordHash))
	u.PasswordHash = hex.EncodeToString(hash.Sum(nil))
}

func (u *User) ValidatePassword(password string) bool {
	newUser := &User{PasswordHash: password}
	newUser.GeneratePasswordHash()
	if u.PasswordHash == newUser.PasswordHash {
		return true
	}
	return false
}

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
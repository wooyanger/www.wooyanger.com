package models

import "time"

type User struct {
	ID				int64		`xorm: "int(12) not null autoincr pk"`
	Username		string		`xorm: "varchar(32) not null unique"`
	PasswordHash	string		`xorm: "varchar(64) not null"`
	Email			string		`xorm: "varchar(64) not null"`
	CreateAt		time.Time	`xorm: "datetime not null"`
}
package models

type Tag struct {
	Id			int64		`xorm: "int(12) not null autoincr pk"`
	Tname		string		`xorm: "varchar(64) not null"`
}
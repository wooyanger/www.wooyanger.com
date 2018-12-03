package models

type PostMapTag struct {
	Id			int64		`xorm: "int(12) not null autoincr pk"`
	Pid			int64		`xorm: "int(12) not null"`
	Tid			int64		`xorm: "int(12) not null"`
}
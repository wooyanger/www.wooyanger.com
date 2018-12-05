package models

type PostMapTag struct {
	Id			int64		`xorm: "int(12) not null autoincr pk"`
	Pid			int64		`xorm: "int(12) not null"`
	Tid			int64		`xorm: "int(12) not null"`
}

func GetAllPidByTid(id int64) []PostMapTag {
	allPostMapTag := make([]PostMapTag, 0)
	x.Find(&allPostMapTag)
	return allPostMapTag
}
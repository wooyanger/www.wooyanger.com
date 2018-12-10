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

func CreatePostMapTag(p *PostMapTag) error {
	if _, err := x.Insert(p); err != nil {
		return err
	}
	return nil
}

func DeletePostMapTag(p *PostMapTag) error {
	if _, err := x.Delete(p); err != nil {
		return err
	}
	return nil
}
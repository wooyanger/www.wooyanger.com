package models

import (
	"fmt"
	"time"
)

type Post struct {
	ID			int64		`xorm: "int(12) not null autoincr pk"`
	Title		string		`xorm: "varchar(256) not null unique"`
	Content		string		`xorm: "text not null"`
	CreateAt	time.Time	`xorm: "datetime not null"`
	UpdateAt	time.Time	`xorm: "datetime not null"`
	Tags		string		`xorm: "varchar(256) not null"`
	UID			int64		`xorm: "int(12) not null"`
}

func IsTitleExist(title string) (bool, error) {
	hasTitle, err := x.Get(&Post{Title: title})
	if err != nil {
		return false, err
	} else if hasTitle {
		return true, nil
	}
	return false, nil
}

func CreatePost(p *Post) error {
	isTitle, err := IsTitleExist(p.Title)
	if err != nil {
		return err
	} else if isTitle {
		return fmt.Errorf("title has been used [title: %s]", p.Title)
	}
	return nil
}
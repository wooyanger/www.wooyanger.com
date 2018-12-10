package models

import (
	"fmt"
	"time"
)

// 定义Tag 模型
type Tag struct {
	Id			int64		`xorm: "int(12) not null autoincr pk"`
	Tname		string		`xorm: "varchar(64) not null"`
	CreateAt	time.Time	`xorm: "datetime not null"`
	UpdateAt	time.Time	`xorm: "datetime not null"`
}

// 判断 Tag 名是否以及存在
func isTagExist(t string) (bool, error) {
	hasTag, err := x.Get(&Tag{Tname: t})
	if err != nil {
		return false, err
	} else if hasTag {
		return true, nil
	}
	return false, nil
}

func (t *Tag) GetAllTag() []Tag {
	tags := make([]Tag, 0)
	x.Find(&tags)
	return tags
}

// 新增 Tag
func NewTag(t *Tag) error {
	isTag, err := isTagExist(t.Tname)
	if err != nil {
		return err
	} else if isTag {
		return fmt.Errorf("tag has been existed [tag: %v]", t.Tname)
	}
	sess := x.NewSession()
	defer sess.Close()
	if _, err = sess.Insert(t); err != nil {
		return err
	}
	return sess.Commit()
}

// 更新 Tag
func UpdateTag(t *Tag) error {
	t.UpdateAt = time.Now()
	_, err := x.ID(t.Id).AllCols().Update(t)
	return err
}

// 删除 Tag
func DeleteTag(t *Tag) error {
	sess := x.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}
	if _, err := x.ID(t.Id).Delete(t);err != nil {
		return err
	}

	if err := sess.Commit();err != nil {
		return err
	}
	return nil
}
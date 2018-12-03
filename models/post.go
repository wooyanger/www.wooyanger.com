package models

import (
	"fmt"
	"time"
)

// 文章模型
type Post struct {
	Id			int64		`xorm: "int(12) not null autoincr pk"`
	Title		string		`xorm: "varchar(256) not null unique"`
	Content		string		`xorm: "text not null"`
	CreateAt	time.Time	`xorm: "datetime not null"`
	UpdateAt	time.Time	`xorm: "datetime not null"`
	//Tags		string		`xorm: "varchar(256) not null"`
	Uid			int64		`xorm: "int(12) not null"`
}

// 获取所有文章
func (p *Post) GetAll() []Post {
	postList := []Post{}
	x.Find(&postList)
	return postList
}


// 检测文章标题是否存在
func IsTitleExist(title string) (bool, error) {
	hasTitle, err := x.Get(&Post{Title: title})
	if err != nil {
		return false, err
	} else if hasTitle {
		return true, nil
	}
	return false, nil
}

// 创建新文章
func CreatePost(p *Post) error {
	isTitle, err := IsTitleExist(p.Title)
	if err != nil {
		return err
	} else if isTitle {
		return fmt.Errorf("title has been used [title: %s]", p.Title)
	}
	sess := x.NewSession()
	defer sess.Close()
	if _, err = sess.Insert(p); err != nil {
		return err
	}
	return sess.Commit()
}

// 更新文章
func UpdatePost(p *Post) error {
	p.UpdateAt = time.Now()
	_, err := x.ID(p.Id).AllCols().Update(p)
	return err
}

// 删除文章
func DeletePost(p *Post) error {
	sess := x.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}
	if _, err := x.ID(p.Id).Delete(p);err != nil {
		return err
	}

	if err := sess.Commit();err != nil {
		return err
	}
	return nil
}
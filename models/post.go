package models

import (
	"fmt"
	"time"
)

// 文章模型
type Post struct {
	Id			int64		`xorm:"int(20) notnull autoincr pk"`
	Title		string		`xorm:"varchar(255) notnull unique"`
	Content		string		`xorm:"text notnull"`
	CreateAt	time.Time	`xorm:"datetime notnull"`
	UpdateAt	time.Time	`xorm:"datetime notnull"`
	Uid			int64		`xorm:"int(20) notnull"`
}

// 获取所有文章
func (p *Post) GetAllPost() []Post {
	posts := make([]Post, 0)
	x.Find(&posts)
	return posts
}

// 获取指定文章
func (p *Post) GetPostById(id int64) *Post {
	post := &Post{Id: id}
	ok, err := x.Get(post)
	if ok && err == nil {
		return post
	} else {
		return nil
	}
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
func CreatePost(p *Post) (int64, error) {
	isTitle, err := IsTitleExist(p.Title)
	if err != nil {
		return 0, err
	} else if isTitle {
		return 0, fmt.Errorf("title has been used [title: %s]", p.Title)
	}
	sess := x.NewSession()
	defer sess.Close()
	if _, err = sess.Insert(p); err != nil {
		return 0, err
	}
	return p.Id, sess.Commit()
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
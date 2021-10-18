package models

import "github.com/astaxie/beego/orm"

type Category struct {
	Id     int
	Pid    int
	Status int
	Name   string `orm:"size(15);unique"`
	Desc   string // 介绍
	Sort   int
	Num    int // 分类下数量
	Icon   string
}

func (m *Category) TableName() string {
	return TNCategory()
}

func (m *Category) GetCategoryList(pid int, status int) (cates []Category, err error) {
	qs := orm.NewOrm().QueryTable(TNCategory())
	if pid > -1 {
		qs = qs.Filter("pid", pid)
	}
	if 1 == status || 2 == status {
		qs = qs.Filter("status", status)
	}
	_, err = qs.OrderBy("sort").All(&cates)
	return
}

func (m *Category) Find(id int) (cate Category) {
	cate.Id = id
	orm.NewOrm().Read(&cate)
	return cate
}
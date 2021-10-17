package models

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

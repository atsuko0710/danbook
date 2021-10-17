package models

type Book struct {
	Id       int
	Name     string
	Identify string `orm:"size(100);unique" json:"identify"` // 唯一标识
	Sort     int
	Desc     string
}

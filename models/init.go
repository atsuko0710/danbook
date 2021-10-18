package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(
		new(Book),
		new(BookCategory),
		new(Category),
	)
}

func TNCategory() string {
	return "category"
}

func TNBook() string {
	return "book"
}

func TNBookCategory() string {
	return "book_category"
}

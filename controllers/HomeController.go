package controllers

import (
	"danbook/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	if cates, err := new(models.Category).GetCategoryList(-1, 1); err == nil {
		c.Data["Cates"] = cates
	} else {
		beego.Error(err.Error())
	}
	c.TplName = "home/list.html"
}

func (c *HomeController) Category() {
	var (
		cid int
		cate models.Category
		// urlPrefix = beego.URLFor("HomeController/Category")
	)

	if cid, _ = c.GetInt("cid"); cid > 0 {
		cateModel := new(models.Category)
		cate = cateModel.Find(cid)
		c.Data["Cate"] = cate
	}
	c.Data["Cid"] = cid


	c.TplName = "home/category.html"
}

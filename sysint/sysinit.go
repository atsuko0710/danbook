package sysint

import (
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

func sysint() {
	uploads := filepath.Join("./", "uploads")
	beego.BConfig.WebConfig.StaticDir["/uploads"] = uploads

	// 注册前端注册函数
	registerFunctions()
}

func registerFunctions() {
	beego.AddFuncMap("cdnjs", func(p string) string  {
		cdn := beego.AppConfig.DefaultString("cdnjs", "")
		if strings.HasPrefix(p, "/") && strings.HasSuffix(cdn, "/") {
			return cdn + string(p[1:])
		}
		return cdn + p
	})
}

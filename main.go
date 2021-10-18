package main

import (
	_ "danbook/routers"
	
	"github.com/astaxie/beego"
	_ "danbook/sysint"
)

func main() {
	beego.Run()
}


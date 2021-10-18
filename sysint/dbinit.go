package sysint

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func dbinit(aliases ...string) {

	isDev := ("dev" == beego.AppConfig.String("runmode"))

	if len(aliases) > 0 {
		for _, alias := range aliases {
			registerDataBase(alias)
			if "w" == alias {
				orm.RunSyncdb("default", false, isDev)
			}
		}
	} else {
		registerDataBase("w")
		orm.RunSyncdb("default", false, isDev)
	}

	if isDev {
		orm.Debug = isDev
	}
}

func registerDataBase(alias string) {
	if len(alias) <= 0 {
		return
	}

	dbAlias := alias
	if "w" == alias || "default" == alias {
		dbAlias = "default"
		alias = "w"
	}

	dbName := beego.AppConfig.String("db_" + alias + "_database")
	dbPass := beego.AppConfig.String("db_" + alias + "_password")
	dbPort := beego.AppConfig.String("db_" + alias + "_port")
	dbHost := beego.AppConfig.String("db_" + alias + "_host")
	dbUser := beego.AppConfig.String("db_" + alias + "_username")

	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8", 30)
}

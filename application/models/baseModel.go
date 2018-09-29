package models

import (
	"github.com/fzpying/beego"
	"github.com/fzpying/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

var (
	dbName     = beego.AppConfig.String("db_name")
	dbUser     = beego.AppConfig.String("db_user")
	dbPassword = beego.AppConfig.String("db_password")
	dbHost     = beego.AppConfig.String("db_host")
	dbPort     = beego.AppConfig.String("db_port")
	dbPrefix   = beego.AppConfig.String("db_prefix")
	dbCharset  = beego.AppConfig.String("db_charset")
	dbDriver   = beego.AppConfig.String("db_driver")
)

func init() {
	// set default database
	orm.RegisterDataBase(
		"default",
		dbDriver,
		dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset+"",
		30)
	registerModel()

	// create table
	orm.RunSyncdb("default", true, true)
}

// register model
func registerModel() {
	orm.RegisterModelWithPrefix("ylx_",
		new(Users),
	)
}

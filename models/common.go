package models

// 通用方法
import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//	"log"
	//	"strconv"
)

func init() {
	//数据库实例化
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	iniconf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		fmt.Println(err)
	}
	var masterHost = iniconf.String("master::mysqlurls")
	var masterUser = iniconf.String("master::mysqluser")
	var masterPass = iniconf.String("master::mysqlpass")
	var masterDb = iniconf.String("master::mysqldb")
	var slaverHost = iniconf.String("slaver::mysqlurls")
	var slaverUser = iniconf.String("slaver::mysqluser")
	var slaverPass = iniconf.String("slaver::mysqlpass")
	var slaverDb = iniconf.String("slaver::mysqldb")
	masterConn := masterUser + ":" + masterPass + "@" + masterHost + "/" + masterDb + "?charset=utf8&loc=Asia%2FShanghai"
	slaverConn := slaverUser + ":" + slaverPass + "@" + slaverHost + "/" + slaverDb + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", masterConn)
	orm.RegisterDataBase("slaver", "mysql", slaverConn)
}

//Init some database settings
package dao

import (
	"janlent-metro/model"
	//"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kylelemons/go-gypsy/yaml"
)

var ConnectionString string
var DbType string
var RedisAddress string

//1. Read some database configuration from settings file
//2. init database orm object
func init() {
	config, err := yaml.ReadFile("settings.yaml")
	if err != nil {
		//TODO: Log the error and panic error
		logs.Error("Read settings.yaml file error, the system cannot run!")
		panic(err)
	}
	ConnectionString, _ = config.Get("connection")
	DbType, _ = config.Get("db_type")
	RedisAddress, _ = config.Get("redis_address")

	err = orm.RegisterDataBase("default", DbType, ConnectionString)
	if err != nil {
		//TODO: the err means that the database cannot be opened
		logs.Error("Cannot open the database which connection : ", ConnectionString)
		panic(err)
	}

	//Register passenger
	orm.RegisterModel(
		new(model.Passenger),
		new(model.MetroEnter),
		new(model.Station),
		new(model.MetroExit),
		new(model.PassengerPath),
		new(model.MetroPath),
		new(model.MetroPathDetail),
		new(model.TimeTableDetail),
		new(model.MetroTransportation))
	//orm.DefaultTimeLoc = time.Local
	orm.SetMaxOpenConns("default", 30) //设置数据库最大连接数
	orm.SetMaxIdleConns("default", 30) //设置数据库最大空闲连接数

	orm.RunSyncdb("default", false, true)
	logs.Info("Database initialization is ok!")
}

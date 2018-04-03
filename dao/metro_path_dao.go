package dao

import (
	"janlent-metro/model"
	//"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type MetroPathDao struct {

}

func (m *MetroPathDao) GetPathDetails(key string) ([]model.MetroPathDetail) {
	client, err := redis.Dial("tcp", RedisAddress)
	if err != nil {
		//TODO:
		fmt.Println("Connect to redis error: ", RedisAddress)
		logs.Error(err)
		return nil
	}
	defer client.Close()
	//Switch the database
	client.Do("SELECT", 1)
	//Get the json string by key
	strJson, err := redis.String(client.Do("GET", key))
	if err != nil {
		//TODO:
		logs.Error(err)
		return nil
	}
	var list []model.MetroPathDetail
	err = json.Unmarshal([]byte(strJson), &list)
	if err != nil {
		//TODO:
		logs.Error(err)
		return nil
	}
	return list
}

func (m *MetroPathDao) GetMetroPath(id int) (*model.MetroPath) {
	var path model.MetroPath
	db := orm.NewOrm()
	err := db.QueryTable(new(model.MetroPath)).Filter("Id", id).One(&path)
	if err != nil {
		logs.Error(err)
		return nil
	}

	return &path
}

func (m *MetroPathDao) GetMetroPaths() ([]model.MetroPath) {
	var list []model.MetroPath
	db := orm.NewOrm()
	_, err := db.QueryTable(new(model.MetroPath)).All(&list)
	if err != nil {
		logs.Error(err)
		return nil
	}
	return list
}

func (m *MetroPathDao) GetMetroPathDetails(pathId int) ([]model.MetroPathDetail) {
	var list []model.MetroPathDetail
	db := orm.NewOrm()
	_, err := db.QueryTable(new(model.MetroPathDetail)).Filter("PathId", pathId).All(&list)
	if err != nil {
		logs.Error(err)
		return nil
	}
	return list
}
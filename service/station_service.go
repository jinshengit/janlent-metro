package service

import (
	"janlent-metro/model"
	"janlent-metro/dao"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/logs"
	"encoding/json"
)

var StationMap map[string]model.Station
var StationIdMap map[int]model.Station
var MetroStationMap map[string]model.MetroStation

func init() {
	StationMap = dao.GetAllStations()
	getMetroStationMapFromRedis()
	GetStationIdMap()
}

func GetStationIdMap() {
	StationIdMap = make(map[int]model.Station)
	for _, v := range StationMap {
		if _, ok := StationIdMap[v.StationTrueNo]; !ok {
			StationIdMap[v.StationTrueNo] = v
		}
	}
}

func getMetroStationMapFromRedis() {
	client, _ := redis.Dial("tcp", dao.RedisAddress)
	defer client.Close()

	strJson, err := redis.String(client.Do("GET", "station_map"))
	if err != nil {
		logs.Error(err)
		return
	}
	if err = json.Unmarshal([]byte(strJson), &MetroStationMap); err != nil {
		logs.Error(err)
	}
}
package dao

import (
	"janlent-metro/model"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type StationDao struct {

}

func GetAllStations() (map[string]model.Station) {
	db := orm.NewOrm()
	var list []model.Station

	_, err := db.QueryTable(new(model.Station)).All(&list)
	if err != nil {
		//TODO:
		logs.Error(err)
		return nil
	}

	var stationMap = make(map[string]model.Station)
	for _, station := range list {
		if _, ok := stationMap[station.StationNo]; !ok {
			stationMap[station.StationNo] = station
		}
	}

	return stationMap
}
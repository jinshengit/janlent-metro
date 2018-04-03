package dao

import (
	"janlent-metro/model"
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type PassengerPathDao struct {

}

func (p *PassengerPathDao) InsertPassengerPaths(paths []model.PassengerPath) (bool, error) {
	if paths == nil || len(paths) <= 0 {
		return false, errors.New("Null passenger path array")
	}
	db := orm.NewOrm()
	result, err := db.InsertMulti(len(paths), paths)
	return result > 0, err
}

func (p *PassengerPathDao) CountPassengerByStation(station string, startTime, endTime time.Time) ([]model.PassengerCount) {
	sql := "SELECT COUNT(id) AS count, station, next_station, direction FROM passenger_path WHERE arrive_time > ? AND arrive_time < ? AND station = ? GROUP BY station, next_station, direction"
	db := orm.NewOrm()
	var list []model.PassengerCount
	db.Raw(sql, startTime, endTime, station).QueryRows(&list)
	return list
}
package dao

import (
	"janlent-metro/model"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"errors"
	"strconv"
)

type PassengerDao struct {

}

//
func (p *PassengerDao) InsertPassengers(passengers []model.Passenger) (bool, error) {
	if passengers == nil || len(passengers) <= 0 {
		return false, errors.New("Null passenger array")
	}
	db := orm.NewOrm()
	result, err := db.InsertMulti(len(passengers), passengers)
	return result > 0, err
}

func (p *PassengerDao) UpdateEnterDataFlag(ids []int64) (bool, error) {
	if ids == nil || len(ids) <= 0 {
		return false, errors.New("Null id list")
	}

	sql := "UPDATE TBL_METRO_ENTER_ALL_TEST201510 SET flag = 2 WHERE "
	for i := 0; i < len(ids); i++ {
		if i == 0 {
			sql += " id = " + strconv.FormatInt(ids[i], 10)
		} else {
			sql += " OR id = " + strconv.FormatInt(ids[i], 10)
		}
	}
	db := orm.NewOrm()
	result, err := db.Raw(sql).Exec()
	if err != nil {
		//TODO: log error message
		logs.Error(err)
		return false, err
	}
	count, err := result.RowsAffected()
	return count > 0, err
}

func (p *PassengerDao) QueryPassengers(startId, endId int64) ([]model.Passenger, error) {
	var list []model.Passenger
	db := orm.NewOrm()
	sql := "SELECT id, card_id, ticket_type, enter_station_id, enter_time, exit_station_id, exit_time, stmt_day FROM passenger WHERE id >= ? AND id <= ?"
	_, err := db.Raw(sql, startId, endId).QueryRows(&list)
	if err != nil {
		//TODO:
		logs.Error(err)
		return nil, err
	}
	return list, nil
}

func (p *PassengerDao) QueryEnterDataById(startId, endId int64) ([]model.MetroEnter, error) {
	var list []model.MetroEnter
	db := orm.NewOrm()
	sql := "SELECT CARDID, TICKET_TYPE, EQUIP_ID, DEAL_TIME, TRANS_TYPE, STMT_DAY, STATION_ID, id FROM TBL_METRO_ENTER_ALL_TEST201510 WHERE id >= ? AND id <= ?"
	_, err := db.Raw(sql, startId, endId).QueryRows(&list)
	if err != nil {
		//TODO:
		logs.Error(err)
		return nil, err
	}
	return list, nil
}

//查询所有Flag为1的进站数据
func (p *PassengerDao) QueryUnprocessdEnterData(count int) ([]model.MetroEnter, error) {
	var list []model.MetroEnter

	db := orm.NewOrm()

	_, err := db.QueryTable(new(model.MetroEnter)).Filter("flag", 1).Limit(count).All(&list)
	if err != nil {
		//TODO: log error message
		logs.Error(err)
		return nil, err
	}
	return list, nil
}

//通过卡号和日期码获取出站数据
func (p *PassengerDao) QueryExitData(cardId, day string) ([]model.MetroExit, error) {
	var list []model.MetroExit
	db := orm.NewOrm()
	_, err := db.QueryTable(new(model.MetroExit)).Filter("CardId", cardId).Filter("StmtDay", day).All(&list)
	if err != nil {
		//TODO: log error message
		logs.Error(err)
		return nil, err
	}
	return list, nil
}

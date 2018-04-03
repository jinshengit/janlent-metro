package dao

import (
	"janlent-metro/model"
	"errors"
	"github.com/astaxie/beego/orm"
)

type MetroTransportDao struct {

}

func (m *MetroTransportDao) InsertMetroTransportations(list []model.MetroTransportation) (bool, error) {
	if list == nil || len(list) <= 0 {
		return false, errors.New("Null transportation array")
	}

	db := orm.NewOrm()
	result, err := db.InsertMulti(len(list), list)
	return result > 0, err
}
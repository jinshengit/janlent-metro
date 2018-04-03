package model

import "time"

type PassengerPath struct {
	Id            int64 `orm:"auto"`
	CardId        string
	ArriveTime    time.Time
	Station       string
	NextStation   string
	Duration      int
	Direction     string
	Line          int
	PathIndex     int //路径的索引号
	IsChangePoint int //是否是换乘站的标识
}

type PassengerCount struct {
	Count       int
	Station     string
	NextStation string
	Direction   string
}

func (this *PassengerPath) TableName() string {
	return "passenger_path"
}

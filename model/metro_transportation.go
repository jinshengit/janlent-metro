package model

import "time"

//地铁运量
type MetroTransportation struct {
	Id                     int       `orm:"auto"`
	StationName            string    `orm:"column(station_name)"`
	NextStation            string    `orm:"column(next_station)"`
	StartTime              time.Time `orm:"column(start_time)"`
	EndTime                time.Time `orm:"column(end_time)"`
	Direction              string    `orm:"column(direction)"`
	TrainCount             int       `orm:"column(train_count)"`
	ExchangePassengerCount int       `orm:"column(exchange_passenger_count)"` //换乘人数
	AboardPassengerCount   int       `orm:"column(aboard_passenger_count)"`   //上车人数
	GetOffPassengerCount   int       `orm:"column(get_off_passenger_count)"`  //出站人数
	PassengerOnTrainCount  int       `orm:"column(passenger_on_train_count)"` //理论在车上的人数
}

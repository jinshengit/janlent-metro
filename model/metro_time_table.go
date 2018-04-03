package model

type TimeTableDetail struct {
	Id int `orm:"auto"`
	TimeTableType int `orm:"column(timetable_type)"`
	LineNo int `orm:"column(line_no)"`
	DepartureStation string `orm:"column(Departure_station)"`
	DestinationStation string `orm:"column(destination_station)"`
	ArrivalTime string `orm:"column(arrival_time)"`
	DepartureTime string `orm:"column(departure_time)"`
	ResidenceTime string `orm:"column(residence_time)"`
	Direction string `orm:"column(udcode)"`
}

func (this *TimeTableDetail) TableName() string {
	return "metro_timetable_detail"
}
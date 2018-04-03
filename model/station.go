package model

type Station struct {
	StationNo     string `orm:"column(StationNo)"`
	StationName   string `orm:"column(StationName)"`
	LineNo        string `orm:"column(LineNO)"`
	StationTrueNo int    `orm:"column(StationTrueNo)"`
	Id            int    `orm:"auto"`
}

func (this *Station) TableName() string {
	return "master_station_info"
}

type MetroStation struct {
	Id           int
	Name         string
	StationLinks []StationLink
}

type StationLink struct {
	FromStation, ToStation string
	Line                   string  //使用的地铁线路
	Wight                  float32 //权重
	Flag                   string  //上行下行的标识
}

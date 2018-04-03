package model

import (
	"time"
)

type Passenger struct {
	Id             int64  `orm:"auto"`
	CardId         string `orm:"size(20)"`
	TicketType     int8
	EnterStationId int
	EnterTime      time.Time
	ExitStationId  int
	ExitTime       time.Time
	StmtDay        string
	Flag           int8
}

type MetroEnter struct {
	Id         int64  `orm:"auto"`
	CardId     string `orm:"column(CARDID)"`
	TicketType int    `orm:"column(TICKET_TYPE)"`
	EquipId    string `orm:"column(EQUIP_ID)"`
	DealTime   string `orm:"column(DEAL_TIME)"`
	TransType  int    `orm:"column(TRANS_TYPE)"`
	StmtDay    string `orm:"column(STMT_DAY)"`
	StationId  string `orm:"column(STATION_ID)"`
	Flag       int8   `orm:"column(flag)"`
}

type MetroExit struct {
	Id         int64  `orm:"auto"`
	CardId     string `orm:"column(CARDID)"`
	TicketType int    `orm:"column(TICKET_TYPE)"`
	EquipId    string `orm:"column(EQUIP_ID)"`
	DealTime   string `orm:"column(DEAL_TIME)"`
	TransType  int    `orm:"column(TRANS_TYPE)"`
	StmtDay    string `orm:"column(STMT_DAY)"`
	StationId  string `orm:"column(STATION_ID)"`
}

func (this *MetroEnter) TableName() string {
	return "TBL_METRO_ENTER_ALL_TEST201510"
}

func (this *MetroExit) TableName() string {
	return "TBL_METRO_EXIT_ALL_TEST201510"
}

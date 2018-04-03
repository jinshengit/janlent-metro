package service

import (
	"janlent-metro/dao"
	"janlent-metro/model"
	"strconv"
	"github.com/astaxie/beego/logs"
	"janlent-metro/utils"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type PassengerService struct {
	PassengerDao dao.PassengerDao
}

func (p *PassengerService) CombineDatas() (bool, error) {
	//1. Get last enter id from redis
	client, err := redis.Dial("tcp", "192.168.0.118:6379")
	if err != nil {
		//TODO:
		fmt.Println("Connect to redis error : 192.168.0.118:6379")
		return false, err
	}
	defer client.Close()
	lastId, err := redis.Int64(client.Do("GET", "last_enter_id"))
	if err != nil {
		//TODO:
		fmt.Println("Get last enter id from redis error")
		return false, err
	}
	startId := lastId + 1
	endId := lastId + 1000

	enterDatas, err := p.PassengerDao.QueryEnterDataById(startId, endId)

	if err != nil {
		return false, err
	}

	if enterDatas == nil || len(enterDatas) <= 0 {
		return false, nil
	}

	var passengers []model.Passenger
	var enterIds []int64

	for _, enter := range enterDatas {
		//Record the enter data id list
		enterIds = append(enterIds, enter.Id)
		if enter.StationId == "0000" {
			continue
		}

		exitDatas, err := p.PassengerDao.QueryExitData(enter.CardId, enter.StmtDay)
		if err != nil || exitDatas == nil || len(exitDatas) <= 0 {
			continue
		}
		inTime, _ := strconv.ParseInt(enter.DealTime, 10, 64)
		for _, exit := range exitDatas {
			outTime, _ := strconv.ParseInt(exit.DealTime, 10 , 64)

			if outTime > inTime {
				var passenger model.Passenger
				enterStation := StationMap[enter.StationId]
				exitStation := StationMap[exit.StationId]

				enterTime := utils.ParseStringToTime(enter.DealTime)
				exitTime := utils.ParseStringToTime(exit.DealTime)

				passenger = model.Passenger{
					CardId: enter.CardId,
					TicketType: int8(enter.TicketType),
					EnterStationId: enterStation.StationTrueNo,
					EnterTime: enterTime,
					ExitStationId: exitStation.StationTrueNo,
					ExitTime: exitTime,
					StmtDay: enter.StmtDay,
					Flag: 1,
				}
				passengers = append(passengers, passenger)
				break
			}
		}
	}
	result, err := p.PassengerDao.InsertPassengers(passengers)
	if err != nil {
		//TODO
		logs.Error(err)
		return false, err
	}
	if result {
		fmt.Println("Insert passenger data successfully, save last enter id to redis")
		_, err = client.Do("SET", "last_enter_id", endId)
		if err != nil {
			logs.Error(err)
			fmt.Println("Save last enter id to redis error")
			return false, err
		}
		fmt.Println("Save last id to redis successfully")
	}
	return true, nil
}
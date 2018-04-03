package service

import (
	"janlent-metro/dao"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/logs"
	"janlent-metro/model"
	"encoding/json"
	"time"
	"fmt"
	"strings"
	"strconv"
)

type PassengerPathService struct {
	PassengerPathDao dao.PassengerPathDao
	PassengerDao     dao.PassengerDao
}

func (p *PassengerPathService) GeneratePassengerPath() (bool, error) {
	//Get last passenger id from redis
	client, err := redis.Dial("tcp", dao.RedisAddress)
	if err != nil {
		logs.Error(err)
		return false, err
	}
	defer client.Close()
	lastId, err := redis.Int64(client.Do("GET", "last_passenger_id"))
	if err != nil {
		logs.Error(err)
		return false, err
	}

	startId := lastId + 1
	endId := lastId + 100
	//Get passenger data from database by start id and end id
	passengers, err := p.PassengerDao.QueryPassengers(startId, endId)

	if err != nil {
		logs.Error(err)
		return false, err
	}
	if passengers == nil || len(passengers) <= 0 {
		return false, nil
	}

	var paths []model.PassengerPath

	client.Do("SELECT", 1)

	for _, passenger := range passengers {
		//Get metro path by passenger start station and end station
		startStation := StationIdMap[passenger.EnterStationId]
		endStation := StationIdMap[passenger.ExitStationId]

		key := strings.Replace(startStation.StationName, " ", "", -1) + "-" + strings.Replace(endStation.StationName, " ", "", -1)
		strJson, err := redis.String(client.Do("GET", key))

		if err != nil {
			continue
		}
		var details []model.MetroPathDetail
		if err = json.Unmarshal([]byte(strJson), &details); err != nil {
			logs.Error(err)
			continue
		}
		totalDuration := passenger.ExitTime.Sub(passenger.EnterTime).Seconds()
		averageDuration := int(totalDuration) / len(details)
		temp := averageDuration

		for i := 0; i <= len(details); i++ {
			var path model.PassengerPath
			if i == len(details) {
				lineNum, _ := strconv.Atoi(strings.Replace(details[i - 1].Line, "号线","", -1))
				path = model.PassengerPath{
					CardId: passenger.CardId,
					Station: details[i - 1].ToStation,
					ArriveTime: passenger.ExitTime,
					Duration: averageDuration,
					Direction: details[i - 1].Direction,
					Line: lineNum,
					PathIndex: details[i - 1].PathIndex + 1, //最后一站，index需要加个1
					IsChangePoint: details[i - 1].IsChangePoint,
				}
			} else if i == 0 {
				lineNum, _ := strconv.Atoi(strings.Replace(details[i].Line, "号线","", -1))
				path = model.PassengerPath{
					CardId: passenger.CardId,
					Station: details[i].FromStation,
					NextStation: details[i].ToStation,
					ArriveTime: passenger.EnterTime,
					Duration: averageDuration,
					Direction: details[i].Direction,
					Line: lineNum,
					PathIndex: details[i].PathIndex,
					IsChangePoint: details[i].IsChangePoint,
				}
			} else {
				duration, _ := time.ParseDuration(fmt.Sprintf("%ds", temp))
				lineNum, _ := strconv.Atoi(strings.Replace(details[i].Line, "号线","", -1))
				path = model.PassengerPath{
					CardId: passenger.CardId,
					Station: details[i].FromStation,
					NextStation: details[i].ToStation,
					ArriveTime: passenger.EnterTime.Add(duration),
					Duration: averageDuration,
					Direction: details[i].Direction,
					Line: lineNum,
					PathIndex: details[i].PathIndex,
					IsChangePoint: details[i].IsChangePoint,
				}
				temp += averageDuration
			}
			paths = append(paths, path)
		}
	}

	result, err := p.PassengerPathDao.InsertPassengerPaths(paths)

	if err != nil {
		//TODO:
		logs.Error(err)
		return false, err
	}

	//Update last passenger id to redis
	if result {
		fmt.Println("Insert passenger path successfully, save last passenger id to redis")
		client.Do("SELECT", 0)
		_, err = client.Do("SET", "last_passenger_id", passengers[len(passengers) - 1].Id)

		if err != nil {
			logs.Error(err)
			fmt.Println("Save last passenger id to redis error")
			return false, err
		}
		fmt.Println("Save last passenger id to redis successfully")
	}

	return true, err
}


















































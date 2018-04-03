package service

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"github.com/astaxie/beego/logs"
	"time"
	"janlent-metro/dao"
	"janlent-metro/model"
	"fmt"
)

type MetroTransportService struct {
	PassengerPathDao dao.PassengerPathDao
	MetroTransportDao dao.MetroTransportDao
}

var BaseStartTime time.Time
var BaseEndTime time.Time
var TimeDuration time.Duration

func init() {
	config, err := yaml.ReadFile("settings.yaml")
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	loc := time.Local

	strStartTime, _ := config.Get("start_time")
	strEndTime, _ := config.Get("end_time")
	strTimeDuration, _ := config.Get("duration")
	BaseStartTime, _ = time.ParseInLocation("2006-01-02 15:04:05", "2015-10-01 " + strStartTime, loc)
	BaseEndTime, _ = time.ParseInLocation("2006-01-02 15:04:05", "2015-10-01 " + strEndTime, loc)
	TimeDuration, _ = time.ParseDuration(strTimeDuration)
}

func (m *MetroTransportService) CalculateTransport() {



	startTime := BaseStartTime

	for {
		endTime := startTime.Add(TimeDuration)



		for k, v := range MetroStationMap {
			beginTime := time.Now()
			var list []model.MetroTransportation
			//按照站名和开始结束时间获取每个方向的乘客人数
			passengerCounts := m.PassengerPathDao.CountPassengerByStation(k, startTime, endTime)
			//遍历StationLink数据，按照上下行将数据区分
			for _, link := range v.StationLinks {
				var item = model.MetroTransportation{
					StationName: k,
					NextStation: link.ToStation,
					StartTime: startTime,
					EndTime: endTime,
					Direction: link.Flag,
				}

				for _, pCount := range passengerCounts {
					if pCount.Station == link.FromStation && pCount.NextStation == link.ToStation {
						item.PassengerOnTrainCount = pCount.Count
					}
					if pCount.Direction == item.Direction && pCount.NextStation == "" {
						item.GetOffPassengerCount = pCount.Count
					}
				}
				list = append(list, item)
			}
			ok, err := m.MetroTransportDao.InsertMetroTransportations(list)

			if err != nil {
				logs.Error(err)
				return
			}

			if ok {
				fmt.Println("Insert metro transportation data successfully")
				elapsed := time.Since(beginTime)
				fmt.Println("Cycle used time : ", elapsed)
			}
		}


		startTime = endTime

		if startTime.After(BaseEndTime) {
			break
		}
	}
}

package service

import (
	"janlent-metro/dao"
	"janlent-metro/model"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"time"
)

type MetroPathService struct {
	MetroPathDao dao.MetroPathDao
}

func (m *MetroPathService) GetMetroPathDetails(key string) ([]model.MetroPathDetail) {
	return m.MetroPathDao.GetPathDetails(key)
}

func (m *MetroPathService) SaveMetroPathDetailToRedis() {
	client, _ := redis.Dial("tcp", dao.RedisAddress)
	defer client.Close()
	client.Do("SELECT", 1)
	for i := 1; i <= 59292; i++ {
		startTime := time.Now()

		path := m.MetroPathDao.GetMetroPath(i)
		if path == nil {
			fmt.Println("Get null metro path")
			break
		}

		details := m.MetroPathDao.GetMetroPathDetails(int(path.Id))
		if details == nil || len(details) <= 0 {
			fmt.Println("Get null metro path detail, path id = ", path.Id)
			break
		}
		key := path.FromNode + "-" + path.ToNode
		strJson, _ := json.Marshal(details)

		_, err := client.Do("SET", key, strJson)
		if err != nil {
			logs.Error(err)
		}
		elapsed := time.Since(startTime)
		fmt.Println("One path used : ", elapsed)
	}
}
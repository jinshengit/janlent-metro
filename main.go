package main

import (
	"fmt"
	//"time"
	_ "janlent-metro/dao"
	_ "janlent-metro/service"
	"time"
	"janlent-metro/service"
)

const (
	VERSION = "1.0.0"
	COMPANY = "Janlent SH"
)

func main() {

	fmt.Println("Please select import data: ")
	fmt.Println("1. Combine Metro Passenger Data")
	fmt.Println("2. Generate Passenger Path")
	fmt.Println("3. Combine Metro Passenger With Train")
	fmt.Println("4. Save Metro Path into Redis")
	fmt.Println("5. Just Run Some Testing")

	var inputNum int

	for inputNum != 1 && inputNum != 2 && inputNum != 3 && inputNum != 4 && inputNum != 5 {
		fmt.Scanln(&inputNum)
	}

	fmt.Println("Your input: ", inputNum)

	switch inputNum {
	case 1:
		//fmt.Println("**********Start to combine metro passenger data**********")
		startTime := time.Now()

		var passengerService service.PassengerService = service.PassengerService{}

		for {
			beginTime := time.Now()

			result, err := passengerService.CombineDatas()
			if err != nil {
				fmt.Println(err)
			}
			if !result {
				break
			}
			cycleElapsed := time.Since(beginTime)
			fmt.Println("Cycle used time : ", cycleElapsed)
		}

		elapsed := time.Since(startTime)
		fmt.Println("Total used time : ", elapsed)
	case 2:
		startTime := time.Now()
		var passengerPathService service.PassengerPathService = service.PassengerPathService{}
		for {
			beginTime := time.Now()
			result, err := passengerPathService.GeneratePassengerPath()
			if err != nil {
				fmt.Println(err)
			}
			if !result {
				break
			}
			cycleElapsed := time.Since(beginTime)
			fmt.Println("Cycle used time : ", cycleElapsed)
		}
		elapsed := time.Since(startTime)
		fmt.Println("Total used time : ", elapsed)
	case 4:
		var metroPathService service.MetroPathService = service.MetroPathService{}
		metroPathService.SaveMetroPathDetailToRedis()
	case 5:
		//Do some testing
	}

}
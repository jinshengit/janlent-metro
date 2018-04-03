package utils

import (
	"strconv"
	"time"
)

func ParseStringToTime(str string) time.Time {
	year, _ := strconv.Atoi(str[0:4])
	month, _ := strconv.Atoi(str[4:6])
	day, _ := strconv.Atoi(str[6:8])
	hour, _ := strconv.Atoi(str[8:10])
	minute, _ := strconv.Atoi(str[10:12])
	second, _ := strconv.Atoi(str[12:14])
	return GetDateTime(year, month, day, hour, minute, second)
}

func GetDateTime(year, month, day, hour, minute, second int) time.Time {
	monthEnum := time.January
	switch month {
	case 1:
		monthEnum = time.January
	case 2:
		monthEnum = time.February
	case 3:
		monthEnum = time.March
	case 4:
		monthEnum = time.April
	case 5:
		monthEnum = time.May
	case 6:
		monthEnum = time.June
	case 7:
		monthEnum = time.July
	case 8:
		monthEnum = time.August
	case 9:
		monthEnum = time.September
	case 10:
		monthEnum = time.October
	case 11:
		monthEnum = time.November
	case 12:
		monthEnum = time.December
	}

	loc := time.Local
	return time.Date(year, monthEnum, day, hour, minute, second, 0, loc)
}

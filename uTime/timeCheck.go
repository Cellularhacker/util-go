package uTime

import "time"

func IsNewHour() bool {
	loc, _ := time.LoadLocation("Asia/Seoul")
	_, minutes, _ := time.Now().In(loc).Clock()
	return minutes < 10
}

func IsNewDay() bool {
	loc, _ := time.LoadLocation("Asia/Seoul")
	hours, _, _ := time.Now().In(loc).Clock()
	return hours == 0 && IsNewHour()
}

func IsNewWeek() bool {
	loc, _ := time.LoadLocation("Asia/Seoul")
	weekday := time.Now().In(loc).Weekday()
	return int(weekday) == 0 && IsNewDay() //Sunday
}

func IsNewMonth() bool {
	loc, _ := time.LoadLocation("Asia/Seoul")
	_, _, day := time.Now().In(loc).Date()
	return day == 1 && IsNewDay()
}

func IsNewYear() bool {
	loc, _ := time.LoadLocation("Asia/Seoul")
	_, month, _ := time.Now().In(loc).Date()
	return int(month) == 1 && IsNewMonth()
}

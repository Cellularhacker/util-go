package uTime

import (
	"fmt"
	"github.com/Cellularhacker/util-go"
	"strconv"
	"time"
)

const baseDayOfWeek = 1 //Monday
const baseDay = 1       // 1st
const baseMonth = 1     // January

func GetYesterdayGteLt() (int64, int64) {
	now := time.Now().Unix()
	lt := (now/86400)*86400 - 32400
	gte := lt - 86400

	return gte, lt
}

func GetLastHourGteLt() (int64, int64) {
	now := time.Now().Unix()
	lt := (now / 3600) * 3600
	gte := lt - 3600

	return gte, lt
}

func GetLastWeekGteLt() (int64, int64) {
	now := time.Now().In(util.Loc)
	dayOfWeek := int(now.Weekday())
	diffOfDayOfWeek := baseDayOfWeek - dayOfWeek

	lt := (now.AddDate(0, 0, diffOfDayOfWeek).Unix()/86400)*86400 - 32400
	gte := lt - 604800

	return gte, lt
}

func GetLastMonthGteLt() (int64, int64) {
	now := time.Now().In(util.Loc)
	day := now.Day()
	diffOfDay := baseDay - day

	lt := (now.AddDate(0, 0, diffOfDay).Unix()/86400)*86400 - 32400
	gte := time.Unix(lt, 0).AddDate(0, -1, 0).Unix()

	return gte, lt
}

func GetLastYearGteLt() (int64, int64) {
	now := time.Now().In(util.Loc)
	_, month, day := now.Date()
	diffOfDay := baseDay - day
	diffOfMonth := baseMonth - int(month)

	lt := (now.AddDate(0, diffOfMonth, diffOfDay).Unix()/86400)*86400 - 32400
	gte := time.Unix(lt, 0).AddDate(-1, 0, 0).Unix()

	return gte, lt
}

func GetThisMonthStartGteLt() (int64, int64) {
	now := time.Now().In(util.Loc)
	day := now.Day()
	diffOfDay := baseDay - day

	gte := (now.AddDate(0, 0, diffOfDay).Unix()/86400)*86400 - 32400
	lt := time.Unix(gte, 0).AddDate(0, 1, 0).Unix()

	return gte, lt
}

func GetDateHourMinute10m(ts int64) int {
	t := time.Unix(ts, 0).In(util.Loc)
	year, month, date := t.Date()
	hour, minute, _ := t.Clock()
	minute = (minute / 10) * 10

	str := fmt.Sprintf("%04d%02d%02d%02d%02d", year, int(month), date, hour, minute)
	i, _ := strconv.Atoi(str)

	return i
}

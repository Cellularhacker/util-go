package uTime

import (
	"fmt"
	"github.com/Cellularhacker/util-go"
	"github.com/jinzhu/now"
	"strconv"
	"time"
)

func GetNowDate() int {
	return GetDateFromTs(time.Now().Unix())
}

func GetKST(ts *time.Time) time.Time {
	if ts == nil {
		return time.Now().In(util.Loc)
	}
	return ts.In(util.Loc)
}

func GetKSTDateStrTs(ts int64) string {
	if ts == 0 {
		return ""
	}

	t := time.Unix(ts, 0)
	return GetKSTDateStr(&t)
}

func GetKSTDateStr(t *time.Time) string {
	if t == nil {
		tt := time.Now()
		t = &tt
	}
	year, month, day := t.In(util.Loc).Date()
	hour, minute, second := t.In(util.Loc).Clock()

	return fmt.Sprintf("%04d%02d%02d-%02d%02d%02d", year, int(month), day, hour, minute, second)
}

func GetKSTDateStrBeautifyTs(ts int64) string {
	if ts <= 0 {
		return ""
	}

	t := time.Unix(ts, 0)
	return GetKSTDateStrBeautify(&t)
}

func GetKSTDateStrBeautify(t *time.Time) string {
	if t == nil {
		tt := time.Now()
		t = &tt
	}
	year, month, day := t.In(util.Loc).Date()
	hour, minute, second := t.In(util.Loc).Clock()

	return fmt.Sprintf("%04d-%02d-%02d_%02d:%02d:%02d", year, int(month), day, hour, minute, second)
}

func GetKoreanDateStrBeautifyTs(ts int64) string {
	t := time.Unix(ts, 0)
	year, month, day := t.In(util.Loc).Date()
	hour, minute, second := t.In(util.Loc).Clock()

	return fmt.Sprintf("%04d년 %02d월 %02d %02d시 %02d분 %02d초", year, int(month), day, hour, minute, second)
}

func GetHour() int {
	return GetHourFromTs(time.Now().Unix())
}

func GetHourFromTs(ts int64) int {
	hours, _, _ := time.Unix(ts, 0).In(util.Loc).Clock()

	return hours
}

func GetDateHourFromTs(ts int64) int {
	date := GetDateFromTs(ts)
	hour := GetHourFromTs(ts)
	dateHour, _ := strconv.Atoi(fmt.Sprintf("%08d%02d", date, hour))

	return dateHour
}

func GetDateHourBeautifiedFromTs(ts int64) string {
	t := time.Unix(ts, 0)

	year, month, date := t.In(util.Loc).Date()
	hour, _, _ := t.In(util.Loc).Clock()

	return fmt.Sprintf("%04d-%02d-%02d %02dh KST", year, int(month), date, hour)
}

func GetDateFromTs(ts int64) int {
	year, month, day := time.Unix(ts, 0).In(util.Loc).Date()
	str := fmt.Sprintf("%4d%02d%02d", year, int(month), day)
	i, _ := strconv.Atoi(str)

	return i
}

func GetKSTDateStrByGte(gte int64) string {
	t := time.Unix(gte, 0)
	year, month, day := t.In(util.Loc).Date()

	return fmt.Sprintf("%04d%02d%02d", year, int(month), day)
}

func GetKSTYearMonthFromTs(ts int64) string {
	t := time.Unix(ts, 0)
	year, month, _ := t.In(util.Loc).Date()

	return fmt.Sprintf("%04d%02d", year, int(month))
}

func GetKSTYearWeekFromTs(ts int64) string {
	year, weekNo := GetWeekNoFromTs(ts)

	return fmt.Sprintf("%04d%02d", year, weekNo)
}

func GetYesterdayKSTDate() int64 {
	t := time.Now().Unix() - 86400
	str := GetKSTDateStrByGte(t)
	i, _ := strconv.Atoi(str)

	return int64(i)
}

func GetWeekNoFromTs(ts int64) (int, int) {
	t := time.Unix(ts, 0)

	return t.In(util.Loc).ISOWeek()
}

func GetMonthNoFromTs(ts int64) (int, int) {
	t := time.Unix(ts, 0)
	year, month, _ := t.In(util.Loc).Date()

	return year, int(month)
}

func GetYearNoFromTs(ts int64) int {
	t := time.Unix(ts, 0)
	year, _, _ := t.In(util.Loc).Date()

	return year
}

func GetMonthTsFromTs(ts int64) (int64, int64) {
	t := time.Unix(ts, 0).In(util.Loc)
	eom := now.With(t).EndOfMonth()
	fom := now.With(t).BeginningOfMonth()

	return fom.Unix(), eom.Unix()
}

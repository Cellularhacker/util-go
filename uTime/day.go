package uTime

import (
	"strings"
	"time"
)

const (
	LanguageKorean   = "Korean"
	LanguageEnglish  = "English"
	LanguageJapanese = "Japanese"
)

var (
	WeekdayKorean = []string{
		"일",
		"월",
		"화",
		"수",
		"목",
		"금",
		"토",
	}
	WeekdayEnglish = []string{
		"Sun",
		"Mon",
		"Tues",
		"Wed",
		"Thur",
		"Fri",
		"Sat",
	}
	WeekdayJapanese = []string{
		"日",
		"月",
		"火",
		"水",
		"木",
		"金",
		"土",
	}
)

func FindWeekDayFromString(str string) time.Weekday {
	str = strings.ToLower(str)
	if strings.Contains(str, "일") || strings.Contains(str, "sun") {
		return time.Sunday
	}

	if strings.Contains(str, "월") || strings.Contains(str, "mon") {
		return time.Monday
	}

	if strings.Contains(str, "화") || strings.Contains(str, "tues") {
		return time.Tuesday
	}

	if strings.Contains(str, "수") || strings.Contains(str, "wed") {
		return time.Wednesday
	}

	if strings.Contains(str, "목") || strings.Contains(str, "thur") {
		return time.Thursday
	}

	if strings.Contains(str, "금") || strings.Contains(str, "fri") {
		return time.Friday
	}

	if strings.Contains(str, "토") || strings.Contains(str, "sat") {
		return time.Saturday
	}

	return time.Weekday(-1)
}

func GetWeekDayNameByLanguage(weekday time.Weekday, language string) string {
	if int(weekday) < 0 || int(weekday) > 6 {
		return "Unknown"
	}

	if language == LanguageKorean {
		return WeekdayKorean[weekday]
	}
	if language == LanguageEnglish {
		return WeekdayEnglish[weekday]
	}
	if language == LanguageJapanese {
		return WeekdayJapanese[weekday]
	}

	return "Unknown"
}

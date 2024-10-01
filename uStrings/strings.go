package uStrings

import "strings"

func RemoveAndTrim(original string, targets ...string) string {
	str := original
	for _, tStr := range targets {
		str = strings.ReplaceAll(str, tStr, "")
	}

	return Trim(str)
}

func Trim(str string) string {
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "  ", " ")
	str = strings.ReplaceAll(str, "  ", " ")
	str = strings.ReplaceAll(str, "  ", " ")
	str = strings.ReplaceAll(str, " ", "")

	return str
}

func UnescapeAllEscapingCharacters(str string) string {
	str = strings.ReplaceAll(str, "\\t", "\t")
	str = strings.ReplaceAll(str, "\\a", "\a")
	str = strings.ReplaceAll(str, "\\n", "\n")
	str = strings.ReplaceAll(str, "\\f", "\f")
	str = strings.ReplaceAll(str, "\\r", "\r")
	str = strings.ReplaceAll(str, "\\v", "\v")

	return str
}

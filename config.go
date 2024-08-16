package util

import "time"

type TelegramBotConfig interface {
	SendMessageAt(message string, at ...time.Time)
	SendStarted(hostname string, localIP string, publicIP string)
	SendStopped(hostname string, localIP string, publicIP string)
	SendFailed(location string, err error, at ...time.Time)
}

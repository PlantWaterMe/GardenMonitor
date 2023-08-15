package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Commands = []tgbotapi.BotCommand{
	{
		Command:     "/probe",
		Description: "probe the depth sensor",
	},
}

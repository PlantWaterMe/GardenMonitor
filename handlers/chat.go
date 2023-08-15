package handlers

import (
	"log"
	"strings"

	"github.com/PlantWaterMe/GardenMonitor/sensor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Chat struct {
	bot      *tgbotapi.BotAPI
	sensor   *sensor.Depth
	commands map[string]func(update *tgbotapi.Update)
}

func NewChat(bot *tgbotapi.BotAPI, sensor *sensor.Depth) *Chat {

	return &Chat{
		bot:    bot,
		sensor: sensor,
	}
}

func (c *Chat) Start() *Chat {
	c.commands = c.generateCommands()
	return c
}

func (c *Chat) Handle(update *tgbotapi.Update) {

	command := strings.Split(update.Message.Text, " ")[0]

	if update.Message != nil { // If we got a message
		if command, ok := c.commands[command]; ok { // If the command exists
			command(update)
		}
	}
}

func (c *Chat) generateCommands() map[string]func(update *tgbotapi.Update) {

	commands := make(map[string]func(update *tgbotapi.Update))

	commands["/probe"] = func(update *tgbotapi.Update) {
		level := c.sensor.Probe()

		msgContent := "The water level is normal"

		if level == sensor.Empty {
			msgContent = "The water level is empty"
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgContent)
		msg.ReplyToMessageID = update.Message.MessageID

		c.bot.Send(msg)
	}

	commands["/help"] = func(update *tgbotapi.Update) {

		var helpMessage = "Available commands:\n"
		helpMessage += "/probe returns level of tank. By default returns empty\n"

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpMessage)
		msg.ReplyToMessageID = update.Message.MessageID

		c.bot.Send(msg)
	}

	return commands
}

func (c *Chat) replyWithError(update *tgbotapi.Update, message string) {
	log.Println(message)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ReplyToMessageID = update.Message.MessageID
	c.bot.Send(msg)
}

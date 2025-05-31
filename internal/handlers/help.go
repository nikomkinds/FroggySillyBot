package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func HandleHelpCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {

	helpMessage := "Hello!\nList of commands:\n/help - for help\n/all - mention all fan-club members"

	_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, helpMessage))
	return err
}

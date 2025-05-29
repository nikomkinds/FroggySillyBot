package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func HandleHelpCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	helpMessage := `
Hello!
Available commands:
/help - for help`

	bot.Send(tgbotapi.NewMessage(
		update.Message.Chat.ID,
		helpMessage,
	))
}

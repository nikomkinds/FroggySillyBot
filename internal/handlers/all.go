package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

func HandleAllCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	roman_r := os.Getenv("ROMAN_R")
	roman_k := os.Getenv("ROMAN_K")
	gaziz := os.Getenv("GAZIZ")
	viktor := os.Getenv("VIKTOR")
	dmitriy := os.Getenv("DMITRIY")

	msg := roman_r + "\n" + roman_k + "\n" + gaziz + "\n" + viktor + "\n" + dmitriy

	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))
}

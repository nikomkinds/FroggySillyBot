package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/nikomkinds/FroggySillyBot/internal/handlers"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	token := os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "help":
				if err := handlers.HandleHelpCommand(bot, update); err != nil {
					log.Println(err)
				}
			case "all":
				if err := handlers.HandleAllCommand(bot, update); err != nil {
					log.Println(err)
				}
			default:
				if _, err := bot.Send(tgbotapi.NewMessage(
					update.Message.Chat.ID,
					"Unknown command...",
				)); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

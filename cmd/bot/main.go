package main

import (
	"log"
	"os"

	"github.com/d-darwin/go-bot/internal/app/commands"
	"github.com/d-darwin/go-bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	godotenv "github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TG_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	productService := product.NewService()
	commander := commands.NewCommander(bot, productService)

	for update := range bot.GetUpdatesChan(u) {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			commander.HandleUpdate(update.Message)
		}
	}
}

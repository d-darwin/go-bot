package commands

import (
	"log"

	"github.com/d-darwin/go-bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot           *tgbotapi.BotAPI
	productSerice *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productSerice *product.Service) *Commander {
	return &Commander{
		bot:           bot,
		productSerice: productSerice,
	}
}

func (c *Commander) HandleUpdate(message *tgbotapi.Message) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("Recoverd from panic: %v", panicValue)
		}
	}()

	if message == nil {
		return
	}

	switch message.Command() {
	case "help":
		c.Help(message)
	case "list":
		c.List(message)
	case "get":
		c.Get(message)
	default:
		c.Default(message)
	}
}

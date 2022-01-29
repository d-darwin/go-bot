package commands

import (
	"github.com/d-darwin/go-bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

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
	// command, ok := registeredCommands[message.Command()]
	// if ok {
	// 	command(c, message)
	// } else {
	// 	c.Default(message)
	// }
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

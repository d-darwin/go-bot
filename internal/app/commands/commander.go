package commands

import (
	"encoding/json"
	"fmt"
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

type CommandData struct {
	Command string `json:"command"`
	Offset  int    `json:"offset"`
}

func (c *Commander) HandleUpdate(update *tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("Recoverd from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Parsed data: %+v", parsedData),
		)
		c.bot.Send(msg)
		return
	}

	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}
}

package commands

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("Wrong args:", args)
		return
	}

	product, err := c.productSerice.Get(idx)
	if err != nil {
		log.Println("Cant find product with idx:", idx)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Founded product: %s", product.Title),
	)

	c.bot.Send(msg)
}

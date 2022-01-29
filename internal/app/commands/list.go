package commands

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here are all products: \n \n"
	products := c.productSerice.List()
	for _, p := range products {
		outputMsgText += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CommandData{
		Command: "list",
		Offset:  10,
	})
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData)),
		),
	)
	c.bot.Send(msg)
}

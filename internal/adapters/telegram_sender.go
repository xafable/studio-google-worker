package adapters

import (
	"fmt"
	"strconv"

	"github.com/xafable/studio-google-worker/internal/entities"
	tele "gopkg.in/telebot.v4"
)

type TelegramSender struct {
	Bot *tele.Bot
}

func NewTelegramSender(bot *tele.Bot) *TelegramSender {
	return &TelegramSender{
		Bot: bot,
	}
}

func (s TelegramSender) Send(m entities.SenderMessage) error {

	chatID, err := strconv.ParseInt(m.To, 10, 64)
	if err != nil {
		fmt.Println("error patse chat id: ", err)
		return err
	}

	chat := tele.Chat{
		ID: chatID,
	}
	_, err = s.Bot.Send(&chat, m.Text)
	if err != nil {
		fmt.Println("error when send to telegram chat: ", err)
		return err
	}

	return nil
}

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/xafable/studio-google-worker/internal/birthdays/adapters"
	"github.com/xafable/studio-google-worker/internal/birthdays/entities"
	"github.com/xafable/studio-google-worker/internal/birthdays/usecases"
	"github.com/xafable/studio-google-worker/internal/db"
	"gopkg.in/telebot.v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error when loading godotenv: %s", err)
		os.Exit(0)
	}

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  os.Getenv("TELEGRAM_API_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := db.NewPostgres(os.Getenv("DB_DSN"))
	if err != nil {
		fmt.Println("error when create gorm postgres")
		os.Exit(0)
	}

	gotm := adapters.NewGormBirthdayRepo(db)

	bot.Handle("/start", func(c telebot.Context) error {
		m := c.Message()
		senderID := m.Sender.ID
		name := m.Sender.FirstName
		usecases.HandleMessage(entities.PollerMessage{
			From: strconv.FormatInt(senderID, 10),
			Text: m.Text,
			Name: name,
			Type: entities.ContactTypeTelegram,
		}, gotm)

		return c.Send("Ви підписалися на сповіщення StefhaniStudio")
	})

	bot.Start()
}

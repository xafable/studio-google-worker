package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/xafable/studio-google-worker/internal/adapters"
	"github.com/xafable/studio-google-worker/internal/db"
	"github.com/xafable/studio-google-worker/internal/interfaces"
	"github.com/xafable/studio-google-worker/internal/worker"
	"gopkg.in/telebot.v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error when loading godotenv: %s", err)
		os.Exit(0)
	}

	telebot, err := telebot.NewBot(telebot.Settings{
		Token: os.Getenv("TELEGRAM_API_TOKEN"),
	})
	if err != nil {
		fmt.Println("error create sender: ", err)
		os.Exit(0)
	}
	sender := adapters.NewTelegramSender(telebot)

	db, err := db.NewPostgres(os.Getenv("DB_DSN"))
	if err != nil {
		fmt.Println("error when create gorm postgres")
		os.Exit(0)
	}

	gotm := adapters.NewGormBirthdayRepo(db)
	sheetOneID := os.Getenv("SHEET_ID")
	bjobOne := worker.NewBirthdaysJob("GB", sheetOneID, gotm)
	sendjob := worker.NewSendBirthdaysJob("send_birthdays", sender, gotm)

	workerErrs := worker.Run([]interfaces.Job{
		bjobOne,
		sendjob,
	})
	if workerErrs != nil {
		fmt.Println("errors when run jobs: ", workerErrs)
	}

	os.Exit(0)
}

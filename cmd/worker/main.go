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

func MakeTelegramAdapter() (interfaces.Sender, error) {
	telebot, err := telebot.NewBot(telebot.Settings{
		Token: os.Getenv("TELEGRAM_API_TOKEN"),
	})
	if err != nil {
		fmt.Println("error create sender: ", err)
		return nil, err
	}
	return adapters.NewTelegramSender(telebot), nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error when loading godotenv: %s", err)
		os.Exit(0)
	}

	teleSender, err := MakeTelegramAdapter()
	if err != nil {
		fmt.Println("error when create TelegramAdapter")
		os.Exit(0)
	}

	db, err := db.NewPostgres(os.Getenv("DB_DSN"))
	if err != nil {
		fmt.Println("error when create gorm postgres")
		os.Exit(0)
	}

	gormBirthRepo := adapters.NewGormBirthdayRepo(db)

	sheetKidsID := os.Getenv("SHEET_KIDS_ID")
	sheetAdultsID := os.Getenv("SHEET_ADULTS_ID")

	getBirthdaysJobKids := worker.NewBirthdaysJob("GB", sheetKidsID, gormBirthRepo)
	getBirthdaysJobAdults := worker.NewBirthdaysJob("GB", sheetAdultsID, gormBirthRepo)

	sendJob := worker.NewSendBirthdaysJob("send_birthdays", teleSender, gormBirthRepo)

	workerErrs := worker.Run([]interfaces.Job{
		getBirthdaysJobKids,
		getBirthdaysJobAdults,
		sendJob,
	})

	if workerErrs != nil {
		fmt.Println("errors when run jobs: ", workerErrs)
	}

	os.Exit(0)
}

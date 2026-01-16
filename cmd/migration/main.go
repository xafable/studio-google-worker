package main

import (
	"fmt"
	"os"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/xafable/studio-google-worker/internal/db"
	"github.com/xafable/studio-google-worker/internal/db/migrations"
)

func main() {
	dsn := "postgres://postgres:root@localhost:5432/studio?sslmode=disable"

	db, err := db.NewPostgres(dsn)
	if err != nil {
		fmt.Println("error when create gorm postgres")
		os.Exit(1)
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		migrations.CreateBirthdaysTable(),
	})
	err = m.Migrate()
	if err != nil {
		fmt.Println("error when migrate: ", err)
		os.Exit(1)
	}
}

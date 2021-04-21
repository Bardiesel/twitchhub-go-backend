package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bardiesel/twitchhub-backend-go/models"
	"github.com/bardiesel/twitchhub-backend-go/utils"

	"github.com/bardiesel/twitchhub-backend-go/boot"
)

func main() {
	app := boot.BootApp()

	doMigrate := flag.Bool("migrate", false, "migrate database")
	flag.Parse()
	if *doMigrate {
		db := utils.DatabaseConnection()
		err := db.AutoMigrate(models.Starting{}, models.Emote{},
			models.Literature{}, models.Command{}, models.Extension{})
		utils.ErrorHandler(err, "database not migrated")
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT"))))
}

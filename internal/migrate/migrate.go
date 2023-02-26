package main

import (
	"log"

	"github.com/nkolentcev/fly_services_srv/config"
	"github.com/nkolentcev/fly_services_srv/internal/models"
)

func init() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("не могу прочитать app.env")
	}
	config.ConnectDB(&cfg)
}

func main() {
	config.DB.AutoMigrate(&models.User{})
}

package main

import (
	"fmt"
	"log"

	"github.com/tejashwinn/sependsense/config"
	"github.com/tejashwinn/sependsense/internal/api"
	"github.com/tejashwinn/sependsense/internal/database"
	"github.com/tejashwinn/sependsense/mode"
)

func main() {
	mode.Set(mode.Dev)
	fmt.Println("Starting sependsense in development mode...")
	config := config.Get()
	log.Println(config)
	db, err := database.New(config.Database.Connection)
	if err != nil {
		panic(err)
	}
	g, err := api.New(db, *config)

	g.Run()

}

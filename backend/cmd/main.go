package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naol86/addis-software-starter/project/backend/api/route"
	"github.com/naol86/addis-software-starter/project/backend/config"
)


func main() {
	app, err := config.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	db := app.DB.Database(app.Env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(app.Env.ContextTimeout) * time.Second
	r := gin.Default()
	
	route.SetUpRoute(app.Env, timeout, db, r)
	r.Run(app.Env.ServerAddress)
}
package main

import (
	"github.com/Revazashvili/easer/config"
	"github.com/Revazashvili/easer/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := config.Init(); err != nil{
		log.Fatalf("%s", err.Error())
	}
	app := server.NewApp()
	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}

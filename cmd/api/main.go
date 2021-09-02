package main

import (
	"fmt"
	"github.com/Revazashvili/easer/config"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := config.Init(); err != nil{
		log.Fatalf("%s", err.Error())
	}

	fmt.Println(viper.GetString("port"))
}

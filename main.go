package main

import (
	"fmt"
	"golang_discordbot/bot"
	"golang_discordbot/config"
)

func main() {

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	
	return
}

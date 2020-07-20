package main

import (
	"time"
	"tlbot/bot"
)

func main() {
	bot.GetUpdate()
	bot.SearchMess()

	for {
		bot.GetUpdate()

		bot.Echo()

		time.Sleep(10 * time.Second)

	}

}

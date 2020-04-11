package main

import (
	"github.com/iluxaorlov/echobot/method"
	"log"
)

func main() {
	var token = "YOUR_TOKEN"
	var api = "https://api.telegram.org/bot" + token
	var offset int

	for {
		updateList, err := method.GetUpdates(api, offset)
		if err != nil {
			log.Println(err.Error())
		}

		for _, update := range updateList {
			offset = update.UpdateId + 1

			err := method.SendMessage(api, update)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}
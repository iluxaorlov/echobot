package method

import (
	"bytes"
	"encoding/json"
	"github.com/iluxaorlov/echobot/entity"
	"net/http"
)

type message struct {
	ChatId int `json:"chat_id"`
	Text string `json:"text"`
}

func SendMessage(api string, update entity.Update) error {
	message := message{
		ChatId: update.Message.Chat.Id,
		Text: update.Message.Text,
	}

	js, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = http.Post(api + "/sendMessage", "application/json", bytes.NewBuffer(js))
	if err != nil {
		return err
	}

	return nil
}
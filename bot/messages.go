package bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func sayHello(body *Update) error {
	reqBody := &NewMessage{
		ChatID: body.Message.Chat.ID,
		Text:   "Hello " + body.Message.From.FirstName,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	sendMessage(reqBytes)
	return nil
}

func sendMessage(reqBytes []byte) error {
	res, err := http.Post(createSendMessageURL(), "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}
	return nil
}

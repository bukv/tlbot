package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var tokenAPI string = ""

var lastReadBlock int = 0
var lastUpdateID int

var response Response

type Chat struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Result struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Response struct {
	Result []Result `json:"result"`
}

func check(text string, e error) {
	if e != nil {
		fmt.Printf("%v : %v", text, e)
	}
}

func send(https string) []byte {
	response, err := http.Get(https)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

func getUpdate() {
	var update []byte

	updHTTP := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", tokenAPI)

	update = send(updHTTP)

	err := json.Unmarshal(update, &response)
	check("processing error JSON", err)
}

func sendMessage(chat int, text string) {
	msgHTTP := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s", tokenAPI, chat, text)

	send(msgHTTP)

}

func echo() {
	if lastReadBlock < len(response.Result) {
		sendMessage(response.Result[lastReadBlock].Message.Chat.ID, response.Result[lastReadBlock].Message.Text)
		lastUpdateID = response.Result[lastReadBlock].UpdateID
		botLog()
		lastReadBlock++
	}
}

//write log to file
func botLog() {
	file, err := os.Create("lastupdateid.txt")
	check("file creation error (lastupdateid.txt)", err)
	_, err = file.WriteString(fmt.Sprintf("%d", lastUpdateID))
	check("file writing error (lastupdateid.txt)", err)
}

//search for the last message read
func searchMess() {
	updID, err := ioutil.ReadFile("lastupdateid.txt")
	check("error read lastupdateid.txt", err)
	updInt, err := strconv.Atoi(fmt.Sprintf("%s", updID))
	check("", err)

	lastUpdateID = updInt

	for lastReadBlock = 0; lastReadBlock < len(response.Result); lastReadBlock++ {
		if response.Result[lastReadBlock].UpdateID == lastUpdateID {

			break
		}
	}
}

func main() {
	getUpdate()
	searchMess()

	for {
		getUpdate()

		echo()

		time.Sleep(1 * time.Second)

	}

}

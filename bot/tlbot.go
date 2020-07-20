package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var tokenAPI string = ""

var lastReadBlock int = 0
var lastUpdateID int

var response Response

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

func GetUpdate() {
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

func Echo() {
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
func SearchMess() {
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

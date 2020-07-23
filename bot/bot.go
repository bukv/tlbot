package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	body := &Update{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}
	if !strings.Contains(strings.ToLower(body.Message.Text), "/start") {
		return
	}
	if err := sayHello(body); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}
	fmt.Println("reply sent")
}

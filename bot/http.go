package bot

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.ListenAndServe(":3000", http.HandlerFunc(Handler))
}

func createSendMessageURL() string {
	sendMessageURL := fmt.Sprintf("https://api.telegram.org/%s/sendMessage", TokenAPI())
	return sendMessageURL
}

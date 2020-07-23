#Telegram Bot

Telegram bot in Golang. Currently implemented only `/start` command.
Incoming messages are updated via webhooks. 

Third-party library used:
- [joho/godotenv](http://github.com/joho/godotenv "github.com/joho/godotenv")

***

###How to run

Step 1: Create file `config.env` and write your API-token to it. (example: file config_env_example)

Step 2: Set webhook. To do this, use the command in the terminal:
`curl -F "url=https://<your_server_url>/"  https://api.telegram.org/bot<your_api_token>/setWebhook`

Step 3: Run Telegram Bot.
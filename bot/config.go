package bot

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from config.env into the system
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalln("No config.env file found")
	}
}

func TokenAPI() string {
	return os.Getenv("TOKEN_API")
}

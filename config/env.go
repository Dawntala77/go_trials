package config

import (
	"fmt"
	"log"

	"github.com/subosito/gotenv"
)

func LoadEnv() {
	err := gotenv.Load()
	if err != nil {
		fmt.Println("\tERROR: failed to load .env file\n\n", err)
		log.Fatal(err)
		return
	}
}

package main

import (
	"fmt"

	"github.com/pimentafm/go-api/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	fmt.Println(config.DBDriver)
}

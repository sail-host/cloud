package main

import (
	"fmt"

	"github.com/sail-host/cloud/config"
)

func main() {
	config.LoadConfig()
	fmt.Println("Hello, World!")
	fmt.Println(config.GetConfig())
}

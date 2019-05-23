package main

import (
	"config"
	"time"
	"fmt"
)

func main() {

	config.LoadConfig()

	fmt.Println("Hello, GO!")
	time.Sleep(5*time.Second)
}

package main

import (
	"awesomeProject/dal/redis"
	"awesomeProject/handler"
	"awesomeProject/utils"
	"time"
)

func main() {
	redis.Init()

	// Set up a timer to trigger block scanner in every 30 seconds
	ticker := utils.Schedule(handler.Scan, time.Second*30)

	defer close(ticker)
}

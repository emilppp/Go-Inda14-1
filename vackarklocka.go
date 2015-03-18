package main

import (
	"fmt"
	"time"
)

func Remind(text string, paus time.Duration) {
	tickChannel := time.Tick(paus)
	for {
		select {
		case now := <-tickChannel:
			fmt.Printf(text, now.Format("15:04:05"))
		}
	}
}

func main() {
	go Remind("Klockan är %s Dags att äta\n", 3*time.Second)
	go Remind("Klockan är %s Dags att arbeta\n", 8*time.Second)
	go Remind("Klockan är %s Dags att sova\n", 24*time.Second)
	select {}
}

package main

import (
	"fmt"
	"time"
)

func Remind(text string, paus time.Duration) {
	for {
		a := time.Now().Hour()
		b := time.Now().Minute()
		c := time.Now().Format("15:04")
		if b == 0 {
			if a%3 == 0 {
				fmt.Println(text + c + " Dags att äta")
			}
			if a%8 == 0 {
				fmt.Println(text + c + " Dags att arbeta")
			}
			if a%24 == 0 {
				fmt.Println(text + c + " Dags att sova")
			}

		}
		time.Sleep(paus * time.Minute)
	}
}

func main() {
	Remind("Klockan är nu ", 1)
}

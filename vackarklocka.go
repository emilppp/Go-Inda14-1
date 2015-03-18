package main

import (
	"fmt"
	"time"
)

func Remind(text string, paus time.Duration) {
	

}

func main() {
ticker := time.Ticker(time.Second).C

	go func() {
	for t := range ticker.C {
		fmt.Println("Klockan är", t.Format("13:37"))
	}
	}()
}


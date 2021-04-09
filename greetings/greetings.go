package greetings

import (
	"fmt"
	"time"
)

func GoodDay() {
	fmt.Println("Good Day")
}

func GoodNight() {
	fmt.Println("Good Night")
}

func IsAM() bool {
	h := time.Now().Hour()
	return h < 12
}

func IsAfternoon() bool {
	h := time.Now().Hour()
	return h >= 12 && h < 18
}

func IsEvening() bool {
	h := time.Now().Hour()
	return h >= 18 && h <= 23
}
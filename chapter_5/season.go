package main

import "fmt"

func main()  {
	month := 10
	season := Season(month)
	fmt.Printf("The month: %d is in season: %s", month, season)
}

func Season(month int) string {
	switch {
	case month < 4 && month > 0:
		return "Spring"
	case month > 3 && month < 7:
		return "Summer"
	case month > 6 && month < 10:
		return "Altumn"
	case month > 9 && month <= 12:
		return "Winter"
	default:
		return "Invalid month"
	}
}
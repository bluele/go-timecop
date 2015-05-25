package main

import (
	"fmt"
	"github.com/bluele/go-timecop"
)

func main() {
	t := timecop.Now()
	fmt.Printf("current: %v\n", t)
	fmt.Println("Dive into the future!")
	timecop.Travel(t.AddDate(1, 0, 0))

	for i := 0; i < 3; i++ {
		fmt.Printf("future: %v\n", timecop.Now())
	}

	fmt.Println("Return to the current.")
	timecop.Return()

	fmt.Printf("current: %v\n", timecop.Now())
}

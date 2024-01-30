package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name        string
	Email       string
	dateOfBirth time.Time
}

func main() {
	p1 := Person{
		Email: "Два",
		Name:  "Один",
	}
	p2 := Person{"Один", "Два", time.Date(2000, 12, 1, 0, 0, 0, 0, time.UTC)}

	fmt.Println(p1)
	// fmt.Println(p2)
	fmt.Printf("Man %#v", p2)

}

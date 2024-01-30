package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	// DateOfBirth time.Time
	List []int `json:"list"`
}

func main() {
	p1 := Person{
		Email: "Aлекс",
		Name:  "alex@yandex.ru",
		List:  []int{11},
	}

	json, _ := json.Marshal(p1)

	fmt.Println(string(json))

}

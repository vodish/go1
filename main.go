package main

import (
	"fmt"
	"strings"
)

func main() {
	// bSlice := []byte(" \t\n a lone gopher \n\t\r\n")
	str := `
	
	string  
	
	`
	fmt.Println("[", strings.Trim(str, " \n\r\t"), "]") // a lone gopher
	// fmt.Printf("%s", bSlice)                  // \t\n a lone gopher \n\t\r\n

}

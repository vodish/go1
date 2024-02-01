package main

import "fmt"

// объявление типа
type thisName int

// объявление метода
func (m *thisName) String() string {
	return fmt.Sprintf("MyType: %d", m)
}

func main() {

	var m thisName = 5
	m = 10

	// вызов метода
	s := m.String()
	fmt.Println(m)
	fmt.Println(s)
}

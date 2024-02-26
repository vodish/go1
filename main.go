package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Внешняя функция\nЗапрос: %s", r.URL.Path)
}

func main() {

	http.HandleFunc("/", home)

	http.ListenAndServe(":7010", nil)
}

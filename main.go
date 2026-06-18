package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("program is working")
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/count", CounterHandler)
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/calculate", HandleCalc)
	http.HandleFunc("/agent", HandleAgent)
	http.ListenAndServe(":8080", nil)

}

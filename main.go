package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	str := fmt.Sprintf("hello docker, host:%s ", host)
	w.Write([]byte(str))
}

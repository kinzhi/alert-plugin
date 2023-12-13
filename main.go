// main.go
package main

import (
	"alert-plugin/internal"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", internal.HandleWebhook)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

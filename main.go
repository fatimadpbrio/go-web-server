package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "page.html")
}
func main() {
	http.HandleFunc("/", home)
	fmt.Println("server started on :8081")
	http.ListenAndServe(":8081", nil)
}

package main

import (
	"fmt"
	"net/http"
	"os"
)

// Dynamic port via os.Getenv — the source regex cannot extract a literal here.
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "go-port-env up on :%s\n", port)
	})
	fmt.Println("Listening on :" + port)
	http.ListenAndServe(":"+port, nil)
}

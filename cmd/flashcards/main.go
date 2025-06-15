package main

import (
	"fmt"
	"net/http"

	"github.com/J-V-S-C/flashCards-go/internal/routes"
)

func main() {
	router := routes.NewRouter()
	port := 3333

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)

	if err != nil {
		panic(err)
	}

}

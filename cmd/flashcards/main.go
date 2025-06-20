package main

import (
	"fmt"
	"github.com/joho/godotenv" // add isto
	"log"
	"net/http"

	"github.com/J-V-S-C/flashCards-go/internal/handlers"
	"github.com/J-V-S-C/flashCards-go/internal/repository"
	"github.com/J-V-S-C/flashCards-go/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
	db, err := repository.NewDB()
	if err != nil {
		log.Fatalln(err)
	}
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handler := handlers.NewHandler(services)
	port := 3333

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)
	err = http.ListenAndServe(addr, handler.InitRoutes())

	if err != nil {
		panic(err)
	}

}

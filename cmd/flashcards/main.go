package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/J-V-S-C/flashCards-go/cmd/flashcards/docs"
	_ "github.com/J-V-S-C/flashCards-go/cmd/flashcards/docs"
	"github.com/J-V-S-C/flashCards-go/internal/handlers"
	"github.com/J-V-S-C/flashCards-go/internal/repository"
	"github.com/J-V-S-C/flashCards-go/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	docs.SwaggerInfo.Title = "FlashCards API"
	docs.SwaggerInfo.Description = "API para gerenciamento de decks e flashcards"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3333"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

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

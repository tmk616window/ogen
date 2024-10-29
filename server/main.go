package main

import (
	"fmt"
	"log"
	"net/http"
	"server/config"
	"server/db"
	"server/handler"
	"server/ogen"
	"server/usecase"

	"github.com/rs/cors"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	dbi, err := db.New(cfg.Database)
	if err != nil {
		panic(err)
	}

	u := usecase.NewUsecase(dbi)

	h := handler.NewHandler(u)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	s, err := ogen.NewServer(h)
	if err != nil {
		log.Fatalln(err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), c.Handler(s)); err != nil {
		log.Fatalln(err)
	}
}

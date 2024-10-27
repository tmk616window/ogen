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

	s, err := ogen.NewServer(h)
	if err != nil {
		log.Fatalln(err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), s); err != nil {
		log.Fatalln(err)
	}
}

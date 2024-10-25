package main

import (
	"log"
	"net/http"
	"server/db"
	"server/handler"
	"server/ogen"
	"server/usecase"
)

func main() {
	dbi, err := db.New()
	if err != nil {
		panic(err)
	}

	u := usecase.NewUsecase(dbi)

	h := handler.NewHandler(u)

	s, err := ogen.NewServer(h)
	if err != nil {
		log.Fatalln(err)
	}

	if err := http.ListenAndServe(":8001", s); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"github.com/Sakenzhassulan/user-service/config"
	"github.com/Sakenzhassulan/user-service/db"
	"github.com/Sakenzhassulan/user-service/internal/user"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	dbi, err := db.NewDBCollection(&conf)
	if err != nil {
		log.Fatal(err)
	}
	userClient, err := user.NewServiceClient(&conf, dbi)
	if err != nil {
		log.Fatal(err)
	}

	user.RegisterRoutes(r, userClient)
	if err := http.ListenAndServe(conf.Port, r); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"p2ngc2/config"
	"p2ngc2/handler"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	h := handler.NewHandler(db)

	router := httprouter.New()
	router.GET("/heroes", h.GetHeroes)
	router.GET("/villains", h.GetVillains)
	router.GET("/inventory", h.GetInventory)
	router.GET("/inventory/:id", h.GetItem)
	router.POST("/inventory", h.PostItem)
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}
	fmt.Println("Starting server...")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

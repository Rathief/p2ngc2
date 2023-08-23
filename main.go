package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"p2ngc2/config"
	"p2ngc2/entity"
	"p2ngc2/handler"

	"github.com/julienschmidt/httprouter"
)

var db *sql.DB

func main() {
	ctx := context.Background()
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	h := handler.NewHandler(db)

	iQuery := `
	SELECT * FROM inventory;
	`

	router := httprouter.New()
	router.GET("/heroes", h.GetHeroes)
	router.GET("/villains", h.GetVillains)
	router.GET("/inventory", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		var item entity.Item
		var sliceOfItems []entity.Item
		rows, err := db.QueryContext(ctx, iQuery)
		if err != nil {
			log.Panic(err)
		}
		for rows.Next() {
			rows.Scan(&item.ID, &item.Name, &item.Stock, &item.Description, &item.Status)
			sliceOfItems = append(sliceOfItems, item)
		}
		json.NewEncoder(w).Encode(sliceOfItems)
	})
	router.GET("/inventory/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		var item entity.Item
		var sliceOfItems []entity.Item
		query := `
		SELECT * FROM Inventory WHERE id = ?;
		`
		rows, err := db.QueryContext(ctx, query, p.ByName("id"))
		if err != nil {
			log.Panic(err)
		}
		for rows.Next() {
			rows.Scan(&item.ID, &item.Name, &item.Stock, &item.Description, &item.Status)
			sliceOfItems = append(sliceOfItems, item)
		}
		json.NewEncoder(w).Encode(sliceOfItems)
	})
	router.POST("/inventory", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)
		var i entity.Item
		err := decoder.Decode(&i)
		if err != nil {
			log.Panic(err)
		}
		invPostQuery := `
		INSERT INTO inventory (id, name, stock, description, status)
		VALUES
			(?, ?, ?, ?, ?)
		;
		`
		_, err = db.ExecContext(ctx, invPostQuery, i.ID, i.Name, i.Stock, i.Description, i.Status)
		if err != nil {
			log.Panic(err)
		}
		fmt.Fprintln(w, "Successfully posted.")
	})
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

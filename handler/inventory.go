package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"p2ngc2/entity"

	"github.com/julienschmidt/httprouter"
)

func (h handler) GetInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var item entity.Item
	var sliceOfItems []entity.Item
	rows, err := h.DB.QueryContext(h.BG, `SELECT * FROM Inventory`)
	if err != nil {
		log.Panic(err)
	}
	for rows.Next() {
		rows.Scan(&item.ID, &item.Name, &item.Stock, &item.Description, &item.Status)
		sliceOfItems = append(sliceOfItems, item)
	}
	json.NewEncoder(w).Encode(sliceOfItems)
}

func (h handler) GetItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var item entity.Item
	query := `
	SELECT * FROM Inventory WHERE id = ?;
	`
	row := h.DB.QueryRowContext(h.BG, query, p.ByName("id"))
	err := row.Scan(&item.ID, &item.Name, &item.Stock, &item.Description, &item.Status)
	if err != nil {
		log.Panic(err)
	}
	json.NewEncoder(w).Encode(item)
}

func (h handler) PostItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	_, err = h.DB.ExecContext(h.BG, invPostQuery, i.ID, i.Name, i.Stock, i.Description, i.Status)
	if err != nil {
		log.Panic(err)
	}
	fmt.Fprintln(w, "Successfully posted.")
}

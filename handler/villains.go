package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"p2ngc2/entity"

	"github.com/julienschmidt/httprouter"
)

func (h handler) GetVillains(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var villain entity.Villain
	var villains []entity.Villain
	rows, err := h.DB.QueryContext(h.BG, `
	SELECT * FROM villains;
	`)
	if err != nil {
		log.Panic(err)
	}
	for rows.Next() {
		rows.Scan(&villain.ID, &villain.Name, &villain.Universe, &villain.ImgURL)
		villains = append(villains, villain)
	}
	err = json.NewEncoder(w).Encode(villains)
	if err != nil {
		log.Panic(err)
	}
}

package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h handler) GetVillains(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		id     int
		name   string
		verse  string
		imgURL string
	)
	rows, err := h.DB.QueryContext(h.BG, `
	SELECT * FROM villains;
	`)
	if err != nil {
		log.Panic(err)
	}
	for rows.Next() {
		rows.Scan(&id, &name, &verse, &imgURL)
		fmt.Fprintf(w, "%-3s %-20s %-10s %s\n", fmt.Sprint(id), name, verse, imgURL)
	}
}

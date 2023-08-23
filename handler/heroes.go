package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h handler) GetHeroes(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		id     int
		name   string
		verse  string
		skill  string
		imgURL string
	)
	rows, err := h.DB.QueryContext(h.BG, `
	SELECT * FROM heroes;
	`)
	if err != nil {
		log.Panic(err)
	}
	for rows.Next() {
		rows.Scan(&id, &name, &verse, &skill, &imgURL)
		fmt.Fprintf(w, "%-3s %-20s %-10s %-20s %s\n", fmt.Sprint(id), name, verse, skill, imgURL)
	}
}

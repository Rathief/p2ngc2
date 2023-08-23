package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"p2ngc2/entity"

	"github.com/julienschmidt/httprouter"
)

func (h handler) GetHeroes(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var hero entity.Hero
	var heroes []entity.Hero
	rows, err := h.DB.QueryContext(h.BG, `
	SELECT * FROM heroes;
	`)
	if err != nil {
		log.Panic(err)
	}
	for rows.Next() {
		rows.Scan(&hero.ID, &hero.Name, &hero.Universe, &hero.Skill, &hero.ImgURL)
		heroes = append(heroes, hero)
	}
	err = json.NewEncoder(w).Encode(heroes)
	if err != nil {
		log.Panic(err)
	}
}

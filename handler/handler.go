package handler

import (
	"context"
	"database/sql"
)

type handler struct {
	DB *sql.DB
	BG context.Context
}

func NewHandler(db *sql.DB) handler {
	return handler{db, context.Background()}
}

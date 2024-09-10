package api

import "database/sql"

type APIServer struct {
	add string
	db  *sql.DB
}

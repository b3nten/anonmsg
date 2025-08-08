package core

import (
	"database/sql"

	"benton.codes/anonmsg/cfg"
	"benton.codes/anonmsg/internal/database"
)

type Context struct {
	cfg.Config
	DB      *sql.DB
	Queries *database.Queries
}

func NewContext(config cfg.Config, db *sql.DB, queries *database.Queries) *Context {
	return &Context{
		Config:  config,
		DB:      db,
		Queries: queries,
	}
}

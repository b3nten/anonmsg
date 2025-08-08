package app

import (
	"database/sql"
	"fmt"

	"benton.codes/anonmsg/cfg"
	databaseschema "benton.codes/anonmsg/database/schema"
	"benton.codes/anonmsg/internal/core"
	"benton.codes/anonmsg/internal/database"
	"benton.codes/anonmsg/internal/server"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func Run(config cfg.Config) {
	// connect to database
	db, err := sql.Open("pgx", config.DatabaseURL)
	if err != nil {
		panic(err)
	}
	// run database migrations
	goose.SetBaseFS(databaseschema.FS)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.Up(db, "."); err != nil {
		panic(err)
	}
	// init queries
	q := database.New(db)
	// create core Context
	c := core.NewContext(config, db, q)
	// start server
	fmt.Println("starting server on port", config.Port)
	server.Run(c)
}

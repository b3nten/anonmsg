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
	//**************************************************
	// Establish DB Conn
	//**************************************************

	db, err := sql.Open("pgx", config.DatabaseURL)
	if err != nil {
		panic(err)
	}

	//**************************************************
	// Run Migrations
	//**************************************************

	goose.SetBaseFS(databaseschema.FS)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.Up(db, "."); err != nil {
		panic(err)
	}

	//**************************************************
	// Init
	//**************************************************

	q := database.New(db)
	c := core.NewContext(config, db, q)

	//**************************************************
	// Start Server!
	//**************************************************

	fmt.Println("starting server on port", config.Port)
	server.Run(c)
}

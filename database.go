package main

import (
	"database/sql"
	//"github.com/mattn/go-sqlite3"
)

func (ctx *context)openDatabase() (err error) {
	ctx.db, err = sql.Open("sqlite3", "./foo.db")
	return
}

func (ctx *context)closeDatabase() {
	if ctx.db != nil {
		ctx.db.Close()
	}
}

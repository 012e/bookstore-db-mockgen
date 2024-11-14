package main

import (
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

func InsertPriceTypes(db *sqlx.Tx) {
	b := sqlbuilder.PostgreSQL.NewInsertBuilder().InsertInto("price_types").Cols("name")
	b.
		Values("price").
		Values("profit").
		Values("discount").
		Values("tax")
	sql, args := b.Build()
	db.MustExec(sql, args...)
}

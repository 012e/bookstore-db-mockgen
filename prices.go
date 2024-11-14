package main

import (
	. "github.com/brianvoe/gofakeit/v7"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

func InsertPrice(db *sqlx.Tx, itemCount int) {
	b := sqlbuilder.
		PostgreSQL.
		NewInsertBuilder().
		InsertInto("item_prices").
		Cols("item_id", "price_type", "divider", "value", "begin_date", "end_date")
	for i := range itemCount {
		itemId := i + 1
		price := Number(1000, 10000)
		b.Values(itemId, 1, 1, price, "2021-01-01", "2029-12-31")
		profit := Number(1, 100)
		b.Values(itemId, 2, 100, profit, "2021-01-01", "2029-12-31")
		discount := Number(1, 100)
		b.Values(itemId, 3, 100, discount, "2021-01-01", "2029-12-31")
		tax := Number(1, 100)
		b.Values(itemId, 4, 100, tax, "2021-01-01", "2029-12-31")
	}
	sql, args := b.Build()
	_, err := db.Exec(sql, args...)
	if err != nil {
		panic(err)
	}

}

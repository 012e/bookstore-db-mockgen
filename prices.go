package main

import (
	"math"

	. "github.com/brianvoe/gofakeit/v7"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

func roundDecimalPlaces(x float64) float64 {
	return math.Round(x*100) / 100
}

func InsertPrice(db *sqlx.Tx, itemCount int) {
	b := sqlbuilder.
		PostgreSQL.
		NewInsertBuilder().
		InsertInto("item_prices").
		Cols("item_id", "price_type", "percentage", "begin_date", "end_date", "ordering")
	for i := range itemCount {
		itemId := i + 1

		tax := roundDecimalPlaces(Float64Range(0.05, 0.08))
		b.Values(itemId, "Tax", tax, "2021-01-01", "2029-12-31", 1)

		profit := roundDecimalPlaces(Float64Range(0.2, 0.6))
		b.Values(itemId, "Profit", profit, "2021-01-01", "2029-12-31", 2)

		if Float64() <= 0.4 {
			discount := roundDecimalPlaces(Float64Range(0.05, 0.2))
			b.Values(itemId, "Discount", discount, "2021-01-01", "2029-12-31", 3)
		}
	}
	sql, args := b.Build()
	_, err := db.Exec(sql, args...)
	if err != nil {
		panic(err)
	}

}

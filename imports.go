package main

import (
	. "github.com/brianvoe/gofakeit/v7"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

func InsertImports(db *sqlx.Tx, importsCount int, providerCount int) {
	builder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	builder.InsertInto("imports").Cols("provider_id", "total_cost")
	for _ = range importsCount {
		randomProvider := Number(1, providerCount)
		cost := Float64Range(1_000_000, 100_000_000)
		builder.Values(randomProvider, cost)
	}
	sql, args := builder.Build()
	_, err := db.Exec(sql, args...)
	if err != nil {
		panic(err)
	}
}

func InsertImportItems(db *sqlx.Tx, importItemCount int, importCount int, itemCount int) {
	builder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	builder.InsertInto("import_items").Cols("import_id", "item_id", "quantity")
	for i := range importCount {
		importId := i + 1
		idGenerator := NewUniqueNumGen(1, itemCount)
		for _ = range importItemCount {
			itemId := idGenerator.Get()
			quantity := Number(1, 100)
			builder.Values(importId, itemId, quantity)
		}
	}
	sql, args := builder.Build()
	_, err := db.Exec(sql, args...)
	if err != nil {
		panic(err)
	}
}

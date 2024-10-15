package main

import (
	. "github.com/brianvoe/gofakeit/v7"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

func InsertCustomers(tx *sqlx.Tx, count int) {
	b := sqlbuilder.PostgreSQL.NewInsertBuilder()
	b.InsertInto("customers").
		Cols("first_name", "last_name", "email", "phone_number")
	for range count {
		b.Values(FirstName(), LastName(), Email(), Phone())
	}
	sql, args := b.Build()
	tx.MustExec(sql, args...)
}

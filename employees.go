package main

import (
	. "github.com/brianvoe/gofakeit/v7"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

func InsertEmployees(tx *sqlx.Tx, count int) {
	b := sqlbuilder.PostgreSQL.NewInsertBuilder()
	b.InsertInto("employees").
		Cols("first_name", "last_name", "email", "salary", "profile_picture", "is_manager")
	for range count {
		isManager := Float64Range(0, 1) <= 0.1
		b.Values(FirstName(), LastName(), Email(), Price(1000, 10000), GetRandomImage(), isManager)
	}
	sql, args := b.Build()
	tx.MustExec(sql, args...)
}

package main

import (
	. "github.com/brianvoe/gofakeit/v7"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

func InsertEmployees(tx *sqlx.Tx, count int) {
	b := sqlbuilder.PostgreSQL.NewInsertBuilder()
	b.InsertInto("employees").
		Cols("first_name", "last_name", "email", "salary", "profile_picture", "is_manager", "password")
	b.Values("admin", "admin", "admin@admin.com", 1, GetRandomImage(), false, "admin")
	for range count {
		isManager := Float64Range(0, 1) <= 0.1
		password :=  Password(true, false, false, false, false, 10)
		b.Values(FirstName(), LastName(), Email(), Price(1000, 10000), GetRandomImage(), isManager, password)
	}
	sql, args := b.Build()
	tx.MustExec(sql, args...)
}

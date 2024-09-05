package main

import (
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

var providers = []struct {
	name    string
	address string
}{

	{"Công ty trách nhiệm hữu hạn thếu người", " Linh Xuan, Thủ Đức, Ho Chi Minh City"},
	{"Công ty Anh Kiệt", "Dĩ An, Bình Dương"},
	{"Công ty Minh Thư", "Dĩ An, Bình Dương"},
	{"Đại học Công nghệ Thông tin", "Phnom Penh, Campuchia"},
	{"Trường Đại học Harvard", "Massachusetts Hall, Cambridge, America"},
}

func InsertProviders(db *sqlx.Tx) {
	b := sqlbuilder.PostgreSQL.NewInsertBuilder().
		InsertInto("providers").
		Cols("name", "address")
	for _, provider := range providers {
		b.
			Values(provider.name, provider.address)
	}
	b.SQL("on conflict do nothing")
	sql, args := b.Build()
	db.MustExec(sql, args...)
}

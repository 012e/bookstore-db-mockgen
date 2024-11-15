package main

import (
	. "github.com/brianvoe/gofakeit/v7"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

var tagNames = []string{
	"book",
	"toy",
	"novel",
	"notebook",
	"pen",
	"sticker",
	"fantasy",
	"sci-fi",
	"pencil",
	"porn",
}

func InsertTags(db *sqlx.Tx) {
	b := sqlbuilder.PostgreSQL.NewInsertBuilder().
		InsertInto("tags").
		Cols("name")
	for _, tag := range tagNames {
		b.Values(tag)
	}
	b.SQL("on conflict do nothing")
	sql, args := b.Build()
	db.Exec(sql, args...)
}

func InsertItemsTags(db *sqlx.Tx, itemCount int, tagCount int) {
	b := sqlbuilder.PostgreSQL.NewInsertBuilder().
		InsertInto("items_tags").
		Cols("item_id", "tag_id")
	for i := range itemCount {
		itemId := i + 1
		totalTags := Number(0, 5)
		for _, tagId := range getRandomDistinctIntSlice(totalTags, 1, tagCount) {
			b.Values(itemId, tagId)
		}
	}
	sql, args := b.Build()
	db.Exec(sql, args...)
}

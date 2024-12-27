package main

import (
	_ "embed"
	"math/rand"
	"strings"

	. "github.com/brianvoe/gofakeit/v7"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

//go:embed items.txt
var itemsTxt string

func InsertItems(db *sqlx.Tx, count int) {
	builder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	builder.InsertInto("items").
		Cols("name", "image", "description", "quantity", "provider_id", "base_price")
	for range count {
		loremCount := Number(10, 50)
		quan := Number(1, 100)
		provider := Number(1, 5)
		basePrice := Number(5, 50)
		builder.Values(getRandomItem(), GetRandomImage(), LoremIpsumSentence(loremCount), quan, provider, basePrice)
	}
	sql, args := builder.Build()
	_, err := db.Exec(sql, args...)
	if err != nil {
		panic(err)
	}
}

func getRandomItem() string {
	var items = strings.Split(itemsTxt, "\n")
	return items[rand.Intn(len(items))]
}

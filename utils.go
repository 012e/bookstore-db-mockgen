package main

import (
	_ "embed"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

//go:embed scripts/cleanup.sql
var cleanUpSQL string

//go:embed scripts/schema.sql
var schemaSQL string

func getRandomPostgresImg() interface{} {
	return sqlbuilder.Raw("'porn'")
}

func CleanUp(db *sqlx.DB) {
	db.MustExec(cleanUpSQL)
}

func InitSchema(db *sqlx.DB) {
	db.MustExec(schemaSQL)
}

// getRandomDistinctIntSlice generates a slice of 'count' distinct random integers between 'min' and 'max'
func getRandomDistinctIntSlice(count int, min int, max int) []int {
	if max-min+1 < count {
		// If there aren't enough distinct integers in the range, return an empty slice
		return []int{}
	}

	numSet := make(map[int]struct{}) // To ensure distinct integers
	result := make([]int, 0, count)

	for len(result) < count {
		num := gofakeit.Number(min, max)
		if _, exists := numSet[num]; !exists {
			numSet[num] = struct{}{}
			result = append(result, num)
		}
	}

	return result
}

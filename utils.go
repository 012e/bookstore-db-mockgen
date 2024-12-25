package main

import (
	_ "embed"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jmoiron/sqlx"
)

//go:embed scripts/cleanup.sql
var cleanUpSQL string

//go:embed scripts/schema.sql
var schemaSQL string

//go:embed scripts/startup.sql
var startupSQL string

func CleanUp(db *sqlx.DB) {
	db.MustExec(cleanUpSQL)
}

func InitSchema(db *sqlx.DB) {
	db.MustExec(schemaSQL)
}

func StartUp(tx *sqlx.Tx) {
	tx.MustExec(startupSQL)
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

func GetRandomImage() string {
	return "https://fastly.picsum.photos/id/591/200/200.jpg?hmac=5agpVWsRchY0DObXs23vYWjjgqLZEBhqSvTwfCAcyng"
}

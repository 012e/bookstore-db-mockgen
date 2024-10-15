package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	employeeCount = 30
	itemCount     = 100
	invoiceCount  = 50
	tagCount      = 10
	providerCount = 5
	customerCount = 50
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	CleanUp(db)
	InitSchema(db)

	tx, err := db.Beginx()
	if err != nil {
		panic(err)
	}
	InsertProviders(tx)
	InsertItems(tx, itemCount)
	InsertTags(tx)
	InsertItemsTags(tx, itemCount, tagCount)
	InsertEmployees(tx, employeeCount)
	InsertCustomers(tx, customerCount)
	InsertInvoices(tx, invoiceCount, itemCount, employeeCount, customerCount)

	tx.Commit()
	defer db.Close()
}

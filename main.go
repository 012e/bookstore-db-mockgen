package main

import (
	"log"
	"math/rand"
	"os"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var (
	employeeCount   = 30
	itemCount       = 100
	invoiceCount    = 70
	tagCount        = 10
	providerCount   = 5
	customerCount   = 50
	importCount     = 80
	importItemCount = 100
)

func init() {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetOutput(os.Stdout)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	rand.Seed(696969)
	err := gofakeit.Seed(69696)
	if err != nil {
		panic(err)
	}
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	logrus.Info("Pinging database")
	err = db.Ping()
	if err != nil {
		logrus.Fatal("can't ping database")
	}

	logrus.Info("cleaning up database")
	CleanUp(db)
	logrus.Info("running schema script")
	InitSchema(db)

	tx, err := db.Beginx()
	if err != nil {
		panic(err)
	}
	logrus.Info("inserting providers")
	InsertProviders(tx)
	logrus.Info("inserting items")
	InsertItems(tx, itemCount)
	logrus.Info("inserting tags")
	InsertTags(tx)
	logrus.Info("inserting tags for items")
	InsertItemsTags(tx, itemCount, tagCount)
	logrus.Info("inserting employees")
	InsertEmployees(tx, employeeCount)
	logrus.Info("inserting customers")
	InsertCustomers(tx, customerCount)
	logrus.Info("inserting invoices")
	InsertInvoices(tx, invoiceCount, itemCount, employeeCount, customerCount)
	logrus.Info("inserting imports")
	InsertImports(tx, importCount, providerCount)
	logrus.Info("inserting items for imports")
	InsertImportItems(tx, importItemCount, importCount, itemCount)
	logrus.Info("inserting price types")
	InsertPriceTypes(tx)
	logrus.Info("inserting price")
	InsertPrice(tx, itemCount)
	logrus.Info("Running startup script")
	StartUp(tx)

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

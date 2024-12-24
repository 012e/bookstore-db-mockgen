package main

import (
	"time"

	. "github.com/brianvoe/gofakeit/v7"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

func InsertInvoices(tx *sqlx.Tx, invoiceCount int, itemCount int, employeeCount int, customerCount int) {
	for i := range invoiceCount {
		invoicesBuilder := insertInvoice(i+1, employeeCount, customerCount)
		sql, args := invoicesBuilder.Build()
		tx.MustExec(sql, args...)

		invoicesItemsBuilder := insertInvoiceItem(i+1, itemCount)
		sql, args = invoicesItemsBuilder.Build()
		tx.MustExec(sql, args...)
	}
}

func insertInvoice(invoiceId int, employeeCount int, customerCount int) *sqlbuilder.InsertBuilder {
	invoicesBuilder := sqlbuilder.PostgreSQL.NewInsertBuilder().
		InsertInto("invoices").
		Cols("id", "total", "employee_id", "customer_id", "created_at")
	employeeId := Number(1, employeeCount)
	customerId := Number(1, customerCount)
	createdAt := DateRange(time.Now().AddDate(0, -2, 0), time.Now())
	// TODO: add correct value
	total := Price(50, 200)

	invoicesBuilder.Values(invoiceId, total, employeeId, customerId, createdAt)
	return invoicesBuilder
}

func insertInvoiceItem(invoiceId int, maxItemId int) *sqlbuilder.InsertBuilder {
	// inserting items in the invoice
	invoicesItemsBuilder := sqlbuilder.PostgreSQL.NewInsertBuilder().
		InsertInto("invoices_items").
		Cols("invoice_id", "item_id", "quantity")
	totalDistinctItems := Number(1, 30)
	for _, itemId := range getRandomDistinctIntSlice(totalDistinctItems, 1, maxItemId) {
		quantity := Number(1, 30)
		invoicesItemsBuilder.Values(invoiceId, itemId, quantity)
	}
	return invoicesItemsBuilder
}

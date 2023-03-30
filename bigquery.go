package main

import (
	"cloud.google.com/go/bigquery"
	"context"
	"el_porvenir.com/cloudfunction/model"
	"encoding/base64"
	"fmt"
	"google.golang.org/api/option"
	"os"
)

const (
	TableProduct  = "productSiigo"
	TableInvoice  = "invoiceSiigo"
	TableCustomer = "customerSiigo"
)

// Structs

type bigQueryPorvenir struct {
	dataset bigquery.Dataset
	ctx     context.Context
}

// SiigoElPorvenir retrieves data from Siigo, processes it and stores it in BigQuery.
func (c bigQueryPorvenir) SiigoElPorvenir() {
	c.ProductBigquery()
	c.InvoiceBigquery()
	c.CustomerBigquery()
}

// ProductBigquery creates a new BigQuery table for product data and populates the schema using the model.ProductSiigo{} struct.
//
// Parameters:
// - c: a bigQueryPorvenir instance
//
// Return values:
// - none
func (c bigQueryPorvenir) ProductBigquery() {
	table := c.dataset.Table(TableProduct)
	schema, _ := bigquery.InferSchema(model.ProductSiigo{})
	err := table.Create(c.ctx, &bigquery.TableMetadata{Schema: schema})
	if err != nil {
		fmt.Printf(err.Error())
	}
}

// InvoiceBigquery creates a BigQuery table for invoice data.
// The table will be created in the dataset associated with the current BigQueryPorvenir client.
//
// This function uses the schema inferred from the model.InvoiceSiigo struct to define the table's schema.
//
// If the table already exists, this function will do nothing.
func (c bigQueryPorvenir) InvoiceBigquery() {
	productTable := c.dataset.Table(TableInvoice)
	schema, _ := bigquery.InferSchema(model.InvoiceSiigo{})
	err := productTable.Create(c.ctx, &bigquery.TableMetadata{Schema: schema})
	if err != nil {
		fmt.Printf(err.Error())
	}
}

// CustomerBigquery creates a BigQuery table named "TableCustomer" in the dataset associated with the bigQueryPorvenir instance.
// The table schema is inferred from the model.CustomerSiigo struct.
//
// No parameters are required.
// The function does not return a value, but logs an error message to the console if the table creation fails.
func (c bigQueryPorvenir) CustomerBigquery() {
	table := c.dataset.Table(TableCustomer)
	schema, err := bigquery.InferSchema(model.CustomerSiigo{})
	err = table.Create(c.ctx, &bigquery.TableMetadata{Schema: schema})
	if err != nil {
		fmt.Printf(err.Error())
	}
}

// InsertProduct inserts an array of ProductSiigo objects into the BigQuery table associated with the current bigQueryPorvenir instance.
//
// Parameters:
// - products: an array of ProductSiigo objects to insert into the table.
//
// Returns:
// - none
func (c bigQueryPorvenir) InsertProduct(products []model.ProductSiigo) {
	table := c.dataset.Table(TableProduct)
	u := table.Inserter()
	if err := u.Put(c.ctx, products); err != nil {
		fmt.Printf(err.Error())
	}
}

// InserCustomer inserts an array of CustomerSiigo models into the BigQuery table associated with this client.
//
// Parameters:
// - customers: the array of CustomerSiigo models to insert
//
// Return values: none
func (c bigQueryPorvenir) InserCustomer(customers []model.CustomerSiigo) {
	table := c.dataset.Table(TableCustomer)
	u := table.Inserter()
	if err := u.Put(c.ctx, customers); err != nil {
		fmt.Printf(err.Error())
	}
}

// InserInvoice inserts a slice of Siigo invoices into the BigQuery table for invoices.
//
// Parameters:
// - invoices: a slice of Siigo invoice models to be inserted
//
// Return values: none
func (c bigQueryPorvenir) InserInvoice(invoices []model.InvoiceSiigo) {
	table := c.dataset.Table(TableInvoice)
	u := table.Inserter()
	if err := u.Put(c.ctx, invoices); err != nil {
		fmt.Printf(err.Error())
	}
}

// Factory functions

// NewBigQueryPorvenir creates a new BigQuery client for the specified project and dataset.
// It uses the credentials provided by the GC_KEY environment variable for authentication.
//
// Parameters:
// - ctx: the context of the function call
//
// Return value:
// - a pointer to the newly created BigQuery client
func NewBigQueryPorvenir(ctx context.Context) *bigQueryPorvenir {
	jsonBytes, _ := base64.StdEncoding.DecodeString(os.Getenv("GC_KEY"))
	client, _ := bigquery.NewClient(ctx, os.Getenv("GC_PROJECT"), option.WithCredentialsJSON(jsonBytes))
	dataset := client.Dataset(os.Getenv("GC_DATESET"))
	return &bigQueryPorvenir{
		dataset: *dataset,
		ctx:     ctx,
	}
}

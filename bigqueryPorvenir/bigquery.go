package bigqueryPorvenir

import (
	"cloud.google.com/go/bigquery"
	"context"
	invoice_repository "el_porvenir.com/cloudfunction/bigqueryPorvenir/repository/invoice"
	bigquery_invoice "el_porvenir.com/cloudfunction/bigqueryPorvenir/repository/invoice/model"
	util_bigquery "el_porvenir.com/cloudfunction/bigqueryPorvenir/util"
	siigo_model "el_porvenir.com/cloudfunction/siigo/model"
	"encoding/base64"
	"fmt"
	"google.golang.org/api/option"
	"os"
)

// Structs

type BigQueryPorvenir struct {
	dataset bigquery.Dataset
	ctx     context.Context
	client  *bigquery.Client
}

// SiigoElPorvenir retrieves data from Siigo, processes it and stores it in BigQuery.
func (c BigQueryPorvenir) SiigoElPorvenir(siggoData siigo_model.SiigoData) {
	invoice_repository.InsertInvoice(c.ctx, c.client, c.dataset, siggoData)
}

func (c BigQueryPorvenir) CreateTable() {
	table := c.dataset.Table(invoice_repository.Invoice)
	table.Delete(c.ctx).Error()
	schema, err := bigquery.InferSchema(bigquery_invoice.InvoiceSiigo{})
	if err != nil {
		fmt.Printf(err.Error())
	}
	util_bigquery.CreateTable(c.ctx, table, schema)
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
func NewBigQueryPorvenir(ctx context.Context) *BigQueryPorvenir {
	jsonBytes, _ := base64.StdEncoding.DecodeString(os.Getenv("GC_KEY"))
	client, _ := bigquery.NewClient(ctx, os.Getenv("GC_PROJECT"), option.WithCredentialsJSON(jsonBytes))
	dataset := client.Dataset(os.Getenv("GC_DATESET"))
	return &BigQueryPorvenir{
		client:  client,
		dataset: *dataset,
		ctx:     ctx,
	}
}

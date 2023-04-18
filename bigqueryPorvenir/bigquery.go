package bigqueryPorvenir

import (
	"cloud.google.com/go/bigquery"
	"context"
	invoice_repository "el_porvenir.com/cloudfunction/bigqueryPorvenir/repository/invoice"
	siigo_model "el_porvenir.com/cloudfunction/siigo/model"
	"encoding/base64"
	"google.golang.org/api/option"
	"os"
)

// Structs

type BigQueryPorvenir struct {
	dataset   bigquery.Dataset
	ctx       context.Context
	siggoData siigo_model.SiigoData
}

// SiigoElPorvenir retrieves data from Siigo, processes it and stores it in BigQuery.
func (c BigQueryPorvenir) SiigoElPorvenir() {
	invoice_repository.InsertInvoice(c.ctx, c.dataset, c.siggoData)
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
func NewBigQueryPorvenir(ctx context.Context, siggoData siigo_model.SiigoData) *BigQueryPorvenir {
	jsonBytes, _ := base64.StdEncoding.DecodeString(os.Getenv("GC_KEY"))
	client, _ := bigquery.NewClient(ctx, os.Getenv("GC_PROJECT"), option.WithCredentialsJSON(jsonBytes))
	dataset := client.Dataset(os.Getenv("GC_DATESET"))
	return &BigQueryPorvenir{
		dataset:   *dataset,
		ctx:       ctx,
		siggoData: siggoData,
	}
}

package util_bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"google.golang.org/api/googleapi"
	"net/http"
)

// CreateTable creates a new table in BigQuery with the given schema if it does not exist.
//
// Parameters:
// - table: a pointer to the BigQuery table to create
// - schema: the schema for the table
func CreateTable(context context.Context, table *bigquery.Table, schema bigquery.Schema) {
	if _, err := table.Metadata(context); err != nil {
		if e, ok := err.(*googleapi.Error); ok {
			if e.Code == http.StatusNotFound {
				err := table.Create(context, &bigquery.TableMetadata{Schema: schema})
				if err != nil {
					fmt.Printf(err.Error())
				}
			}
		}
	}
}

func SetStringToEmptyIfNil(strPtr *string) string {
	if strPtr == nil {
		// If the string pointer is nil, create a new empty string and assign it to the pointer
		return ""
	}
	return *strPtr
}

func ConvertUint64PointerToUint64(uintPtr *uint64) uint64 {
	if uintPtr != nil {
		// If the uint64 pointer is not nil, dereference it and return the value
		return *uintPtr
	}
	// If the uint64 pointer is nil, return a default value (e.g., 0)
	return 0
}

func ConvertUint64PointerToInt64(uintPtr *int64) int64 {
	if uintPtr != nil {
		// If the uint64 pointer is not nil, dereference it and return the value
		return *uintPtr
	}
	// If the uint64 pointer is nil, return a default value (e.g., 0)
	return 0
}

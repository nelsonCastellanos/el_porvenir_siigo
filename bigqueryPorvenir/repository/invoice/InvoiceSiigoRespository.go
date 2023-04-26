package invoice_repository

import (
	"cloud.google.com/go/bigquery"
	"context"
	bigquery_customer "el_porvenir.com/cloudfunction/bigqueryPorvenir/repository/customer/model"
	bigquery_invoice "el_porvenir.com/cloudfunction/bigqueryPorvenir/repository/invoice/model"
	bigquery_product "el_porvenir.com/cloudfunction/bigqueryPorvenir/repository/product/model"
	util_bigquery "el_porvenir.com/cloudfunction/bigqueryPorvenir/util"
	siigo_model "el_porvenir.com/cloudfunction/siigo/model"
	siigo_customer "el_porvenir.com/cloudfunction/siigo/model/siigo/customer"
	siigo_invoice "el_porvenir.com/cloudfunction/siigo/model/siigo/invoice"
	siigo_product "el_porvenir.com/cloudfunction/siigo/model/siigo/product"
	"fmt"
	"strings"
)

const (
	Invoice = "invoice_porvenir"
)

func InsertInvoice(ctx context.Context, dataset bigquery.Dataset, listsSiigo siigo_model.SiigoData) {
	var listBigQuery []bigquery_invoice.InvoiceSiigo
	for _, item := range listsSiigo.Invoices {
		invoice := bigquery_invoice.InvoiceSiigo{
			ID:                 item.ID,
			Name:               item.Name,
			Date:               item.Date,
			CustomerId:         item.Customer.ID,
			CustomerName:       getCustomerName(item, listsSiigo),
			CustomerIdTypeCode: getCustomerIdTypeCode(item, listsSiigo),
			CustomerIdTypeName: getCustomerNameIdTypeName(item, listsSiigo),
			Total:              item.Total,
			StampStatus:        item.Stamp.Status,
			StampCufe:          item.Stamp.Cufe,
			MailStatus:         item.Mail.Status,
			MailObservations:   item.Mail.Observations,
			Created:            item.Metadata.Created,
			LastUpdated:        util_bigquery.SetStringToEmptyIfNil(item.Metadata.LastUpdated),
		}
		for _, payment := range item.Payments {
			invoice.Payments = append(invoice.Payments, bigquery_invoice.Payment{
				ID:    payment.ID,
				Name:  payment.Name,
				Value: payment.Value,
			})
		}
		for _, itemInvoice := range item.Items {
			invoice.Items = append(invoice.Items, bigquery_invoice.Item{
				Product:  getProduct(itemInvoice, listsSiigo),
				Quantity: itemInvoice.Quantity,
				Price:    itemInvoice.Price,
				Total:    itemInvoice.Total,
			})
		}
		listBigQuery = append(listBigQuery, invoice)
	}
	table := dataset.Table(Invoice)
	schema, err := bigquery.InferSchema(bigquery_invoice.InvoiceSiigo{})
	if err != nil {
		fmt.Printf(err.Error())
	}
	util_bigquery.CreateTable(ctx, table, schema)
	u := table.Inserter()
	if err := u.Put(ctx, listBigQuery); err != nil {
		fmt.Printf(err.Error())
	}
}

func getCustomerNameIdTypeName(item siigo_invoice.InvoiceSiigo, siigo siigo_model.SiigoData) string {
	for _, typesSiigo := range siigo.Customers {
		if typesSiigo.ID == item.Customer.ID {
			return typesSiigo.IDType.Name
		}
	}
	return ""
}

func getCustomerIdTypeCode(item siigo_invoice.InvoiceSiigo, siigo siigo_model.SiigoData) string {
	for _, typesSiigo := range siigo.Customers {
		if typesSiigo.ID == item.Customer.ID {
			return typesSiigo.IDType.Code
		}
	}
	return ""
}

func getCustomerName(item siigo_invoice.InvoiceSiigo, siigo siigo_model.SiigoData) string {
	for _, typesSiigo := range siigo.Customers {
		if typesSiigo.ID == item.Customer.ID {
			return strings.Join(*typesSiigo.Name, " ")
		}
	}
	return ""
}

func getProduct(invoice siigo_invoice.Item, listsSiigo siigo_model.SiigoData) bigquery_product.ProductSiigo {
	var productBigQuery bigquery_product.ProductSiigo
	for _, product := range listsSiigo.Products {
		if invoice.Code == product.Code || invoice.ID == product.ID {
			account := getAccountGroup(product.AccountGroup, listsSiigo.AccountGroup)
			return bigquery_product.ProductSiigo{
				Code:             product.Code,
				Name:             product.Name,
				AccountGroupID:   account.ID,
				AccountGroupName: account.Name,
			}
		}
	}
	return productBigQuery
}

func getAccountGroup(group siigo_product.AccountGroup, accountGroups []siigo_customer.AccountGroupSiigo) bigquery_customer.AccountGroupSiigo {
	for _, accountGroup := range accountGroups {
		if group.ID == accountGroup.ID {
			return bigquery_customer.AccountGroupSiigo(accountGroup)
		}
	}
	return bigquery_customer.AccountGroupSiigo{}
}

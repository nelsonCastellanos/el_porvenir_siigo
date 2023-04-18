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
	Invoice = "Invoice"
)

func InsertInvoice(ctx context.Context, dataset bigquery.Dataset, listsSiigo siigo_model.SiigoData) {
	var listBigQuery []bigquery_invoice.InvoiceSiigo
	for _, item := range listsSiigo.Invoices {
		invoice := bigquery_invoice.InvoiceSiigo{
			ID:       item.ID,
			Document: getDocument(item, listsSiigo),
			Prefix:   item.Prefix,
			Number:   item.Number,
			Name:     item.Name,
			Date:     item.Date,
			Customer: getCustomer(item, listsSiigo),
			Seller:   item.Seller,
			Total:    item.Total,
			Balance:  item.Balance,
			Stamp:    bigquery_invoice.Stamp(item.Stamp),
			Mail:     bigquery_invoice.Mail(item.Mail),
			Metadata: bigquery_invoice.Metadata{
				Created:     item.Metadata.Created,
				LastUpdated: util_bigquery.SetStringToEmptyIfNil(item.Metadata.LastUpdated),
			},
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
				Product:     getProduct(itemInvoice, listsSiigo),
				Quantity:    itemInvoice.Quantity,
				Price:       itemInvoice.Price,
				Description: itemInvoice.Description,
				Taxes:       getTaxes(itemInvoice.Taxes),
				Total:       itemInvoice.Total,
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

func getProduct(invoice siigo_invoice.Item, listsSiigo siigo_model.SiigoData) bigquery_product.ProductSiigo {
	var productBigQuery bigquery_product.ProductSiigo
	for _, product := range listsSiigo.Products {
		if invoice.Code == product.Code || invoice.ID == product.ID {
			return bigquery_product.ProductSiigo{
				ID:                  product.ID,
				Code:                product.Code,
				Name:                product.Name,
				AccountGroup:        getAccountGroup(product.AccountGroup, listsSiigo.AccountGroup),
				Type:                product.Type,
				StockControl:        product.StockControl,
				Active:              product.Active,
				TaxClassification:   product.TaxClassification,
				TaxIncluded:         product.TaxIncluded,
				TaxConsumptionValue: util_bigquery.ConvertUint64PointerToInt64(product.TaxConsumptionValue),
				Unit:                bigquery_product.Unit(product.Unit),
				UnitLabel:           util_bigquery.SetStringToEmptyIfNil(product.UnitLabel),
				Reference:           product.Reference,
				Description:         product.Description,
				AdditionalFields: bigquery_product.AdditionalFields{
					Barcode: product.AdditionalFields.Barcode,
				},
				AvailableQuantity: product.AvailableQuantity,
				Metadata: bigquery_product.Metadata{
					Created:     product.Metadata.Created,
					LastUpdated: util_bigquery.SetStringToEmptyIfNil(product.Metadata.LastUpdated),
				},
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

func getTaxes(taxes []siigo_invoice.Tax) []bigquery_invoice.Tax {
	var taxesBigquery []bigquery_invoice.Tax
	for _, tax := range taxes {
		taxes = append(taxes, tax)
	}
	return taxesBigquery
}

func getCustomer(item siigo_invoice.InvoiceSiigo, siigo siigo_model.SiigoData) bigquery_customer.CustomerSiigo {
	customer := bigquery_customer.CustomerSiigo{}
	for _, typesSiigo := range siigo.Customers {
		if typesSiigo.ID == item.Customer.ID {
			customer = bigquery_customer.CustomerSiigo{
				ID:             typesSiigo.ID,
				Type:           typesSiigo.Type,
				PersonType:     typesSiigo.PersonType,
				IDType:         bigquery_customer.IDType(typesSiigo.IDType),
				Identification: typesSiigo.Identification,
				BranchOffice:   typesSiigo.BranchOffice,
				CheckDigit:     util_bigquery.SetStringToEmptyIfNil(typesSiigo.CheckDigit),
				Name:           strings.Join(*typesSiigo.Name, " "),
				Active:         typesSiigo.Active,
				VatResponsible: typesSiigo.VatResponsible,
				Address:        bigquery_customer.Address(typesSiigo.Address),
				CommercialName: util_bigquery.SetStringToEmptyIfNil(typesSiigo.CommercialName),
				Metadata: bigquery_customer.Metadata{
					Created:     typesSiigo.Metadata.Created,
					LastUpdated: util_bigquery.SetStringToEmptyIfNil(typesSiigo.Metadata.LastUpdated),
				},
			}
			for _, responsibility := range typesSiigo.FiscalResponsibilities {
				customer.FiscalResponsibilities = append(customer.FiscalResponsibilities,
					bigquery_customer.FiscalResponsibility(responsibility),
				)
			}
			for _, phone := range typesSiigo.Phones {
				customer.Phones = append(customer.Phones,
					bigquery_customer.PhoneElement{
						Indicative: util_bigquery.SetStringToEmptyIfNil(phone.Indicative),
						Number:     phone.Number,
						Extension:  util_bigquery.SetStringToEmptyIfNil(phone.Extension),
					},
				)
			}
			for _, contact := range typesSiigo.Contacts {
				contactSiigo := bigquery_customer.Contact{
					FirstName: contact.FirstName,
					LastName:  contact.LastName,
					Email:     contact.Email,
				}
				if contact.Phone != nil {
					contactSiigo.Phone.Indicative = util_bigquery.SetStringToEmptyIfNil(contact.Phone.Indicative)
					contactSiigo.Phone.Number = util_bigquery.SetStringToEmptyIfNil(contact.Phone.Number)
				} else {
					contactSiigo.Phone = bigquery_customer.ContactPhone{
						Indicative: "",
						Number:     "",
					}
				}
				customer.Contacts = append(customer.Contacts)
			}
			return customer
		}

	}
	return customer
}

func getDocument(item siigo_invoice.InvoiceSiigo, siigo siigo_model.SiigoData) bigquery_invoice.DocumentTypesSiigo {
	document := bigquery_invoice.DocumentTypesSiigo{}
	for _, typesSiigo := range siigo.DocumentType {
		if typesSiigo.ID == item.Document.ID {
			return bigquery_invoice.DocumentTypesSiigo(typesSiigo)
		}
	}
	return document
}

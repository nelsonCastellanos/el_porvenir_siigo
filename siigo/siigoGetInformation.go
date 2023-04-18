package siigo

import (
	siigo_model "el_porvenir.com/cloudfunction/siigo/model"
	"el_porvenir.com/cloudfunction/siigo/model/auth"
	siigo_customer "el_porvenir.com/cloudfunction/siigo/model/siigo/customer"
	siigo_invoice "el_porvenir.com/cloudfunction/siigo/model/siigo/invoice"
	siigo_product "el_porvenir.com/cloudfunction/siigo/model/siigo/product"
	siigo_request "el_porvenir.com/cloudfunction/siigo/request"
	"fmt"
)

const (
	PageSize      = "100"
	PRODUCT       = "/v1/products?page_size=" + PageSize + "&page=1"
	CUSTOMER      = "/v1/customers?page_size=" + PageSize + "&page=1"
	INVOICE       = "/v1/invoices?page_size=" + PageSize + "&page=1"
	AccountGroup  = "/v1/account-groups"
	DocumentTypes = "/v1/document-types?type=FV"
	PaymentsTypes = "/v1/payment-types?document_type=FV"
	Taxes         = "/v1/taxes"
)

func GetSiigoData(token auth.TokenResult) siigo_model.SiigoData {
	taxes := GetTaxes(token)
	paymentsTypes := GetPaymentsTypes(token)
	documentType := GetDocumentType(token)
	accountGroup := GetAccountGroup(token)
	products := GetProducts(token)
	customers := GetCustomer(token)
	invoices := GetInvoice(token)
	return siigo_model.SiigoData{
		Taxes:         taxes,
		PaymentsTypes: paymentsTypes,
		DocumentType:  documentType,
		AccountGroup:  accountGroup,
		Products:      products,
		Customers:     customers,
		Invoices:      invoices,
	}
}

func GetTaxes(token auth.TokenResult) []siigo_invoice.TaxesSiigo {
	var taxes []siigo_invoice.TaxesSiigo
	requestPaginate := siigo_request.NewSimpleRequest(siigo_invoice.TaxesSiigo{}, Taxes)
	allItems := requestPaginate.GetItems(token)
	for _, item := range allItems {
		tax, ok := item.(*siigo_invoice.TaxesSiigo)
		if ok {
			taxes = append(taxes, *tax)
		} else {
			fmt.Println("Error: el item no es de tipo siigo_invoice.TaxesSiigo")
		}
	}
	return taxes
}

func GetPaymentsTypes(token auth.TokenResult) []siigo_invoice.PaymentTypesSiigo {
	var paymentsTypes []siigo_invoice.PaymentTypesSiigo
	requestPaginate := siigo_request.NewSimpleRequest(siigo_invoice.PaymentTypesSiigo{}, PaymentsTypes)
	allItems := requestPaginate.GetItems(token)
	for _, item := range allItems {
		paymentType, ok := item.(*siigo_invoice.PaymentTypesSiigo)
		if ok {
			paymentsTypes = append(paymentsTypes, *paymentType)
		} else {
			fmt.Println("Error: el item no es de tipo siigo_invoice.PaymentTypesSiigo")
		}
	}
	return paymentsTypes
}

func GetDocumentType(token auth.TokenResult) []siigo_invoice.DocumentTypesSiigo {
	var documentTypes []siigo_invoice.DocumentTypesSiigo
	requestPaginate := siigo_request.NewSimpleRequest(siigo_invoice.DocumentTypesSiigo{}, DocumentTypes)
	allItems := requestPaginate.GetItems(token)
	for _, item := range allItems {
		documentType, ok := item.(*siigo_invoice.DocumentTypesSiigo)
		if ok {
			documentTypes = append(documentTypes, *documentType)
		} else {
			fmt.Println("Error: el item no es de tipo siigo_invoice.DocumentTypesSiigo")
		}
	}
	return documentTypes
}

func GetAccountGroup(token auth.TokenResult) []siigo_customer.AccountGroupSiigo {
	var accountGroups []siigo_customer.AccountGroupSiigo
	requestPaginate := siigo_request.NewSimpleRequest(siigo_customer.AccountGroupSiigo{}, AccountGroup)
	allItems := requestPaginate.GetItems(token)
	for _, item := range allItems {
		accountGroup, ok := item.(*siigo_customer.AccountGroupSiigo)
		if ok {
			accountGroups = append(accountGroups, *accountGroup)
		} else {
			fmt.Println("Error: el item no es de tipo siigo_product.ProductSiigo")
		}
	}
	return accountGroups
}

// GetProducts retrieves all products from Siigo using the provided access token.
// It paginates through all available results and returns them as an array of ProductSiigo objects.
//
// Parameters:
// - token: the access token to use for authentication with Siigo
//
// Returns:
// - an array of ProductSiigo objects representing all products in Siigo
func GetProducts(token auth.TokenResult) []siigo_product.ProductSiigo {
	var products []siigo_product.ProductSiigo
	requestPaginate := siigo_request.NewRequestPaginate(siigo_product.ProductSiigo{}, PRODUCT)
	allItems := requestPaginate.GetAllItems(token)
	for _, item := range allItems {
		product, ok := item.(*siigo_product.ProductSiigo)
		if ok {
			products = append(products, *product)
		} else {
			fmt.Println("Error: el item no es de tipo siigo_product.ProductSiigo")
		}
	}
	return products
}

// GetInvoice retrieves all invoice data from Siigo API using the given access token.
//
// Parameters:
// - token: the authentication token needed to access the Siigo API.
//
// Returns:
// - a slice of InvoiceSiigo containing all the invoice data retrieved from Siigo API.
func GetInvoice(token auth.TokenResult) []siigo_invoice.InvoiceSiigo {
	var invoices []siigo_invoice.InvoiceSiigo
	requestPaginate := siigo_request.NewRequestPaginate(siigo_invoice.InvoiceSiigo{}, INVOICE)
	allItems := requestPaginate.GetAllItems(token)
	for _, item := range allItems {
		invoice, ok := item.(*siigo_invoice.InvoiceSiigo)
		if ok {
			invoices = append(invoices, *invoice)
		} else {
			fmt.Println("Error: el item no es de tipo siigo_product.InvoiceSiigo")
		}
	}
	return invoices
}

// GetCustomer retrieves all customers from the Siigo API using the provided token.
// Returns an array of siigo.CustomerSiigo structs containing the customer information.
//
// Parameters:
// - token: a siigo.TokenResult struct containing the access token for authenticating with the Siigo API
//
// Returns:
// - an array of siigo.CustomerSiigo structs representing all customers returned by the Siigo API
func GetCustomer(token auth.TokenResult) []siigo_customer.CustomerSiigo {
	var customers []siigo_customer.CustomerSiigo
	requestPaginate := siigo_request.NewRequestPaginate(siigo_customer.CustomerSiigo{}, CUSTOMER)
	allItems := requestPaginate.GetAllItems(token)
	for _, item := range allItems {
		customer, ok := item.(*siigo_customer.CustomerSiigo)
		if ok {
			customers = append(customers, *customer)
		} else {
			fmt.Println("Error: el item no es de tipo siigo_product.CustomerSiigo")
		}
	}
	return customers
}

package p

import (
	"el_porvenir.com/cloudfunction/model"
	"encoding/json"
)

// getInvoice retrieves all invoice data from Siigo API using the given access token.
//
// Parameters:
// - token: the authentication token needed to access the Siigo API.
//
// Returns:
// - a slice of InvoiceSiigo containing all the invoice data retrieved from Siigo API.
func getInvoice(token model.TokenResult) []model.InvoiceSiigo {
	var allProducts []model.InvoiceSiigo
	responseSiigo := new(model.ResponseSiigo)
	err := getWithBearerToken(INVOICE, token.AccessToken, responseSiigo)
	if err != nil {
		panic("Error to get value auth from siigo")
	}
	allProducts = castInvoice(responseSiigo.Results, allProducts)
	next := responseSiigo.LinksV2.Next != nil
	for next {
		path := getPath(responseSiigo.LinksV2.Next.Href)
		responseSiigo = new(model.ResponseSiigo)
		getWithBearerToken(path, token.AccessToken, responseSiigo)
		allProducts = castInvoice(responseSiigo.Results, allProducts)
		next = responseSiigo.LinksV2.Next != nil
	}
	return allProducts
}

// castInvoice function is to convert the data received as a response from Siigo into JSON format, and assign
// them to a slice of products (products) to later join them with an existing slice of products (allProducts).
//
//	Params:
//		responseSiigo: a value of type model.InvoiceSiigo, which is expected to contain the Siigo response data in JSON format.
//		allProducts: a slice of model.InvoiceSiigo that contains all the products obtained so far.
//	Return:
//		A slice of model.InvoiceSiigo, which contains all the products obtained so far, including the new products
//		obtained from responseSiigo.
func castInvoice(responseSiigo interface{}, allInvoice []model.InvoiceSiigo) []model.InvoiceSiigo {
	invoces := new([]model.InvoiceSiigo)
	result, _ := json.Marshal(responseSiigo)
	json.Unmarshal(result, invoces)
	for _, invoice := range *invoces {
		allInvoice = append(allInvoice, invoice)
	}
	return allInvoice
}

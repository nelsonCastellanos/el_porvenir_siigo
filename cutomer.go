package p

import (
	"el_porvenir.com/cloudfunction/model"
	"encoding/json"
	"fmt"
)

// getCustomer retrieves all customers from the Siigo API using the provided token.
// Returns an array of model.CustomerSiigo structs containing the customer information.
//
// Parameters:
// - token: a model.TokenResult struct containing the access token for authenticating with the Siigo API
//
// Returns:
// - an array of model.CustomerSiigo structs representing all customers returned by the Siigo API
func getCustomer(token model.TokenResult) []model.CustomerSiigo {
	var allCustomers []model.CustomerSiigo
	responseSiigo := new(model.ResponseSiigo)
	err := getWithBearerToken(CUSTOMER, token.AccessToken, responseSiigo)
	if err != nil {
		panic("Error to get value auth from siigo")
	}
	allCustomers = castCustomer(responseSiigo.Results, allCustomers)
	next := responseSiigo.LinksV2.Next != nil
	for next {
		path := getPath(responseSiigo.LinksV2.Next.Href)
		responseSiigo = new(model.ResponseSiigo)
		getWithBearerToken(path, token.AccessToken, responseSiigo)
		allCustomers = castCustomer(responseSiigo.Results, allCustomers)
		next = responseSiigo.LinksV2.Next != nil
	}
	return allCustomers
}

// castCustomer function is to convert the data received as a response from Siigo into JSON format, and assign
// them to a slice of customers (customers) to later join them with an existing slice of customers (allcustomers).
//
//	Params:
//		responseSiigo: a value of type model.CustomerSiigo, which is expected to contain the Siigo response data in JSON format.
//		allcustomers: a slice of model.CustomerSiigo that contains all the customers obtained so far.
//	Return:
//		A slice of model.CustomerSiigo, which contains all the customers obtained so far, including the new customers
//		obtained from responseSiigo.
func castCustomer(responseSiigo interface{}, allCustomers []model.CustomerSiigo) []model.CustomerSiigo {
	customers := new([]model.CustomerSiigo)
	result, _ := json.Marshal(responseSiigo)
	err := json.Unmarshal(result, customers)
	if err != nil {
		fmt.Print(err.Error())
	}
	for _, customer := range *customers {
		allCustomers = append(allCustomers, customer)
	}
	return allCustomers
}

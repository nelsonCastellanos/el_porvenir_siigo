package p

import (
	"el_porvenir.com/cloudfunction/model"
	"encoding/json"
	"fmt"
)

// getProducts retrieves all products from Siigo using the provided access token.
// It paginates through all available results and returns them as an array of ProductSiigo objects.
//
// Parameters:
// - token: the access token to use for authentication with Siigo
//
// Returns:
// - an array of ProductSiigo objects representing all products in Siigo
func getProducts(token model.TokenResult) []model.ProductSiigo {
	var allProducts []model.ProductSiigo
	responseSiigo := new(model.ResponseSiigo)
	err := getWithBearerToken(PRODUCT, token.AccessToken, responseSiigo)
	if err != nil {
		panic("Error to get value auth from siigo")
	}
	allProducts = castProducts(responseSiigo.Results, allProducts)
	next := responseSiigo.LinksV1.Next != nil
	for next {
		path := getPath(responseSiigo.LinksV1.Next.Href)
		responseSiigo = new(model.ResponseSiigo)
		getWithBearerToken(path, token.AccessToken, responseSiigo)
		allProducts = castProducts(responseSiigo.Results, allProducts)
		next = responseSiigo.LinksV1.Next != nil
	}
	return allProducts
}

// castProducts function is to convert the data received as a response from Siigo into JSON format, and assign
// them to a slice of products (products) to later join them with an existing slice of products (allProducts).
//
//	Params:
//		responseSiigo: a value of type interface{}, which is expected to contain the Siigo response data in JSON format.
//		allProducts: a slice of model.ProductSiigo that contains all the products obtained so far.
//	Return:
//		A slice of model.ProductSiigo, which contains all the products obtained so far, including the new products
//		obtained from responseSiigo.
func castProducts(responseSiigo interface{}, allProducts []model.ProductSiigo) []model.ProductSiigo {
	products := new([]model.ProductSiigo)
	result, err := json.Marshal(responseSiigo)
	if err != nil {
		fmt.Print(err.Error())
	}
	err = json.Unmarshal(result, products)
	if err != nil {
		fmt.Print(err.Error())
	}
	for _, product := range *products {
		allProducts = append(allProducts, product)
	}
	return allProducts
}

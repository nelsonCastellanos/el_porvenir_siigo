// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"context"
	"el_porvenir.com/cloudfunction/model"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	PageSize = "5"
	PRODUCT  = "/v1/products?page_size=" + PageSize + "&page=1"
	CUSTOMER = "/v1/customers?page_size=" + PageSize + "&page=1"
	INVOICE  = "/v1/invoices?page_size=" + PageSize + "&page=1"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

// ElPorvenirSiigo consumes a Pub/Sub message.
func ElPorvenirSiigo(ctx context.Context, m PubSubMessage) error {
	bigQueryPorvenir := NewBigQueryPorvenir(ctx)
	bigQueryPorvenir.SiigoElPorvenir()
	token := getAuthToken()
	products := getProducts(token)
	customers := getCustomer(token)
	invoices := getInvoice(token)
	bigQueryPorvenir.InsertProduct(products)
	bigQueryPorvenir.InserCustomer(customers)
	bigQueryPorvenir.InserInvoice(invoices)
	return nil
}

// The getPath function extracts the path and query parameters of the next link in a Siigo API
// response and returns a string representation of them with the page query parameter preserved.

// Return:
//
//	The getPath function returns a string representing the path and query parameters of the next link
//	in the responseSiigo struct, with the page query parameter set to the same value it had before.
//	If there is any error during the parsing of the URL, an empty string is returned.
func getPath(href string) string {
	url, _ := url.Parse(href)
	url.Query().Set("page", url.Query().Get("page"))
	return url.Path + "?" + url.RawQuery
}

// getAuthToken realiza una solicitud POST para obtener un token de autenticación desde el servidor de Siigo
// utilizando las credenciales de acceso proporcionadas a través de las variables de entorno EMAIL y SiGGO_API_KEY.
// Retorna:
// - Una estructura de modelo TokenResult que contiene el token de autenticación generado por el servidor de Siigo.
// - Si se produce un error durante la solicitud POST, se lanza una excepción.
func getAuthToken() model.TokenResult {
	auth := model.TokenResult{}
	err := Post("/auth", model.Auth{
		UserName:  os.Getenv("EMAIL"),
		AccessKey: os.Getenv("SiGGO_API_KEY"),
	}, &auth)
	if err != nil {
		panic("Error to get value auth from siigo")
	}
	return auth
}

// getWithBearerToken realiza una solicitud GET a una URL específica con un token de autenticación de tipo Bearer.
// Parámetros:
// - path: La URL a la que se debe enviar la solicitud GET.
// - authToken: El token de autenticación de tipo Bearer a incluir en el encabezado de autorización de la solicitud.
// Retorna:
// - Un string que representa la respuesta del servidor a la solicitud GET.
// - Un puntero al error encontrado al realizar la solicitud GET, si lo hay.
func getWithBearerToken(path string, authToken string, result interface{}) *error {
	host := os.Getenv("HOST")
	client := &http.Client{}
	req, err := http.NewRequest("GET", host+path, nil)
	if err != nil {
		return &err
	}
	req.Header.Set("Authorization", "Bearer "+authToken)
	resp, err := client.Do(req)
	if err != nil {
		return &err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &err
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		return &err
	}
	return nil
}

// Post realiza una solicitud POST a una URL específica con un payload JSON proporcionado
// y devuelve la respuesta JSON decodificada en la estructura de resultado proporcionada.
// Parámetros:
// - url: La URL a la que se debe enviar la solicitud POST.
// - payload: El payload JSON que se debe incluir en la solicitud POST.
// - result: Un puntero a una estructura en la que se almacenará la respuesta decodificada del servidor.
// Retorna:
// - Un puntero al error encontrado al realizar la solicitud POST, si lo hay.
func Post(path string, payload interface{}, result interface{}) *error {
	host := os.Getenv("HOST")
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return &err
	}
	payloadReader := strings.NewReader(string(payloadJSON))
	resp, err := http.Post(host+path, "application/json", payloadReader)
	if err != nil {
		return &err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &err
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		return &err
	}
	return nil
}

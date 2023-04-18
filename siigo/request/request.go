package siigo_request

import (
	siigo_model "el_porvenir.com/cloudfunction/siigo/model"
	"el_porvenir.com/cloudfunction/siigo/model/auth"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
)

// The GetPath function extracts the path and query parameters of the next link in a Siigo API
// response and returns a string representation of them with the page query parameter preserved.
// Return:
//
//	The getPath function returns a string representing the path and query parameters of the next link
//	in the responseSiigo struct, with the page query parameter set to the same value it had before.
//	If there is any error during the parsing of the URL, an empty string is returned.
func GetPath(href string) string {
	url, _ := url.Parse(href)
	url.Query().Set("page", url.Query().Get("page"))
	return url.Path + "?" + url.RawQuery
}

// GetAuthToken realiza una solicitud POST para obtener un token de autenticación desde el servidor de Siigo
// utilizando las credenciales de acceso proporcionadas a través de las variables de entorno EMAIL y SiGGO_API_KEY.
// Retorna:
// - Una estructura de modelo TokenResult que contiene el token de autenticación generado por el servidor de Siigo.
// - Si se produce un error durante la solicitud POST, se lanza una excepción.
func GetAuthToken() auth.TokenResult {
	authToken := auth.TokenResult{}
	err := Post("/auth", auth.Auth{
		UserName:  os.Getenv("EMAIL"),
		AccessKey: os.Getenv("SiGGO_API_KEY"),
	}, &authToken)
	if err != nil {
		panic("Error to get value auth from siigo")
	}
	return authToken
}

// GetWithBearerToken realiza una solicitud GET a una URL específica con un token de autenticación de tipo Bearer.
// Parámetros:
// - path: La URL a la que se debe enviar la solicitud GET.
// - authToken: El token de autenticación de tipo Bearer a incluir en el encabezado de autorización de la solicitud.
// Retorna:
// - Un string que representa la respuesta del servidor a la solicitud GET.
// - Un puntero al error encontrado al realizar la solicitud GET, si lo hay.
func GetWithBearerToken(path string, authToken string, result interface{}) *error {
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

func CastItems(responseSiigo interface{}, item interface{}) interface{} {
	items := reflect.New(reflect.TypeOf(item)).Interface()
	result, err := json.Marshal(responseSiigo)
	if err != nil {
		fmt.Print(err.Error())
	}
	err = json.Unmarshal(result, items)
	if err != nil {
		fmt.Print(err.Error())
	}
	return items
}

func getLinks(response *siigo_model.ResponseSiigo) siigo_model.Links {
	if response.LinksV1 != nil {
		return *response.LinksV1
	}
	return *response.LinksV2
}

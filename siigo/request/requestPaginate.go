package siigo_request

import (
	siigo_model "el_porvenir.com/cloudfunction/siigo/model"
	"el_porvenir.com/cloudfunction/siigo/model/auth"
)

type RequestPaginate struct {
	items []interface{}
	item  interface{}
	path  string
}

func NewRequestPaginate(item interface{}, path string) *RequestPaginate {
	return &RequestPaginate{
		item: item,
		path: path,
	}
}

func (s *RequestPaginate) GetAllItems(token auth.TokenResult) []interface{} {
	responseSiigo := new(siigo_model.ResponseSiigo)
	err := GetWithBearerToken(s.path, token.AccessToken, responseSiigo)
	if err != nil {
		panic("Error to get value auth from siigo")
	}
	allItems := make([]interface{}, 0)
	for _, result := range responseSiigo.Results {
		allItems = append(allItems, CastItems(result, s.item))
	}
	next := getLinks(responseSiigo).Next != nil
	for next {
		path := GetPath(getLinks(responseSiigo).Next.Href)
		responseSiigo = new(siigo_model.ResponseSiigo)
		GetWithBearerToken(path, token.AccessToken, responseSiigo)
		for _, result := range responseSiigo.Results {
			allItems = append(allItems, CastItems(result, s.item))
		}
		next = getLinks(responseSiigo).Next != nil
	}
	return allItems
}

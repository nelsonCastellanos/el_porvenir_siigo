package siigo_request

import "el_porvenir.com/cloudfunction/siigo/model/auth"

type SimpleRequest struct {
	items []interface{}
	item  interface{}
	path  string
}

func NewSimpleRequest(item interface{}, path string) *RequestPaginate {
	return &RequestPaginate{
		item: item,
		path: path,
	}
}

func (s *RequestPaginate) GetItems(token auth.TokenResult) []interface{} {
	responseSiigo := new([]interface{})
	err := GetWithBearerToken(s.path, token.AccessToken, responseSiigo)
	if err != nil {
		panic("Error to get value auth from siigo")
	}
	allItems := make([]interface{}, 0)
	for _, result := range *responseSiigo {
		allItems = append(allItems, CastItems(result, s.item))
	}
	return allItems
}

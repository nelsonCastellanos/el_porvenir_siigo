package siigo_model

type ResponseSiigo struct {
	LinksV1    *Links        `json:"__links"`
	LinksV2    *Links        `json:"_links"`
	Pagination Pagination    `json:"pagination"`
	Results    []interface{} `json:"results"`
}
type Next struct {
	Href string `json:"href"`
}
type Self struct {
	Href string `json:"href"`
}
type Links struct {
	Next *Next `json:"next"`
	Self Self  `json:"self"`
}
type Pagination struct {
	Page         int `json:"page"`
	PageSize     int `json:"page_size"`
	TotalResults int `json:"total_results"`
}

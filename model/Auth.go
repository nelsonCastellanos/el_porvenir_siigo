package model

type Auth struct {
	UserName  string `json:"username"`
	AccessKey string `json:"access_key"`
}

type TokenResult struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

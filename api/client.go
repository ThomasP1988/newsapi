package newsapi

import "net/http"

type Client struct {
	APIKey     string
	HTTPClient *http.Client
}

func NewClient(APIKey string, httpClient *http.Client) Client {
	return Client{
		APIKey:     APIKey,
		HTTPClient: httpClient,
	}
}

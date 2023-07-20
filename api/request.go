package newsapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) Request(ctx context.Context, endpoint Endpoint, params *url.Values, output any) error {
	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	urlRequest := BaseURL + string(endpoint)

	if params != nil {
		urlRequest = urlRequest + "?" + params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlRequest, nil)
	if err != nil {
		return err
	}

	req.Header[Header_XAPIKey] = []string{c.APIKey}

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	statusOK := res.StatusCode >= 200 && res.StatusCode < 300
	if !statusOK {
		return handleError(body)
	}

	fmt.Printf("res: %+v\n", res)

	return json.Unmarshal(body, output)
}

func handleError(body []byte) error {

	newsAPIError := NewsAPIError{}

	err := json.Unmarshal(body, &newsAPIError)

	if err != nil {
		return errors.New("unknown error response " + string(body))
	}

	return newsAPIError
}

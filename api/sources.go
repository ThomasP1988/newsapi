package newsapi

import (
	"context"
	"net/url"
)

type SourcesArgs struct {
	Country  string
	Category Category
	Language string
}

type SourcesResponse struct {
	Status  string   `json:"status"`
	Sources []Source `json:"sources"`
}

func (c *Client) Sources(ctx context.Context, args SourcesArgs) (*SourcesResponse, error) {
	params := &url.Values{}

	FormatParamString("category", params, string(args.Category))
	FormatParamString("country", params, args.Country)
	FormatParamString("language", params, args.Language)

	output := &SourcesResponse{}

	err := c.Request(ctx, Endpoint_Sources, params, output)

	if err != nil {
		return nil, err
	}

	return output, nil
}

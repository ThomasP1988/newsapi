package newsapi

import (
	"context"
	"net/url"
)

type TopHeadlinesArgs struct {
	Country  string
	Category Category
	Q        string
	QInTitle string
	Sources  []string
	PageSize int16
	Page     int32
}

type TopHeadlinesResponse struct {
	Status       string    `json:"status"`
	TotalResults int16     `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

func (c *Client) TopHeadlines(ctx context.Context, args TopHeadlinesArgs) (*TopHeadlinesResponse, error) {
	params := &url.Values{}
	FormatParamString("q", params, args.Q)
	FormatParamString("qInTitle", params, args.QInTitle)
	FormatParamString("category", params, string(args.Category))
	FormatParamString("country", params, args.Country)
	FormatParamMultipleString("sources", params, args.Sources)
	FormatParamInt("pageSize", params, int(args.PageSize))
	FormatParamInt("page", params, int(args.Page))

	output := &TopHeadlinesResponse{}

	err := c.Request(ctx, Endpoint_TopHeadlines, params, output)

	if err != nil {
		return nil, err
	}

	return output, nil
}

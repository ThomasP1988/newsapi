package newsapi

import (
	"context"
	"net/url"
	"time"
)

type EverythingArgs struct {
	Q              string
	QInTitle       string
	SearchIn       []SearchIn
	Sources        []string
	Domains        []string
	ExcludeDomains []string
	From           time.Time
	To             time.Time
	Language       string
	SortBy         SortBy
	PageSize       int16
	Page           int32
}

type EverythingResponse struct {
	Status       string    `json:"status"`
	TotalResults int16     `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

func (c *Client) Everything(ctx context.Context, args EverythingArgs) (*EverythingResponse, error) {
	params := &url.Values{}
	FormatParamString("q", params, args.Q)
	FormatParamString("qInTitle", params, args.QInTitle)
	FormatParamMultipleString("searchIn", params, args.SearchIn)
	FormatParamMultipleString("sources", params, args.Sources)
	FormatParamMultipleString("domains", params, args.Domains)
	FormatParamMultipleString("excludeDomains", params, args.ExcludeDomains)
	FormatParamDate("from", params, args.From)
	FormatParamDate("to", params, args.To)
	FormatParamString("language", params, args.Language)
	FormatParamString("sortBy", params, string(args.SortBy))
	FormatParamInt("pageSize", params, int(args.PageSize))
	FormatParamInt("page", params, int(args.Page))

	output := &EverythingResponse{}

	err := c.Request(ctx, Endpoint_Everything, params, output)

	if err != nil {
		return nil, err
	}

	return output, nil
}

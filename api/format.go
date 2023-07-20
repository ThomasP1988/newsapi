package newsapi

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

func FormatParamDate(key string, params *url.Values, input time.Time) {
	if !input.IsZero() {
		params.Add(key, input.Format(time.RFC3339))
	}
}

func FormatParamString(key string, params *url.Values, input string) {
	if input != "" {
		params.Add(key, input)
	}
}

func FormatParamMultipleString(key string, params *url.Values, input []string) {
	if len(input) > 0 {
		params.Add(key, strings.Join(input, ","))
	}
}

func FormatParamInt(key string, params *url.Values, input int) {
	if input != 0 {
		params.Add("key", strconv.Itoa(int(input)))
	}
}

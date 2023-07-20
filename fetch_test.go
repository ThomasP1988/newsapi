package news

import (
	"context"
	"fmt"
	newsapi "newsapi/api"
	"testing"
)

var API_KEY = ""

func TestEverything(t *testing.T) {

	client := newsapi.NewClient(API_KEY, nil)

	response, err := client.Everything(context.Background(), newsapi.EverythingArgs{
		QInTitle: "test",
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("response: %v\n", response)
	}

	t.Fail()

}
func TestTopHeadlines(t *testing.T) {

	client := newsapi.NewClient(API_KEY, nil)

	response, err := client.TopHeadlines(context.Background(), newsapi.TopHeadlinesArgs{
		Category: newsapi.Category_Entertainment,
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("response: %v\n", response)
	}

	t.Fail()
}
func TestSources(t *testing.T) {

	client := newsapi.NewClient(API_KEY, nil)

	response, err := client.Sources(context.Background(), newsapi.SourcesArgs{})

	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("response: %v\n", response)
	}

	t.Fail()
}

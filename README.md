Implementatin of News API v2 in Go

## install

`go get github.com/ThomasP1988/newsapi`

## example

    var API_KEY = ""
	
	client := newsapi.NewClient(API_KEY, nil)

	response, err := client.Everything(context.Background(), newsapi.EverythingArgs{
		QInTitle: "test",
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("response: %v\n", response)
	}

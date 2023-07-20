package newsapi

const BaseURL = "https://newsapi.org/v2"

const (
	Header_XAPIKey = "X-Api-Key"
)

type Endpoint = string

const (
	Endpoint_Everything   Endpoint = "/everything"
	Endpoint_TopHeadlines Endpoint = "/top-headlines"
	Endpoint_Sources      Endpoint = "/top-headlines/sources"
)

type SearchIn = string

const (
	SearchIn_Title       SearchIn = "title"
	SearchIn_Description SearchIn = "description"
	SearchIn_Content     SearchIn = "content"
)

type SortBy string

const (
	SortBy_Relevancy   SortBy = "relevancy"
	SortBy_Popularity  SortBy = "popularity"
	SortBy_PublishedAt SortBy = "publishedAt"
)

type Category string

const (
	Category_Business      Category = "business"
	Category_Entertainment Category = "entertainment"
	Category_General       Category = "general"
	Category_Health        Category = "health"
	Category_Science       Category = "science"
	Category_Sports        Category = "sports"
	Category_Technology    Category = "technology"
)

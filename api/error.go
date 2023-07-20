package newsapi

type NewsAPIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (err NewsAPIError) Error() string {
	return err.Code + " " + err.Message
}

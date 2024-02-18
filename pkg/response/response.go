package response

type HttpOkResponse struct {
	Data interface{} `json:"data"`
}

type HttpErrorResponse struct {
	Errors []string `json:"errors"`
}

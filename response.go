package sendinblue

type OKResponse struct {
	MessageID string `json:"messageId"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

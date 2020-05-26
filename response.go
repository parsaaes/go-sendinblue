package sendinblue

type OKResponse struct {
	MessageID string `json:"message_id"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

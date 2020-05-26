package sendinblue

const (
	SendTransactionalEmailURL = "https://api.sendinblue.com/v3/smtp/email"
)

type Client struct {
	APIKey string
}

func New(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

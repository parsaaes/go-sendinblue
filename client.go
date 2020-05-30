package sendinblue

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// SendTransactionalEmailEndpoint is the endpoint for sending transactional emails.
const SendTransactionalEmailEndpoint = "/v3/smtp/email"

// ErrSentWithBadResponse means server responded with HTTP code 201
// but body cannot be processed to retrieve the message id.
var ErrSentWithBadResponse = errors.New("email sent but with error on processing the response")

// Client represents a sendinblue client.
type Client struct {
	Client *resty.Client
}

// New creates a new sendinblue client configured with the given api-key and timeout.
func New(url, apiKey string, timeout time.Duration) *Client {
	return &Client{
		Client: resty.New().
			SetHostURL(url).
			SetHeader("api-key", apiKey).
			SetTimeout(timeout),
	}
}

// SendTransactionalEmail sends the given email and returns messageID returned by the server
// or error if it fails.
//
// It will return empty message id and ErrSentWithBadResponse if server responds
// with 201 and without a proper response.
//
// It's based on sendinblue's v3 API https://developers.sendinblue.com/reference#sendtransacemail.
func (c *Client) SendTransactionalEmail(email Email) (string, error) {
	resp, err := c.Client.R().SetBody(email).Post(SendTransactionalEmailEndpoint)

	if err != nil {
		return "", err
	}

	if code := resp.StatusCode(); code != http.StatusCreated {
		if code == http.StatusBadRequest {
			var body ErrorResponse
			if err := json.Unmarshal(resp.Body(), &body); err != nil {
				return "", fmt.Errorf("bad request. error on processing the response: %v", err)
			}

			return "", fmt.Errorf("bad request. response: %+v", body)
		}

		return "", fmt.Errorf("failed. code: %d", code)
	}

	var body OKResponse
	if err := json.Unmarshal(resp.Body(), &body); err != nil {
		return "", ErrSentWithBadResponse
	}

	return body.MessageID, nil
}

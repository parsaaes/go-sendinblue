package sendinblue

type User struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email"`
}

type Email struct {
	Sender      User   `json:"sender"`
	To          []User `json:"to"`
	ReplyTo     User   `json:"reply_to"`
	Subject     string `json:"subject"`
	HTMLContent string `json:"html_content"`
}

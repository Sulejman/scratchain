package model

// Block model
type Block struct {
	Header       string
	Previous     string
	Next         string
	Transactions []string
}

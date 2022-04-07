package models

type Auction struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
}

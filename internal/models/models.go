package models

type Auction struct {
	Title       string `json:"title"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
}

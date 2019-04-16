package models

type User struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Link  `json:"link,omitempty"`
}

type Link struct {
	Title string `json:"title,omitempty"`
	Url   string `json:"url,omitempty"`
}

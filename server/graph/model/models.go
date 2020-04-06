package model

type Book struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Picture   *string   `json:"picture"`
	Publisher string    `json:"publisher"`
	Statuss   []*Status `json:"statuss"`
}

type Status struct {
	ID         string  `json:"id"`
	Currently  string  `json:"currently"`
	Icon       *string `json:"icon"`
	Importance int     `json:"importance"`
}

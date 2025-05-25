package book

type Criteria struct {
	Title       string 	`json:"title"`
	Author      string 	`json:"author"`
	ISBN        string 	`json:"isbn"`
	ReleaseDate string 	`json:"release_date"`
	SortBy      string 	`json:"sort_by"`
	Order       string 	`json:"order"`
	Page        int 	`json:"page"`
	Limit       int 	`json:"limit"`
}

package model

type MovieSearchRequest struct {
	Keyword string `json:"keyword" validate:"required"`
	Page    int32  `json:"page" validate:"required"`
}

type Movie struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbId string `json:"imdbId"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

type MovieSearchResponse struct {
	Movies []*Movie `json:search`
	TotalResult string   `json:"totalResults"`
	Response    string   `json:"Response"`
}


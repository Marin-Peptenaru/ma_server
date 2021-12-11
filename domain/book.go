package domain

type Book struct {
	Isbn            string      `json:"isbn"`
	Title           string      `json:"title"`
	Author          string      `json:"author"`
	Pages           int         `json:"pages"`
	Genre           string      `json:"genre"`
	PublicationDate string  `json:"publicationDate"`
}

package main

// Book Struct
type Books struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func NewBook(ID string, Isbn string, Title string) (*Books, error) {

	return &Books{
		ID:    ID,
		Isbn:  Isbn,
		Title: Title,
	}, nil
}

func NewAuthor(FirstName string, LastName string) (*Author, error) {
	return &Author{
		FirstName: FirstName,
		LastName:  LastName,
	}, nil

}

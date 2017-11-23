package api

import (
	"encoding/json"
	"net/http"
)

//Book type with Name, Author and ISBN
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
	//define the book
}

//ToJSON to be used for marshalling of Book type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// FromJSON to be used for marshalling of Book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

// Books slice of all known books
var Books = []Book{
	Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Admin", ISBN: "0345391802"},
	Book{Title: "Cloud Native Go", Author: "M. -L. Reimer", ISBN: "0123456789"},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

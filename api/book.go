package api

import (
	"encoding/json"
	"net/http"
)

//Book type with Name, Author and ISBN
type Book struct {
	Title  string
	Author string
	ISBN   string
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
	return Book{}
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {

}

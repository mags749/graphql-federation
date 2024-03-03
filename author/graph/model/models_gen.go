// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Author struct {
	ID   int     `json:"id"`
	Name *string `json:"name,omitempty"`
	Age  *int    `json:"age,omitempty"`
}

func (Author) IsEntity() {}

type Book struct {
	ID     int     `json:"id"`
	Author *Author `json:"author,omitempty"`
}

func (Book) IsEntity() {}

type Query struct {
}

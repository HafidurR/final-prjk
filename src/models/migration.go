package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string    `json:"password"`
	Age      int    `json:"age"`
	Books 		[]Book
}

type Book struct {
	gorm.Model
	Title  				string `json:"title"`
	Description   string `json:"description"`
	Publisher 		string	`json:"publisher`
	AuthorID 			int			`json:"author_id"`
}
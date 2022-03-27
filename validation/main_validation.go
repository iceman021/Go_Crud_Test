package main

import (
	"fmt"
	"github.com/go-playground/validator"
)

type Author struct {
	ID      int    `validate:"gte=1,lte=5"`
	Name    string `json:"name" validate:"alpha"`
	Surname string `json:"surname" validate:"alpha"`
}

type Book struct {
	ID              int       `validate:"gte=1,lte=5"`
	Title           string    `json:"title" validate:"alpha"`
	Description     string    `json:"description" validate:"alpha"`
	PublicationDate string 	  `validate:"datetime"`
}

/* User contains user information
*/

// use single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {

	validate = validator.New()

	book := &Book{
		ID:   7,
		Title:    "Test",
		Description:  "markovic",
		PublicationDate: "OneOne",
	}

	author := &Author{
		ID:   6,
		Name:    "Mark",
		Surname:  "twain",
	}


	err := validate.Struct(author)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		err := validate.Struct(book)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}
	

		fmt.Println("----- Listing of tag fields with error ------")

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.StructField())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println("---------------")
		}
		return
	}
}}
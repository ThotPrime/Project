package mocks

//import "github.com/tutorials/go/crud/pkg/models"
import "github.com/ThotPrime/Project/tree/master/Project/pkg/models"

var Books = []models.Book{
	{
		Id:     1,
		Title:  "Golang",
		Author: "Gopher",
		Desc:   "A book for Go",
		Pages:  27,
	},
}
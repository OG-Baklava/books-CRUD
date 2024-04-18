package handlers

import (
	"books/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB
func GetBooks(w http.ResponseWriter, r *http.Request){
	var books []models.Book
	DB.Find(&books)
}

func GetBook(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r)
	var book models.Book
	if err := DB.First (&book, vars["id"]).Error; err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(book)
}
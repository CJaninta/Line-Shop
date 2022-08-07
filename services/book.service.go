package services

import (
	"assignment/models"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBook(book *models.BookReq) error {
	_, err := bookCollection.InsertOne(ctx, book)
	return err
}

func GetBookById(id primitive.ObjectID) (*models.BookRes, error) {

	var book *models.BookRes

	err := bookCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	if err != nil {
		return book, err
	}

	return book, err
}

func GetBook() ([]*models.BookRes, error) {

	var books []*models.BookRes

	data, err := bookCollection.Find(ctx, bson.M{})

	for data.Next(ctx) {
		var book models.BookRes
		err := data.Decode(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, err
}

func GetSearchBook(search string) ([]*models.SearchBookRes, error) {

	var books []*models.SearchBookRes

	texts := strings.Fields(search)

	searchText := []map[string]interface{}{}
	for _, text := range texts {
		searchText = append(searchText, bson.M{"title": bson.M{"$regex": text}})
	}
	filter := bson.M{"$and": searchText}

	data, err := bookCollection.Find(ctx, filter)
	for data.Next(ctx) {
		var book models.SearchBookRes

		err := data.Decode(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, err
}

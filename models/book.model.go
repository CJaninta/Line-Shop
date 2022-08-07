package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Image struct {
	JPEG string `json:"jpeg"`
	AVIF string `json:"avif"`
	WEBP string `json:"webp"`
}

type BookReq struct {
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Rating      int     `json:"rating"`
	Author      string  `json:"author"`
	Images      Image   `json:"images`
}

type BookRes struct {
	Id          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title       string             `json:"title"`
	Price       float64            `json:"price"`
	Description string             `json:"description"`
	Rating      int                `json:"rating"`
	Author      string             `json:"author"`
	Images      Image              `json:"images"`
}

type SearchBookRes struct {
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	Rating int     `json:"rating"`
	Author string  `json:"author"`
	Images Image   `json:"images"`
}

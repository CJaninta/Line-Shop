package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//For request
type OrderReq struct {
	ProductID primitive.ObjectID `bson:"product_id" json:"id,omitempty"`
	Quantity  int                `json:"quantity"`
}

//For send to db
type Order struct {
	Products []OrderReq `json:"products"`
	Summary  float64    `json:"summary"`
	Checkout bool       `json:"checkout"`
}

//For model of db
type OrderModel struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Products []OrderReq         `json:"products"`
	Summary  float64            `json:"summary"`
	Checkout bool               `json:"checkout"`
}

//For Response summary
type OrderProduct struct {
	Title    string  `json:"name"`
	Price    float64 `json:"price"`
	Rating   int     `json:"rating"`
	Author   string  `json:"author"`
	Images   Image   `json:"coverImage"`
	Quantity int     `bson:"quantity" json:"quantity"`
}

//For response summary
type OrderSummaryRes struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Products []OrderProduct     `json:"products"`
	Discount float64            `json:"discount"`
	Summary  float64            `json:"summary"`
}

type OrderSummary struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Products []BookRes          `json:"products"`
	Discount float64            `json:"discount"`
	Summary  float64            `json:"summary"`
}

type OrderCheckoutReq struct {
	Id primitive.ObjectID `json:"id,omitempty"`
}

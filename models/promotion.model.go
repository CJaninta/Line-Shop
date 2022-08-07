package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PromotionRes struct {
	Id        primitive.ObjectID   `bson:"_id" json:"id,omitempty"`
	Type      string               `json:"type"`
	Name      string               `json:"name"`
	TargetIds []primitive.ObjectID `bson:"targetIds_id" json:"targetIds,omitempty"`
	Discounts []string             `json:"discounts"`
}

type PromotionReq struct {
	Type      string               `json:"type"`
	Name      string               `json:"name"`
	TargetIds []primitive.ObjectID `bson:"targetIds_id" json:"targetIds,omitempty"`
	Discounts []string             `json:"discounts"`
}

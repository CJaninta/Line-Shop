package services

import (
	"assignment/models"

	"go.mongodb.org/mongo-driver/bson"
)

//TODO: delete (for test)
func CreatePromotion(promotion *models.PromotionReq) error {

	//check targetid is in db
	for _, v := range promotion.TargetIds {
		_, err := GetBookById(v)
		if err != nil {
			return err
		}
	}

	_, err := promotionCollection.InsertOne(ctx, promotion)
	return err
}

func GetPromotion() ([]*models.PromotionRes, error) {

	var promotions []*models.PromotionRes

	data, err := promotionCollection.Find(ctx, bson.M{})

	for data.Next(ctx) {
		var promotion models.PromotionRes
		err := data.Decode(&promotion)
		if err != nil {
			return nil, err
		}
		promotions = append(promotions, &promotion)
	}
	return promotions, err
}

package services

import (
	"assignment/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateFirstOrder(order *models.OrderReq) error {
	var preOrderModel = []models.OrderReq{{
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
	}}

	book, err := GetBookById(order.ProductID)
	if err != nil {
		return err
	}

	sum := float64(book.Price * float64(order.Quantity))
	var preOrder = models.Order{
		Products: preOrderModel,
		Summary:  sum,
		Checkout: false,
	}

	_, err = orderCollection.InsertOne(ctx, preOrder)
	return err
}

func AddProductExistOrder(lastRecord models.OrderModel, order *models.OrderReq) error {

	var orderProduct models.OrderModel
	filterProduct := bson.M{"_id": lastRecord.Id, "products.product_id": order.ProductID}
	err := orderCollection.FindOne(ctx, filterProduct).Decode(&orderProduct)

	if err != nil {
		updateOrder := bson.M{"$push": bson.M{"products": bson.M{
			"product_id": order.ProductID,
			"quantity":   order.Quantity,
		}}}
		filter := bson.M{"_id": lastRecord.Id}
		_, err := orderCollection.UpdateOne(ctx, filter, updateOrder)
		if err != nil {
			return err
		}

	} else {
		var amount int = 0
		for _, v := range orderProduct.Products {
			if v.ProductID == order.ProductID {
				amount = v.Quantity
				break
			}
		}

		filter := bson.M{
			"_id":                 lastRecord.Id,
			"products.product_id": order.ProductID,
		}

		updateOrderAmount := bson.M{
			"$set": bson.M{
				"products.$.quantity": amount + order.Quantity,
			},
		}
		_, err := orderCollection.UpdateOne(ctx, filter, updateOrderAmount)
		if err != nil {
			return err
		}

	}

	err = UpdateSummary(lastRecord.Id, order, lastRecord.Summary)
	return err
}

func UpdateSummary(id primitive.ObjectID, order *models.OrderReq, prevAmount float64) error {

	book, err := GetBookById(order.ProductID)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": id,
	}

	amount := prevAmount + float64(book.Price*float64(order.Quantity))

	updateOrderSummary := bson.M{
		"$set": bson.M{
			"summary": amount,
		},
	}
	_, err = orderCollection.UpdateOne(ctx, filter, updateOrderSummary)
	return err
}

func GetLastRecord() (models.OrderModel, error) {
	opts := options.FindOne().SetSort(bson.M{"$natural": -1})

	var lastRecord models.OrderModel
	err := orderCollection.FindOne(ctx, bson.M{}, opts).Decode(&lastRecord)

	return lastRecord, err
}

func CreateOrder(order *models.OrderReq) error {

	_, err := GetBookById(order.ProductID)
	if err != nil {
		return err
	}

	opts := options.FindOne().SetSort(bson.M{"$natural": -1})

	var lastRecord models.OrderModel
	err = orderCollection.FindOne(ctx, bson.M{}, opts).Decode(&lastRecord)

	if lastRecord.Checkout || err != nil {
		err := CreateFirstOrder(order)
		if err != nil {
			return err
		}
		return nil
	} else {
		err := AddProductExistOrder(lastRecord, order)
		if err != nil {
			return err
		}
	}
	return err
}

func GetOrderById(id primitive.ObjectID) (*models.OrderModel, error) {
	var order *models.OrderModel

	err := orderCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)
	if err != nil {
		return order, err
	}

	return order, err
}

func CheckoutOrder(id primitive.ObjectID) error {
	orderProduct, err := GetOrderById(id)
	if err != nil {
		return errors.New("Please try again later.")
	}

	if orderProduct.Checkout {
		return errors.New("This order has checked out.")
	}

	filter := bson.M{
		"_id": id,
	}

	updateCheckout := bson.M{
		"$set": bson.M{
			"checkout": true,
		},
	}

	_, err = orderCollection.UpdateOne(ctx, filter, updateCheckout)
	return err
}

func Summary() (*models.OrderSummary, error) {

	var order *models.OrderSummary

	filter := bson.M{"$match": bson.M{"checkout": false}}

	populate := bson.M{"$lookup": bson.M{
		"from":         "book",
		"localField":   "products.product_id",
		"foreignField": "_id",
		"as":           "products",
	}}

	data, _ := orderCollection.Aggregate(ctx, []bson.M{
		filter, populate,
	})

	if data.Next(ctx) {
		err := data.Decode(&order)
		if err != nil {
			return nil, err
		}

	}

	return order, nil
}

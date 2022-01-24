package persistence

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var reviewsCollection *mongo.Collection

func init() {

	reviewsCollection = DB.Collection("reviews")

	dateRangeDigestIndexModel := mongo.IndexModel{
		Keys: bson.M{
			"applicationId": 1,
			"postedAt":      -1,
			"wasDigested":   1,
		}, Options: nil,
	}

	reviewsCollection.Indexes().CreateOne(context.TODO(), dateRangeDigestIndexModel)

}

func AddReviewsBulk(applicationId string, reviews []*Review) {

	if len(reviews) == 0 {
		log.Println("No reviews to be added")
		return
	}

	documents := make([]interface{}, len(reviews))

	for i, review := range reviews {
		documents[i] = review
	}

	_, err := reviewsCollection.InsertMany(context.TODO(), documents)

	if err != nil {
		log.Printf("Error while adding bulk reviews for application %s", applicationId)
		log.Fatal(err.Error())
	}

}

func RetrieveLastReviewAdded(applicationId string) *Review {

	filter := bson.M{
		"applicationId": applicationId,
	}

	findOptions := options.Find().SetLimit(1)

	lastReviewAdded := &Review{}

	cursor, err := reviewsCollection.Find(context.TODO(), filter, findOptions)

	if err != nil {
		log.Printf("Error while retrieving last review added for application %s", applicationId)
		log.Fatal(err.Error())
	}

	for cursor.Next(context.Background()) {
		err = cursor.Decode(lastReviewAdded)
		if err != nil {
			log.Printf("Error while decoding last review for application %s", applicationId)
			log.Fatal(err.Error())
		}
	}

	log.Println("Found last review Added", lastReviewAdded)

	return lastReviewAdded

}

func RetrieveNotDigestedReviews(applicationId string, limit int64, currentExecutionTime time.Time) []*Review {

	filter := bson.D{{"applicationId", applicationId}, {"postedAt", bson.D{{"$gt", currentExecutionTime}}}, {"wasDigested", false}}

	findOptions := options.Find().SetLimit(limit)

	cursor, err := reviewsCollection.Find(context.TODO(), filter, findOptions)

	if err != nil {
		log.Printf("Error while retrieving not digested reviews for application %s", applicationId)
		log.Fatal(err.Error())
	}

	reviews := []*Review{}

	for cursor.Next(context.Background()) {
		review := &Review{}

		err := cursor.Decode(review)

		if err != nil {
			log.Printf("Error while parsing retrieved review for application %s", applicationId)
			log.Fatal(err.Error())
		}

		reviews = append(reviews, review)
	}

	return reviews

}

func UpdateReviewsAsDigested(applicationId string, reviews []*Review) {
	opts := options.Update().SetUpsert(false)

	for _, review := range reviews {
		review.WasDigested = true
	}

	ids := make([]string, len(reviews))

	for i, review := range reviews {
		ids[i] = review.ReviewId
	}

	filter := bson.D{{"_id", bson.D{{"$in", ids}}}}

	update := bson.D{{"$set", bson.D{{"wasDigested", true}}}}

	_, err := reviewsCollection.UpdateMany(context.TODO(), filter, update, opts)

	if err != nil {
		log.Printf("Error while updating reviews for application %s as digested", applicationId)
		log.Fatal(err.Error())
	}

	log.Printf("Updated all documents successfuly for application %s", applicationId)
}

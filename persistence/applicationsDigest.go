package persistence

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var applicationsDigestCollection *mongo.Collection

func init() {
	log.Println("Initiating the applicationsDigest...")

	applicationsDigestCollection = DB.Collection("applicationDigests")

	applicationActiveIndexModel := mongo.IndexModel{
		Keys: bson.M{
			"isActive": 1, // data range in descending order
		}, Options: nil,
	}

	applicationsDigestCollection.Indexes().CreateOne(context.TODO(), applicationActiveIndexModel)
}

func AddApplicationAndActivate(applicationId string, sleepTime uint64, startTimeAt time.Time, pageSize uint64) *Application {

	opts := options.Update().SetUpsert(true)

	application := &Application{
		SleepTime:     sleepTime,
		StartDigestAt: startTimeAt,
		ApplicationId: applicationId,
		IsActive:      true,
		PageSize:      pageSize,
	}

	filter := bson.D{{"_id", applicationId}}

	update := bson.D{{"$set", application}}

	_, err := applicationsDigestCollection.UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		log.Printf("Error while inserting application %s into database", applicationId)
		log.Fatal(err)
	}

	log.Printf("Sucessfully inserted application %s into the database", applicationId)

	return application

}

func InactivateApplicationDigest(applicationId string) {

	filter := bson.D{{"_id", applicationId}}

	applicationChange := bson.D{{"isActive", false}}

	update := bson.D{{"$set", applicationChange}}

	result := applicationsDigestCollection.FindOneAndUpdate(context.TODO(), filter, update)

	if result.Err() != nil {
		log.Printf("Error while inactivating application %s into database", applicationId)
		log.Fatal(result.Err().Error())
	}

	log.Printf("Sucessfully inactivated application %s into the database", applicationId)

}

func RetrieveActiveApplications() []*Application {

	filter := bson.D{{"isActive", true}}

	cursor, err := applicationsDigestCollection.Find(context.TODO(), filter)

	if err != nil {
		log.Println("Error while querying active applications")
		log.Fatal(err)
	}

	var applications []*Application

	for cursor.Next(context.TODO()) {
		// Declare a result BSON object
		var result *Application
		err := cursor.Decode(&result)

		if err != nil {
			log.Println("Error while querying active applications")
			log.Fatal(err)
		}

		applications = append(applications, result)

	}

	log.Println("Sucessfully retrieved active applications")

	return applications
}

func UpdateApplicationDigest(application *Application) {

	filter := bson.D{{"_id", application.ApplicationId}}

	update := bson.D{{"$set", bson.D{{"lastDigest", application.LastDigest}}}}

	result := applicationsDigestCollection.FindOneAndUpdate(context.TODO(), filter, update)

	if result.Err() != nil {
		log.Printf("Error while updating application %s into database after digest", application.ApplicationId)
		log.Fatal(result.Err().Error())
	}

	log.Printf("Updated Last Digest for application %s successfully", application.ApplicationId)

}

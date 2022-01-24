package services

import (
	"log"
	"time"

	"github.com/Neiva07/review-digest-creator/persistence"
	"github.com/go-co-op/gocron"
)

func ExecuteRunningJobs() {
	applications := persistence.RetrieveActiveApplications()

	for _, app := range applications {
		ExecuteNewJob(app)
	}
}

func ExecuteNewJob(application *persistence.Application) {
	log.Printf("Starting execution of job of application %s", application.ApplicationId)

	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.StartAt(application.StartDigestAt)

	_, err := scheduler.Every(int(application.SleepTime)).Hours().Do(syncReviewAndGenerateReport, application)

	if err != nil {
		log.Printf("Error while creating cron job for application %s", application.ApplicationId)
		log.Fatalln(err.Error())
	}

	scheduler.StartAsync()

	_, nextRun := scheduler.NextRun()

	log.Printf("Application %s scheduled successfully. Next run will be in: ", application.ApplicationId)
	log.Println(nextRun)
}

func syncReviewAndGenerateReport(application *persistence.Application) {

	syncReviewsCollection(application.ApplicationId)

	rangeTime := time.Now().Add(time.Hour * (-time.Duration(application.SleepTime)))

	reviews := persistence.RetrieveNotDigestedReviews(application.ApplicationId, int64(application.PageSize), rangeTime)

	err := GenerateReport(application, reviews)

	if err != nil {
		log.Printf("Error while creating digest for application %s", application.ApplicationId)
		log.Println(err)
		return
	}

	persistence.UpdateReviewsAsDigested(application.ApplicationId, reviews)

	persistence.UpdateApplicationDigest(application)

}

func syncReviewsCollection(applicationId string) {

	lastReview := persistence.RetrieveLastReviewAdded(applicationId)

	reviewsToAdd := RetrieveAppleReviews(applicationId, lastReview)

	persistence.AddReviewsBulk(applicationId, reviewsToAdd)
}

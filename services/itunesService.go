package services

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Neiva07/review-digest-creator/persistence"
)

const ITUNES_RSS_API = "https://itunes.apple.com/us/rss"

const ITUNES_RSS_PAGE_LIMIT = 10

func RetrieveAppleReviews(applicationId string, lastReview *persistence.Review) []*persistence.Review {

	log.Printf("Retrieving Apple Reviews for appicationId %s", applicationId)

	hasSavedReview := false

	reviews := []*persistence.Review{}

	for page := 0; page < ITUNES_RSS_PAGE_LIMIT && !hasSavedReview; page++ {

		urlString := fmt.Sprintf("%s/customerreviews/id=%s/sortBy=mostRecent/page=%d/json", ITUNES_RSS_API, applicationId, page+1)

		itunesResponseObject := &ItunesRssModel{}

		err := GetJson(urlString, itunesResponseObject)

		if err != nil {
			log.Fatalln("Fatal error retrieving from Reviews from Itunes API: ", err)
		}

		for _, entry := range itunesResponseObject.Feed.Entry {

			review := convertItunesReviewIntoReview(entry, applicationId)

			if lastReview.ReviewId == review.ReviewId {
				hasSavedReview = true
				break
			}

			reviews = append(reviews, review)
		}
	}

	log.Printf("Found %d reviews not saved from application %s", len(reviews), applicationId)

	return reviews

}

func convertItunesReviewIntoReview(itunesReview ItunesReviewModel, applicationId string) *persistence.Review {

	postedAt, err := time.Parse(time.RFC3339, itunesReview.UpdatedAt.Label)

	if err != nil {
		log.Printf("Error while converting date of review %s", itunesReview.ReviewId.Label)
		log.Fatalln(err)
	}

	u64, err := strconv.ParseUint(itunesReview.Rating.Label, 10, 16)

	if err != nil {
		log.Printf("Error while converting rating of review %s", itunesReview.ReviewId.Label)
		fmt.Println(err)
	}
	rating := uint8(u64)

	return &persistence.Review{
		ApplciationId: applicationId,
		ReviewId:      itunesReview.ReviewId.Label,
		Title:         itunesReview.Title.Label,
		Author:        itunesReview.Author.Name.Label,
		Content:       itunesReview.Content.Label,
		Rating:        rating,
		PostedAt:      postedAt,
		WasDigested:   false,
	}
}

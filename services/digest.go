package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Neiva07/review-digest-creator/persistence"
)

const TITLE = "###Author\t\tTitle\t\tContent\t\tRating\t\tPosted Time"

const REVIEW_LINE_TEMPLATE = "%s\t\t%s\t\t%s\t\t%d\t\t%s"

const BREAK_LINE = "\n\n"

const DATE_FORMAT = "02/Jan/2006 15:04:05"

func GenerateReport(application *persistence.Application, reviews []*persistence.Review) error {

	if len(reviews) == 0 {
		errorMessage := fmt.Sprintf("No reviews to digest for application %s", application.ApplicationId)
		return errors.New(errorMessage)
	}

	content := TITLE + BREAK_LINE

	for _, review := range reviews {

		content += fmt.Sprintf(REVIEW_LINE_TEMPLATE,
			review.Author,
			review.Title,
			review.Content,
			review.Rating,
			review.PostedAt.Format(DATE_FORMAT))

		content += BREAK_LINE
	}

	updatedTime := time.Now()

	fileNameTemplate := os.Getenv("DIGEST_PATH") + "/%s-%s.md"

	fileName := fmt.Sprintf(fileNameTemplate,
		application.ApplicationId,
		strings.ReplaceAll(updatedTime.Format(DATE_FORMAT), "/", "-"))

	file, err := os.Create(fileName)

	if err != nil {
		log.Printf("Errored while creating file for application %s", application.ApplicationId)
		return err
	}

	_, err = file.WriteString(content)

	if err != nil {
		log.Printf("Errored while wrting content into file for application %s", application.ApplicationId)
		return err
	}

	log.Printf("Application digest created with file name of %s", fileName)

	application.LastDigest = updatedTime

	return nil

}

package cdn

import (
	"fmt"
	"log"
	"os"
	"path"

	dblocker "github.com/pawl0wski/technikinformatyk-backend/db_locker"
	"github.com/pawl0wski/technikinformatyk-backend/model"
)

func getQuestionsWithImages() []model.Question {
	dbLocker := dblocker.GetDBLockerInstance()
	questions, _ := dbLocker.GetQuestions()

	questionsToReturn := []model.Question{}
	for _, question := range questions {
		if question.HaveImage {
			questionsToReturn = append(questionsToReturn, question)
		}
	}
	return questionsToReturn
}

func RebuildCdn() {
	if IsCDNEnabled() {
		dbLocker := dblocker.GetDBLockerInstance()

		questions := getQuestionsWithImages()
		for _, question := range questions {

			filePath := path.Join(os.Getenv("CDN_PATH"), fmt.Sprintf("%s.png", question.Uuid))
			if _, err := os.Stat(filePath); err != nil {
				image := dbLocker.GetImage(question.Uuid)
				err := os.WriteFile(filePath, image, 0644)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
		log.Println("CDN rebuilded")
	}
}

package cdn

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/pawl0wski/technikinformatyk-backend/cache"
	"github.com/pawl0wski/technikinformatyk-backend/model"
)

func getQuestionsWithImages() []model.Question {
	cacheInstance := cache.GetCacheInstance()
	questions, _ := cacheInstance.GetQuestions()

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
		cacheInstance := cache.GetCacheInstance()

		questions := getQuestionsWithImages()
		for _, question := range questions {

			filePath := path.Join(os.Getenv("CDN_PATH"), fmt.Sprintf("%s.png", question.Uuid))
			if _, err := os.Stat(filePath); err != nil {
				image := cacheInstance.GetImage(question.Uuid)
				err := os.WriteFile(filePath, image, 0644)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
		log.Println("CDN rebuilded")
	}
}

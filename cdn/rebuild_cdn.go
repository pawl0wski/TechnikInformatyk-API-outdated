package cdn

import (
	"log"
	"os"
	"path"

	"github.com/jpawlowskii/TechnikInformatykBackend/cache"
	"github.com/jpawlowskii/TechnikInformatykBackend/structs"
)

func getQuestionsWithImages() []structs.Question {
	cacheInstance := cache.GetCacheInstance()
	questions := cacheInstance.GetQuestions()

	questionsToReturn := []structs.Question{}
	for _, question := range questions {
		if question.HaveImage {
			questionsToReturn = append(questionsToReturn, question)
		}
	}
	return questionsToReturn
}

func RebuildCdn() {
	cacheInstance := cache.GetCacheInstance()

	questions := getQuestionsWithImages()
	for _, question := range questions {

		filePath := path.Join(os.Getenv("CDN_PATH"), question.Uuid)
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

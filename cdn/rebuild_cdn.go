package cdn

import (
	"archive/tar"
	"bufio"
	"fmt"
	dblocker "github.com/pawl0wski/technikinformatyk-backend/db_locker"
	"github.com/pawl0wski/technikinformatyk-backend/model"
	"log"
	"os"
	"path"
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

func createTarFileWithAllImages(filePath string, database *dblocker.DBLocker) {
	questionWithImages := getQuestionsWithImages()

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	fileBuffer := bufio.NewWriter(file)
	tw := tar.NewWriter(fileBuffer)

	for _, questionWithImage := range questionWithImages {
		image := database.GetImage(questionWithImage.Uuid)
		hdr := &tar.Header{
			Name: fmt.Sprintf("%s.jpeg", questionWithImage.Uuid),
			Size: int64(len(image)),
		}
		err := tw.WriteHeader(hdr)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = tw.Write(image)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tw.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func RebuildCdn() {
	if IsCDNEnabled() {
		dbLocker := dblocker.GetDBLockerInstance()

		questions := getQuestionsWithImages()
		log.Println("Coping all images to cdn folder 📷")
		for _, question := range questions {

			filePath := path.Join(os.Getenv("CDN_PATH"), fmt.Sprintf("%s.jpg", question.Uuid))
			if _, err := os.Stat(filePath); err != nil {
				image := dbLocker.GetImage(question.Uuid)
				err := os.WriteFile(filePath, image, 0644)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
		log.Println("Creating Tar file with all question images 🗃")
		createTarFileWithAllImages(path.Join(os.Getenv("CDN_PATH"), "images.tar"), dbLocker)
		log.Println("CDN rebuilded 🏗")
	}
}

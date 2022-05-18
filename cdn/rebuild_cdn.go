package cdn

import (
	"bufio"
	"fmt"
	dblocker "github.com/pawl0wski/technikinformatyk-backend/db_locker"
	"log"
	"os"
	"path"
	"runtime"
)

func RebuildCdn() {
	if IsCDNEnabled() {
		dbLocker := dblocker.GetDBLockerInstance()

		questions := getQuestionsWithImages()
		log.Println("Coping all images to cdn folder 🖼")
		for _, question := range questions {

			filePath := path.Join(GetCDNPath(), fmt.Sprintf("%s.jpg", question.Uuid))
			if _, err := os.Stat(filePath); err != nil {
				image := dbLocker.GetImage(question.Uuid)
				err := os.WriteFile(filePath, image, 0644)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}

		log.Println("Creating images snapshot 📸")
		file, err := os.Create(path.Join(GetCDNPath(), "images.tar"))
		if err != nil {
			log.Fatalln(err)
		}
		fileBuffer := bufio.NewWriter(file)
		CreateImagesSnapshot(fileBuffer, dbLocker)

		log.Println("Running garbage collector 🗑️")
		runtime.GC()
		log.Println("CDN rebuilded 🏗")
	}
}

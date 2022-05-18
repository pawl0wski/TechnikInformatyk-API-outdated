package cdn

import (
	"archive/tar"
	"bufio"
	"fmt"
	dblocker "github.com/pawl0wski/technikinformatyk-backend/db_locker"
	"log"
)

func CreateImagesSnapshot(buffer *bufio.Writer, database *dblocker.DBLocker) {
	questionWithImages := getQuestionsWithImages()

	tw := tar.NewWriter(buffer)

	for _, questionWithImage := range questionWithImages {
		image := database.GetImage(questionWithImage.Uuid)
		hdr := &tar.Header{
			Name: fmt.Sprintf("%s.jpeg", questionWithImage.Uuid),
			Size: int64(len(image)),
			Mode: 0777,
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
	err := tw.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

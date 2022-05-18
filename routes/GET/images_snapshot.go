package GET

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"github.com/pawl0wski/technikinformatyk-backend/cdn"
	dblocker "github.com/pawl0wski/technikinformatyk-backend/db_locker"
	"log"
	"net/http"
	"os"
	"path"
)

func ImagesSnapshot(c *gin.Context) {
	if cdn.IsCDNEnabled() {
		cdnPath := cdn.GetCDNPath()
		if cdnPath[0] == '.' {
			cdnPath = cdnPath[1:]
		}
		cdnPath = path.Join(cdnPath, "images.tar")
		c.Redirect(http.StatusMovedPermanently, cdnPath)
	} else {
		tempFile, err := os.CreateTemp(os.TempDir(), "tmp-")
		if err != nil {
			log.Fatalln()
		}
		buffWriter := bufio.NewWriter(tempFile)
		dbLocker := dblocker.GetDBLockerInstance()
		cdn.CreateImagesSnapshot(buffWriter, dbLocker)
		c.File(tempFile.Name())
		err = os.Remove(tempFile.Name())
		if err != nil {
			log.Println(err)
		}
	}
}

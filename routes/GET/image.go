package GET

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/pawl0wski/technikinformatyk-backend/cdn"
	dblocker "github.com/pawl0wski/technikinformatyk-backend/db_locker"
)

func Image(c *gin.Context) {
	questionUuid := c.Param("questionUuid")
	if cdn.IsCDNEnabled() {
		cdnPath := os.Getenv("CDN_PATH")
		if cdnPath[0] == '.' {
			cdnPath = cdnPath[1:]
		}
		c.Redirect(http.StatusMovedPermanently, path.Join(cdnPath, fmt.Sprintf("%s.png", questionUuid)))
	} else {
		dbLocker := dblocker.GetDBLockerInstance()
		image := dbLocker.GetImage(questionUuid)

		if len(image) != 0 {
			c.Data(200, "image/png", image)
		} else {
			NotFound(c)
		}
	}

}

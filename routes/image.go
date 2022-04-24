package routes

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/pawl0wski/technikinformatyk-backend/cache"
	"github.com/pawl0wski/technikinformatyk-backend/cdn"
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
		cacheInstance := cache.GetCacheInstance()
		image := cacheInstance.GetImage(questionUuid)

		if len(image) != 0 {
			c.Data(200, "image/png", image)
		} else {
			NotFound(c)
		}
	}

}

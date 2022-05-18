package GET

import (
	"github.com/gin-gonic/gin"
	"github.com/pawl0wski/technikinformatyk-backend/cdn"
	"net/http"
	"path"
)

func ImagesSnapshot(c *gin.Context) {
	if cdn.IsCDNEnabled() {
		cdnPath := path.Join(cdn.GetCDNPath(), "images.tar")
		if cdnPath[0] == '.' {
			cdnPath = cdnPath[1:]
		}
		c.Redirect(http.StatusMovedPermanently, cdnPath)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "You need to enable CDN to get images snapshot."})
	}
}

package cdn

import (
	"os"
)

func GetCDNPath() string {
	return os.Getenv("CDN_PATH")
}

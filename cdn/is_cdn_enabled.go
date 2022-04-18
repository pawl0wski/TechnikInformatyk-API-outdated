package cdn

import (
	"os"
)

func IsCDNEnabled() bool {
	return os.Getenv("ENABLE_CDN") == "1"
}

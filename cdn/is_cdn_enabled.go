package cdn

import (
	"os"
	"strconv"
)

func IsCDNEnabled() bool {
	result, err := strconv.ParseBool(os.Getenv("ENABLE_CDN"))
	if err != nil {
		result = false
	}
	return result
}

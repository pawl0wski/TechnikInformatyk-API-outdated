package cache

import (
	"os"
)

func isCacheEnabled() bool {
	return os.Getenv("ENABLE_CACHE") == "1"
}

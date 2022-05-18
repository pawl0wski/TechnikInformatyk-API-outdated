package cdn

import (
	"os"
	"testing"
)

func TestGetCDNPath(t *testing.T) {
	testCDNPath := "./testpath"
	os.Setenv("CDN_PATH", testCDNPath)
	if GetCDNPath() != testCDNPath {
		t.Error("GetCDNPath returns bad path")
	}

}

package cdn

import (
	"os"
	"testing"
)

func TestIsCdnEnabledIfNot(t *testing.T) {
	env := os.Getenv("ENABLE_CDN")
	if env != "" {
		os.Setenv("ENABLE_CDN", "")
	}
	if IsCDNEnabled() {
		t.Error("IsCDNEnabled returns true if CDN is disabled")
	}
}

func TestIsCdnEnabledIfIsEnabled(t *testing.T) {
	os.Setenv("ENABLE_CDN", "1")
	if !IsCDNEnabled() {
		t.Error("IsCDNEnabled returns false if CDN is enabled")
	}
}

func TestIsCdnEnabledIfIs0(t *testing.T) {
	os.Setenv("ENABLE_CDN", "0")
	if IsCDNEnabled() {
		t.Error("IsCDNEnabled returns true if CDN is disabled")
	}
}

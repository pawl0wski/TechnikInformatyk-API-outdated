package cdn

import (
	"os"
	"strconv"
	"testing"
)

type IsCDNEnabledResults struct {
	env      string
	expected bool
}

var isCdnEnabledResults = []IsCDNEnabledResults{
	{"", false},
	{"0", false},
	{"1", true},
	{"true", true},
	{"11", false},
	{"-1", false},
}

func TestIsCdnEnabled(t *testing.T) {
	for _, test := range isCdnEnabledResults {
		os.Setenv("ENABLE_CDN", test.env)
		if result := IsCDNEnabled(); result != test.expected {
			t.Errorf("IsCDNEnabled returns %s if ENV is %s", strconv.FormatBool(result), test.env)
		}
	}

}

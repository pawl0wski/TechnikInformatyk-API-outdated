package db

import (
	"os"
	"testing"
)

func TestGetDBConnectionPath(t *testing.T) {
	os.Setenv("MYSQL_CONNECTION_PATH", "test")
	connectionPath := GetDbConnectionPath()
	if connectionPath != "test" {
		t.Error("GetDBConnectionPath() results incorrect path")
	}
	os.Unsetenv("MYSQL_CONNECTION_PATH")
}

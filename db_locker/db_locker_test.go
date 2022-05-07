package dblocker

import "testing"

func TestGetDbLocker(t *testing.T) {
	dbLockerFirstInstance := GetDBLockerInstance()
	dbLockerSecondInstance := GetDBLockerInstance()

	if dbLockerFirstInstance != dbLockerSecondInstance {
		t.Error("First instance and second instance are not the same")
	}

}

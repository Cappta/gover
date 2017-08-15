package gover

import (
	"errors"
	"fmt"
)

const (
	success = ""
)

var (
	errDataType               = errors.New("Version asserts can only be used with Version values")
	errInsufficientParameters = errors.New("Insufficient parameters for the assert")
)

func convertToVersion(actual interface{}, expected ...interface{}) (actualVersion *Version, expectedVersions []*Version, err error) {
	actualVersion, ok := actual.(*Version)
	if ok == false {
		return nil, nil, errDataType
	}

	expectedVersions = make([]*Version, len(expected))
	for index, value := range expected {
		expectedVersions[index], ok = value.(*Version)
		if ok == false {
			return nil, nil, errDataType
		}
	}
	return
}

func VersionShouldBeGreaterThan(actual interface{}, expected ...interface{}) string {
	actualVersion, expectedVersions, err := convertToVersion(actual, expected...)
	if err != nil {
		return err.Error()
	}
	if len(expectedVersions) < 0 {
		return errInsufficientParameters.Error()
	}

	if actualVersion.IsHigherThan(expectedVersions[0]) == false {
		return fmt.Sprintf("%v is not greater than %v", actualVersion, expectedVersions[0])
	}
	return success
}

func VersionShouldBeLessThan(actual interface{}, expected ...interface{}) string {
	actualVersion, expectedVersions, err := convertToVersion(actual, expected...)
	if err != nil {
		return err.Error()
	}
	if len(expectedVersions) < 0 {
		return errInsufficientParameters.Error()
	}

	if actualVersion.IsLowerThan(expectedVersions[0]) == false {
		return fmt.Sprintf("%v is not less than %v", actualVersion, expectedVersions[0])
	}
	return success
}

func VersionShouldEqual(actual interface{}, expected ...interface{}) string {
	actualVersion, expectedVersions, err := convertToVersion(actual, expected...)
	if err != nil {
		return err.Error()
	}
	if len(expectedVersions) < 0 {
		return errInsufficientParameters.Error()
	}

	if actualVersion.IsEqualTo(expectedVersions[0]) == false {
		return fmt.Sprintf("%v is not equal to %v", actualVersion, expectedVersions[0])
	}
	return success
}

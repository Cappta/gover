package gover

import (
	"strconv"
	"strings"
)

// Version represents a version
type Version struct {
	Values []int
}

const (
	major    = iota
	minor    = iota
	build    = iota
	revision = iota
)

const (
	separator = "."
)

// NewVersion creates a Version structure from the provided values
func NewVersion(values ...int) (version *Version) {
	return &Version{Values: values}
}

// NewVersionFromString creates a Version structure from a dot separated string
func NewVersionFromString(versionString string) (version *Version, err error) {
	version = NewVersion()
	err = version.Scan(versionString)
	if err != nil {
		version = nil
	}
	return
}

func (version *Version) String() string {
	stringValues := make([]string, len(version.Values))
	for index, value := range version.Values {
		stringValues[index] = strconv.Itoa(value)
	}
	return strings.Join(stringValues, separator)
}

// GetMajor retrieves the major value of the version
func (version *Version) GetMajor() int {
	return version.GetValueFromIndex(major)
}

// GetMinor retrieves the minor value of the version
func (version *Version) GetMinor() int {
	return version.GetValueFromIndex(minor)
}

// GetBuild retrieves the build value of the version
func (version *Version) GetBuild() int {
	return version.GetValueFromIndex(build)
}

// GetRevision retrieves the revision value of the version
func (version *Version) GetRevision() int {
	return version.GetValueFromIndex(revision)
}

// GetValueFromIndex returns the value from the specified index or 0 if it doesn't exist
func (version *Version) GetValueFromIndex(index int) int {
	if len(version.Values) > index {
		return version.Values[index]
	}
	return 0
}

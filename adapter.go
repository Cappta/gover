package gover

import (
	"database/sql/driver"
	"errors"
	"strconv"
	"strings"
)

// Scan implements the sql's Scanner interface
func (version *Version) Scan(src interface{}) (err error) {
	versionString, ok := src.(string)
	if version == nil {
		return errors.New("Cannot Scan to a nil pointer")
	}
	if ok == false {
		return errors.New("Expected a string")
	}
	splitString := strings.Split(versionString, separator)

	version.Values = make([]int, len(splitString))
	for index, value := range splitString {
		version.Values[index], err = strconv.Atoi(value)
		if err != nil {
			return errors.New("Provided string has an invalid format")
		}
	}
	return
}

// Value implements the driver's Valuer interface
func (version *Version) Value() (output driver.Value, err error) {
	if version != nil {
		output = version.String()
	}
	return
}

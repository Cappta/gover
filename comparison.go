package gover

import "github.com/Cappta/gohelpmath"

const (
	// Lower means the Left Version is lower than the one in the Right
	Lower = iota - 1
	// Equal means both versions are equal
	Equal = iota - 1
	// Higher means the Left Version is higher than the one in the Right
	Higher = iota - 1
)

// Compare retrieves the major value of the version
func (version *Version) Compare(targetVersion *Version) int {
	versionLength := len(version.Values)
	targetVersionLength := len(targetVersion.Values)
	highestLength := gohelpmath.IntMax(versionLength, targetVersionLength)

	for index := 0; index < highestLength; index++ {
		versionValue := version.GetValueFromIndex(index)
		targetVersionValue := targetVersion.GetValueFromIndex(index)

		if versionValue < targetVersionValue {
			return Lower
		} else if versionValue > targetVersionValue {
			return Higher
		}
	}
	return Equal
}

// IsLowerThan returns true if the current version is lower than provided version
func (version *Version) IsLowerThan(targetVersion *Version) bool {
	return version.Compare(targetVersion) == Lower
}

// IsEqualTo returns true if the current version is equal to the provided version
func (version *Version) IsEqualTo(targetVersion *Version) bool {
	return version.Compare(targetVersion) == Equal
}

// IsHigherThan returns true if the current version is higher than provided version
func (version *Version) IsHigherThan(targetVersion *Version) bool {
	return version.Compare(targetVersion) == Higher
}

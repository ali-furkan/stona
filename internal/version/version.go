package version

import "fmt"

func GetVersion() string {
	if Version == 0 {
		return "unknown version"
	}

	version := fmt.Sprintf("v%g", Version)

	if VersionName != "" {
		version = fmt.Sprintf("v%g-%s", Version, VersionName)
	}

	return version
}

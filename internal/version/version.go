// Package version provides the version information for the application.
package version

import (
	"runtime/debug"
)

// GetVersion gets the version information for the application.
func GetVersion() string {
	buildInfo, hasBuildInfo := debug.ReadBuildInfo()

	if !hasBuildInfo {
		return "unknown"
	}

	for _, info := range buildInfo.Settings {
		if info.Key == "vcs.revision" {
			return info.Value
		}
	}

	return "unknown"
}

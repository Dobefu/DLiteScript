package version

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	t.Parallel()

	version := GetVersion()

	if version == "" {
		t.Fatalf("expected version, got empty string")
	}
}

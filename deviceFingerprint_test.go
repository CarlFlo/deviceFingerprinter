package devicefingerprint

import (
	"testing"
)

func TestFingerprint(t *testing.T) {

	fingerprint := GetFingerprint()
	if len(fingerprint) == 0 {
		t.Error("Expected fingerprint, got nothing")
	}

	t.Error("Device Fingerprint:", fingerprint)
}

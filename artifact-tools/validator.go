package artifact

import (
	"errors"
	"strings"
)

// Validate checks structural integrity of an artifact.
// No semantic interpretation — structure only.
func Validate(a Artifact) error {
	if strings.TrimSpace(a.ID) == "" {
		return errors.New("artifact.ID must not be empty")
	}
	if strings.TrimSpace(a.StateRoot) == "" {
		return errors.New("artifact.StateRoot must not be empty")
	}
	if a.Height == 0 {
		return errors.New("artifact.Height must be greater than zero")
	}
	if a.Timestamp <= 0 {
		return errors.New("artifact.Timestamp must be a positive unix epoch")
	}
	return nil
}

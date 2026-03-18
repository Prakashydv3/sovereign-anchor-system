package artifact

import "fmt"

// Generate returns the hex-encoded SHA-256 hash for a given artifact.
func Generate(a Artifact) (string, error) {
	h, err := Hash(a)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h), nil
}

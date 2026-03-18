package nodeops

import (
	"errors"
	"net/http"
	"time"
)

// Check performs a liveness probe against the node RPC endpoint.
// Returns nil if the node is reachable and responding.
func Check(rpcURL string) error {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(rpcURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 500 {
		return errors.New("node returned server error")
	}
	return nil
}

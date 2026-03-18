package client

// Config holds the minimum required parameters to connect to L1.
type Config struct {
	RPCURL          string // L1 node RPC endpoint
	ContractAddress string // Deployed AnchorRegistry address
	PrivateKey      string // Signing key (never commit real values)
}

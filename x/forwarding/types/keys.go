package types

const (
	ModuleName        = "forwarding"
	StoreKey          = "forwarding"
	TransientStoreKey = "transient_forwarding"
)

var (
	NumOfAccountsPrefix   = []byte("num_of_accounts")
	NumOfForwardsPrefix   = []byte("num_of_forwards")
	TotalForwardedPrefix  = []byte("total_forwarded")
	PendingForwardsPrefix = []byte("pending_forwards")
)

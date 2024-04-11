package types

const (
	// ModuleName defines the module name
	ModuleName = "invoice"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_invoice"
)

var (
	ParamsKey = []byte("p_invoice")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

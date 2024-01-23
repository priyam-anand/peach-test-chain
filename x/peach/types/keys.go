package types

const (
	// ModuleName defines the module name
	ModuleName = "peach"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_peach"
)

var (
	ParamsKey = []byte("p_peach")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

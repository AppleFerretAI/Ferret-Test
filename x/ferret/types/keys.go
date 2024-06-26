package types

const (
	// ModuleName defines the module name
	ModuleName = "ferret"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ferret"
)

var (
	ParamsKey = []byte("p_ferret")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

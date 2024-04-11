package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// InvoiceKeyPrefix is the prefix to retrieve all Invoice
	InvoiceKeyPrefix = "Invoice/value/"
)

// InvoiceKey returns the store key to retrieve a Invoice from the index fields
func InvoiceKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

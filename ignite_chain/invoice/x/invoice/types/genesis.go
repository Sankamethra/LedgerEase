package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		InvoiceList: []Invoice{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in invoice
	invoiceIndexMap := make(map[string]struct{})

	for _, elem := range gs.InvoiceList {
		index := string(InvoiceKey(elem.Index))
		if _, ok := invoiceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for invoice")
		}
		invoiceIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

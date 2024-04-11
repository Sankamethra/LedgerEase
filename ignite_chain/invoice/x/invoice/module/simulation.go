package invoice

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"invoice/testutil/sample"
	invoicesimulation "invoice/x/invoice/simulation"
	"invoice/x/invoice/types"
)

// avoid unused import issue
var (
	_ = invoicesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgStoreinvoice = "op_weight_msg_storeinvoice"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStoreinvoice int = 100

	opWeightMsgCreateInvoice = "op_weight_msg_invoice"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateInvoice int = 100

	opWeightMsgUpdateInvoice = "op_weight_msg_invoice"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateInvoice int = 100

	opWeightMsgDeleteInvoice = "op_weight_msg_invoice"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteInvoice int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	invoiceGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		InvoiceList: []types.Invoice{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&invoiceGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgStoreinvoice int
	simState.AppParams.GetOrGenerate(opWeightMsgStoreinvoice, &weightMsgStoreinvoice, nil,
		func(_ *rand.Rand) {
			weightMsgStoreinvoice = defaultWeightMsgStoreinvoice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStoreinvoice,
		invoicesimulation.SimulateMsgStoreinvoice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateInvoice int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateInvoice, &weightMsgCreateInvoice, nil,
		func(_ *rand.Rand) {
			weightMsgCreateInvoice = defaultWeightMsgCreateInvoice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateInvoice,
		invoicesimulation.SimulateMsgCreateInvoice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateInvoice int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateInvoice, &weightMsgUpdateInvoice, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateInvoice = defaultWeightMsgUpdateInvoice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateInvoice,
		invoicesimulation.SimulateMsgUpdateInvoice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteInvoice int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteInvoice, &weightMsgDeleteInvoice, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteInvoice = defaultWeightMsgDeleteInvoice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteInvoice,
		invoicesimulation.SimulateMsgDeleteInvoice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgStoreinvoice,
			defaultWeightMsgStoreinvoice,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				invoicesimulation.SimulateMsgStoreinvoice(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateInvoice,
			defaultWeightMsgCreateInvoice,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				invoicesimulation.SimulateMsgCreateInvoice(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateInvoice,
			defaultWeightMsgUpdateInvoice,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				invoicesimulation.SimulateMsgUpdateInvoice(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteInvoice,
			defaultWeightMsgDeleteInvoice,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				invoicesimulation.SimulateMsgDeleteInvoice(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (

	// EpcBech32Prefix defines the Bech32 prefix of an account's address
	EpcBech32Prefix = "epc"

	// EpcCoinType Epc coin in https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	EpcCoinType = 118

	// EpcFullFundraiserPath BIP44Prefix is the parts of the BIP44 HD path that are fixed by what we used during the fundraiser.
	// use the registered cosmos stake token ATOM 118 as coin_type
	// m / purpose' / coin_type' / account' / change / address_index
	EpcFullFundraiserPath = "44'/118'/0'/0/0"

	// EpcBech32PrefixAccAddr defines the Bech32 prefix of an account's address
	EpcBech32PrefixAccAddr = EpcBech32Prefix
	// EpcBech32PrefixAccPub defines the Bech32 prefix of an account's public key
	EpcBech32PrefixAccPub = EpcBech32Prefix + sdk.PrefixPublic
	// EpcBech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	EpcBech32PrefixValAddr = EpcBech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator
	// EpcBech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	EpcBech32PrefixValPub = EpcBech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	// EpcBech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	EpcBech32PrefixConsAddr = EpcBech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus
	// EpcBech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	EpcBech32PrefixConsPub = EpcBech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)

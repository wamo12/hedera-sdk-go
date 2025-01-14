//go:build all || unit
// +build all unit

package hedera

import (
	"testing"

	"github.com/hashgraph/hedera-protobufs-go/services"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUnitAccountAllowanceAdjustTransaction(t *testing.T) {
	tokenID1 := TokenID{Token: 1}
	tokenID2 := TokenID{Token: 141}
	serialNumber1 := int64(3)
	serialNumber2 := int64(4)
	nftID1 := tokenID2.Nft(serialNumber1)
	nftID2 := tokenID2.Nft(serialNumber2)
	spenderAccountID1 := AccountID{Account: 7}
	spenderAccountID2 := AccountID{Account: 7890}
	nodeAccountID := []AccountID{{Account: 10}, {Account: 11}, {Account: 12}}
	hbarAmount := HbarFromTinybar(100)
	tokenAmount := int64(101)

	transactionID := TransactionIDGenerate(AccountID{Account: 324})

	transaction, err := NewAccountAllowanceAdjustTransaction().
		SetTransactionID(transactionID).
		SetNodeAccountIDs(nodeAccountID).
		AddHbarAllowance(spenderAccountID1, hbarAmount).
		AddTokenAllowance(tokenID1, spenderAccountID1, tokenAmount).
		AddTokenNftAllowance(nftID1, spenderAccountID1).
		AddTokenNftAllowance(nftID2, spenderAccountID1).
		AddTokenNftAllowance(nftID2, spenderAccountID2).
		AddAllTokenNftAllowance(tokenID1, spenderAccountID1).
		Freeze()
	require.NoError(t, err)

	data := transaction._Build()

	switch d := data.Data.(type) {
	case *services.TransactionBody_CryptoAdjustAllowance:
		require.Equal(t, d.CryptoAdjustAllowance.CryptoAllowances, []*services.CryptoAllowance{
			{
				Spender: spenderAccountID1._ToProtobuf(),
				Amount:  hbarAmount.AsTinybar(),
			},
		})
		require.Equal(t, d.CryptoAdjustAllowance.NftAllowances, []*services.NftAllowance{
			{
				TokenId:        tokenID2._ToProtobuf(),
				Spender:        spenderAccountID1._ToProtobuf(),
				SerialNumbers:  []int64{serialNumber1, serialNumber2},
				ApprovedForAll: &wrapperspb.BoolValue{Value: false},
			},
			{
				TokenId:        tokenID2._ToProtobuf(),
				Spender:        spenderAccountID2._ToProtobuf(),
				SerialNumbers:  []int64{serialNumber2},
				ApprovedForAll: &wrapperspb.BoolValue{Value: false},
			},
			{
				TokenId:        tokenID1._ToProtobuf(),
				Spender:        spenderAccountID1._ToProtobuf(),
				SerialNumbers:  []int64{},
				ApprovedForAll: &wrapperspb.BoolValue{Value: true},
			},
		})
		require.Equal(t, d.CryptoAdjustAllowance.TokenAllowances, []*services.TokenAllowance{
			{
				TokenId: tokenID1._ToProtobuf(),
				Spender: spenderAccountID1._ToProtobuf(),
				Amount:  tokenAmount,
			},
		})
	}
}

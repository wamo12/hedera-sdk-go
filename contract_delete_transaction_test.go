package hedera

import (
	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerializeContractDeleteTransaction(t *testing.T) {
	mockClient, err := newMockClient()
	assert.NoError(t, err)

	privateKey, err := Ed25519PrivateKeyFromString(mockPrivateKey)
	assert.NoError(t, err)

	tx, err := NewContractDeleteTransaction().
		SetContractID(ContractID{Contract: 5}).
		SetMaxTransactionFee(HbarFromTinybar(1e6)).
		SetTransactionID(testTransactionID).
		Build(mockClient)

	assert.NoError(t, err)

	tx.Sign(privateKey)

	cupaloy.SnapshotT(t, tx.String())
}

func TestSerializeContractDeleteTransaction_WithAccountIDObtainer(t *testing.T) {
	mockClient, err := newMockClient()
	assert.NoError(t, err)

	privateKey, err := Ed25519PrivateKeyFromString(mockPrivateKey)
	assert.NoError(t, err)

	tx, err := NewContractDeleteTransaction().
		SetContractID(ContractID{Contract: 5}).
		SetMaxTransactionFee(HbarFromTinybar(1e6)).
		SetTransferAccountID(AccountID{Account: 3}).
		SetTransactionID(testTransactionID).
		Build(mockClient)

	assert.NoError(t, err)

	tx.Sign(privateKey)

	cupaloy.SnapshotT(t, tx.String())
}

func TestSerializeContractDeleteTransaction_WithContractIDObtainer(t *testing.T) {
	mockClient, err := newMockClient()
	assert.NoError(t, err)

	privateKey, err := Ed25519PrivateKeyFromString(mockPrivateKey)
	assert.NoError(t, err)

	tx, err := NewContractDeleteTransaction().
		SetContractID(ContractID{Contract: 5}).
		SetMaxTransactionFee(HbarFromTinybar(1e6)).
		SetTransferContractID(ContractID{Contract: 3}).
		SetTransactionID(testTransactionID).
		Build(mockClient)

	assert.NoError(t, err)

	tx.Sign(privateKey)

	cupaloy.SnapshotT(t, tx.String())
}
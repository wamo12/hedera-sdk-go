package hedera

import (
	"github.com/hashgraph/hedera-sdk-go/proto"
)

type TransactionReceipt struct {
	Status                  Status
	ExchangeRate            *ExchangeRate
	TopicID                 *TopicID
	FileID                  *FileID
	ContractID              *ContractID
	AccountID               *AccountID
	TopicSequenceNumber     uint64
	TopicRunningHash        []byte
	TopicRunningHashVersion uint64
}

func newTransactionReceipt(
	status Status, exchangeRate *ExchangeRate,
	topicID TopicID, fileID FileID,
	contractID ContractID, accountID AccountID,
	topicSequenceNumber uint64, topicRunningHash []byte,
	topicRunningHashVersion uint64) TransactionReceipt {

	receipt := TransactionReceipt{
		Status:                  status,
		ExchangeRate:            exchangeRate,
		TopicID:                 &topicID,
		FileID:                  &fileID,
		ContractID:              &contractID,
		AccountID:               &accountID,
		TopicSequenceNumber:     topicSequenceNumber,
		TopicRunningHash:        topicRunningHash,
		TopicRunningHashVersion: topicRunningHashVersion,
	}

	return receipt

}

func transactionReceiptFromProtobuf(protoReceipt *proto.TransactionReceipt) TransactionReceipt {
	var accountID *AccountID
	if protoReceipt.AccountID != nil {
		accountIDValue := accountIDFromProto(protoReceipt.AccountID)
		accountID = &accountIDValue
	}

	var contractID *ContractID
	if protoReceipt.ContractID != nil {
		contractIDValue := contractIDFromProto(protoReceipt.ContractID)
		contractID = &contractIDValue
	}

	var fileID *FileID
	if protoReceipt.FileID != nil {
		fileIDValue := fileIDFromProto(protoReceipt.FileID)
		fileID = &fileIDValue
	}

	var topicID *TopicID
	if protoReceipt.TopicID != nil {
		topicIDValue := TopicIDFromProto(protoReceipt.TopicID)
		topicID = &topicIDValue
	}

	var rate *ExchangeRate
	if protoReceipt.ExchangeRate != nil {
		exchangeRateValue := exchangeRateFromProtobuf(protoReceipt.ExchangeRate.GetCurrentRate())
		rate = &exchangeRateValue
	}

	return TransactionReceipt{
		Status:                  Status(protoReceipt.Status),
		ExchangeRate:            rate,
		TopicID:                 topicID,
		FileID:                  fileID,
		ContractID:              contractID,
		AccountID:               accountID,
		TopicSequenceNumber:     protoReceipt.TopicSequenceNumber,
		TopicRunningHash:        protoReceipt.TopicRunningHash,
		TopicRunningHashVersion: protoReceipt.TopicRunningHashVersion,
	}
}

func (receipt TransactionReceipt) toProtobuf() *proto.TransactionReceipt {
	return &proto.TransactionReceipt{
		Status:     proto.ResponseCodeEnum(receipt.Status),
		AccountID:  receipt.AccountID.toProtobuf(),
		FileID:     receipt.FileID.toProto(),
		ContractID: receipt.ContractID.toProto(),
		ExchangeRate: &proto.ExchangeRateSet{
			CurrentRate: receipt.ExchangeRate.toProtobuf(),
			NextRate:    receipt.ExchangeRate.toProtobuf(),
		},
		TopicID:                 receipt.TopicID.toProto(),
		TopicSequenceNumber:     receipt.TopicSequenceNumber,
		TopicRunningHash:        receipt.TopicRunningHash,
		TopicRunningHashVersion: receipt.TopicRunningHashVersion,
		TokenId:                 nil,
	}
}
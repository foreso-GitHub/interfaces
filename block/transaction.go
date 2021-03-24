package block

import (
	"github.com/tokentransfer/interfaces/core"
	"github.com/tokentransfer/interfaces/crypto"
)

type TransactionType uint32

type StateType uint32

type TransactionResult uint32

type Transaction interface {
	core.Indexable
	crypto.Signable

	GetTransactionType() TransactionType

	GetDestination() core.Address
	GetSequence() uint64
	GetAmount() int64
	GetGas() int64
	GetType() string
}

type TransactionWithData interface {
	crypto.Hashable

	GetTransaction() Transaction
	GetReceipt() Receipt
}

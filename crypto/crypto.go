package crypto

import (
	"github.com/tokentransfer/interfaces/account"
	"github.com/tokentransfer/interfaces/core"
)

type Hashable interface {
	core.Binariable

	GetHash() core.Hash
	SetHash(h core.Hash)
}

type Signable interface {
	Hashable

	GetAccount() core.Address

	GetSignature() core.Signature
	SetSignature(s core.Signature)

	SetPublicKey(p core.PublicKey)
	GetPublicKey() core.PublicKey
	Raw() ([]byte, error)
}

type CryptoService interface {
	GetSize() int
	Hash(msg []byte) (core.Hash, error)
	Raw(h Hashable) (core.Hash, []byte, error)
	Sign(p account.Key, s Signable) error
	Verify(s Signable) (bool, error)
}

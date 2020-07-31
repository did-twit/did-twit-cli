package did

import (
	"crypto/ed25519"
	"errors"
	"fmt"
	"strings"

	"github.com/btcsuite/btcutil/base58"

	"github.com/did-twit/did-twit-cli/pkg"
)

const (
	// Base58BTC https://github.com/multiformats/go-multibase/blob/master/multibase.go
	Base58BTCMultiBase = "z"

	// ed25519-pub https://github.com/multiformats/multicodec/blob/master/table.csv
	Ed25519MultiCodec = 0xed

	// DID Twit prefix
	DIDPrefix = "did:twit"
)

func CreateDID(username string) (*string, ed25519.PrivateKey, error) {
	pubKey, privKey, err := pkg.GenerateEd25519Key()
	if err != nil {
		return nil, nil, err
	}
	multiCodec := append([]byte{Ed25519MultiCodec}, pubKey...)
	did := fmt.Sprintf("%s:%s:%s%s", DIDPrefix, username, Base58BTCMultiBase, base58.Encode(multiCodec))
	return &did, privKey, nil
}

func ExpandDID(did string) (ed25519.PublicKey, error) {
	split := strings.Split(did, ":")
	if len(split) != 4 {
		return nil, errors.New("malformed did:twit")
	}
	maybePrefix := fmt.Sprintf("%s:%s", split[0], split[1])
	if maybePrefix != DIDPrefix {
		return nil, errors.New("prefix is not did:twit")
	}
	if !strings.HasPrefix(split[3], Base58BTCMultiBase) {
		return nil, errors.New("unrecognized multi base prefix")
	}

	decoded := base58.Decode(split[3][1:])
	// check first byte is known coded
	if decoded[0] != Ed25519MultiCodec {
		return nil, errors.New("first byte does not contain known Ed25519 multi codec")
	}
	return decoded[1:], nil
}

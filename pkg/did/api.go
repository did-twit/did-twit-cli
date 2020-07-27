package did

import (
	"crypto/ed25519"
	"errors"

	"github.com/did-twitter/did-twitter-cli/internal/lib/did"
)

type DIDManagement interface {
	CreateDID(username string) (*did.SignedDIDDoc, error)
	UpdateDID(username string, currentDoc did.DIDDoc) (*did.SignedDIDDoc, error)
	DeactivateDID(username string, currentDoc did.DIDDoc) (*did.SignedDIDDoc, error)
}

type didTwit struct {
	username string
	privKey  ed25519.PrivateKey
}

func (d *didTwit) CreateDID() (*did.SignedDIDDoc, error) {
	if d.username == "" {
		return nil, errors.New("username cannot be empty")
	}
	doc, pk, err := did.GenerateSignedDIDDocument(d.username)
	if err != nil {
		return nil, err
	}
	d.privKey = pk
	return doc, nil
}

func (d *didTwit) UpdateDID(_ string, _ did.DIDDoc) (*did.SignedDIDDoc, error) {
	return nil, errors.New("not implemented")
}

func (d *didTwit) DeactivateDID(_ string, _ did.DIDDoc) (*did.SignedDIDDoc, error) {
	return nil, errors.New("not implemented")
}

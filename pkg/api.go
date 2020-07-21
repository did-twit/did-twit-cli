package pkg

import (
	"crypto/ed25519"
	"errors"

	"github.com/did-twitter/did-twitter-cli/internal/lib"
)

type DIDTwitter interface {
	CreateDID(username string) (*lib.SignedDIDDoc, error)
	UpdateDID(username string, currentDoc lib.DIDDoc) (*lib.SignedDIDDoc, error)
	DeactivateDID(username string, currentDoc lib.DIDDoc) (*lib.SignedDIDDoc, error)
}

type didTwitter struct {
	username string
	privKey  ed25519.PrivateKey
}

func (d *didTwitter) CreateDID() (*lib.SignedDIDDoc, error) {
	if d.username == "" {
		return nil, errors.New("username cannot be empty")
	}
	doc, pk, err := lib.GenerateSignedDIDDocument(d.username)
	if err != nil {
		return nil, err
	}
	d.privKey = pk
	return doc, nil
}

func (d *didTwitter) UpdateDID(username string, currentDoc lib.DIDDoc) (*lib.SignedDIDDoc, error) {
	return nil, nil
}

func (d *didTwitter) DeactivateDID(username string, currentDoc lib.DIDDoc) (*lib.SignedDIDDoc, error) {
	return nil, nil
}

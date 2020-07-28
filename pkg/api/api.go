package api

import (
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/btcsuite/btcutil/base58"

	"github.com/did-twit/did-twit-cli/internal/lib"
	"github.com/did-twit/did-twit-cli/internal/lib/did"
)

type DIDTwitAPI interface {
	CreateDIDTweet(username string) (*string, ed25519.PrivateKey, error)
	GenerateTweet(username, tweet string, privKey ed25519.PrivateKey) (*string, error)

	DIDAPI
	TweetAPI
}

type DIDAPI interface {
	CreateDID(username string) (*lib.SignedDIDDoc, ed25519.PrivateKey, error)
	RecoverDID(username string, key ed25519.PrivateKey) (*lib.SignedDIDDoc, error)
	AddKey(doc lib.DIDDoc, privKey ed25519.PrivateKey) (*lib.SignedDIDDoc, error)
	RemoveKey(doc lib.DIDDoc, keyID string, privKey ed25519.PrivateKey) (*lib.SignedDIDDoc, error)
	DeactivateDID(doc lib.DIDDoc, privKey ed25519.PrivateKey) (*lib.SignedDIDDoc, error)
}

type TweetAPI interface {
	PostTweet(tweet lib.Tweet) (*string, error)
	GetTweet(id string) (*lib.Tweet, error)
	DeleteTweet(id string) error
}

type didTwit struct{}

func (d *didTwit) CreateDID(username string) (*lib.SignedDIDDoc, ed25519.PrivateKey, error) {
	doc, privKey, err := did.GenerateSignedDIDDocument(username)
	if err != nil {
		return nil, nil, err
	}
	return doc, privKey, nil
}

func (d *didTwit) RecoverDID(username string, privKey ed25519.PrivateKey) (*lib.SignedDIDDoc, error) {
	return did.RecoverDIDDocument(username, privKey)
}

func (d *didTwit) AddKey(_ lib.DIDDoc, _ ed25519.PrivateKey) (*lib.SignedDIDDoc, error) {
	return nil, errors.New("not implemented")
}

func (d *didTwit) RemoveKey(_ lib.DIDDoc, _ string, _ ed25519.PrivateKey) (*lib.SignedDIDDoc, error) {
	return nil, errors.New("not implemented")
}

func (d *didTwit) DeactivateDID(doc lib.DIDDoc, privKey ed25519.PrivateKey) (*lib.SignedDIDDoc, error) {
	return did.DeactivateDIDDocument(doc, privKey)
}

func (d *didTwit) PostTweet(tweet lib.Tweet) (*string, error) {
	return nil, errors.New("not implemented")
}

func (d *didTwit) GetTweet(id string) (*lib.Tweet, error) {
	return nil, errors.New("not implemented")
}

func (d *didTwit) DeleteTweet(id string) error {
	return errors.New("not implemented")
}

func (d *didTwit) CreateDIDTweet(username string) (*string, ed25519.PrivateKey, error) {
	doc, privKey, err := d.CreateDID(username)
	if err != nil {
		return nil, nil, err
	}
	bytes, err := json.Marshal(doc)
	if err != nil {
		return nil, nil, err
	}
	tweet := fmt.Sprintf("%s?create=%s", doc.ID, base58.Encode(bytes))
	return &tweet, privKey, nil
}

func (d *didTwit) GenerateTweet(username, tweet string, privKey ed25519.PrivateKey) (*string, error) {

}
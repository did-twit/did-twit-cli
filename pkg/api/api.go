package api

import (
	"crypto/ed25519"
	"errors"

	"github.com/did-twit/did-twit-cli/internal/lib/did"

	"github.com/did-twit/did-twit-cli/internal/lib"
)

type DIDTwitAPI interface {
	DIDAPI
	TweetAPI
}

type DIDAPI interface {
	CreateDID(username string) (*string, ed25519.PrivateKey, error)
	ExpandDID(username string) (ed25519.PublicKey, error)
}

type TweetAPI interface {
	PostTweet(tweet lib.Tweet) (*string, error)
	GetTweet(id string) (*lib.Tweet, error)
	DeleteTweet(id string) error
}

type didTwit struct{}

func (d *didTwit) CreateDID(username string) (*string, ed25519.PrivateKey, error) {
	doc, privKey, err := did.CreateDID(username)
	if err != nil {
		return nil, nil, err
	}
	return doc, privKey, nil
}

func (d *didTwit) ExpandDID(username string) (ed25519.PublicKey, error) {
	return did.ExpandDID(username)
}

func (d *didTwit) PostTweet(_ lib.Tweet) (*string, error) {
	return nil, errors.New("not implemented")
}

func (d *didTwit) GetTweet(_ string) (*lib.Tweet, error) {
	return nil, errors.New("not implemented")
}

func (d *didTwit) DeleteTweet(_ string) error {
	return errors.New("not implemented")
}

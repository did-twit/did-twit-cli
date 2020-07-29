package api

import (
	"crypto/ed25519"
	"errors"

	"github.com/did-twit/did-twit-cli/internal"
	"github.com/did-twit/did-twit-cli/internal/did"
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
	PostTweet(tweet internal.Tweet) (*string, error)
	GetTweet(id string) (*internal.Tweet, error)
	DeleteTweet(id string) error
}

type didTwit struct{}

func (d *didTwit) CreateDID(username string) (*string, ed25519.PrivateKey, error) {
	return did.CreateDID(username)
}

func (d *didTwit) ExpandDID(username string) (ed25519.PublicKey, error) {
	return did.ExpandDID(username)
}

func (d *didTwit) PostTweet(_ internal.Tweet) (*string, error) {
	return nil, errors.New("not implemented")
}

func (d *didTwit) GetTweet(_ string) (*internal.Tweet, error) {
	return nil, errors.New("not implemented")
}

func (d *didTwit) DeleteTweet(_ string) error {
	return errors.New("not implemented")
}

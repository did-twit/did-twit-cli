package api

import (
	"testing"

	"github.com/btcsuite/btcutil/base58"
	"github.com/stretchr/testify/assert"

	"github.com/did-twit/did-twit-cli/internal/lib/did"
	"github.com/did-twit/did-twit-cli/pkg/tweet"
)

var (
	priv = "ttVFGrTDz922rCVTF9DFh1UkGZco1miUbvwkLmaK59Qa6bKAKavau6xK7eVHqAgrttyUR5vxjR913UKfJgzZXvZ"
)

func TestRecoverDIDTweet(t *testing.T) {
	doc, err := did.RecoverDIDDocument("didtwitt3r", base58.Decode(priv))
	assert.NoError(t, err)
	assert.NotEmpty(t, doc)

	tweet, err := tweet.CreateDIDTweet(*doc)
	assert.NoError(t, err)
	assert.NotEmpty(t, tweet)
}

func TestCreateDIDTweet(t *testing.T) {
	d := didTwit{}
	doc, privKey, err := d.CreateDID("test")
	assert.NoError(t, err)
	assert.NotEmpty(t, doc)
	assert.NotEmpty(t, privKey)

	tweet, err := tweet.CreateDIDTweet(*doc)
	assert.NoError(t, err)
	assert.NotEmpty(t, tweet)
}

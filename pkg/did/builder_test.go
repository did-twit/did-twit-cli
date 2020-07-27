package did

import (
	"testing"

	"github.com/btcsuite/btcutil/base58"
	"github.com/stretchr/testify/assert"

	"github.com/did-twitter/did-twitter-cli/internal/lib/did"
)

var (
	priv = "ttVFGrTDz922rCVTF9DFh1UkGZco1miUbvwkLmaK59Qa6bKAKavau6xK7eVHqAgrttyUR5vxjR913UKfJgzZXvZ"
)

func TestRecoverDIDTweet(t *testing.T) {
	doc, err := did.RecoverDIDDocument("didtwitt3r", base58.Decode(priv))
	assert.NoError(t, err)
	assert.NotEmpty(t, doc)

	tweet, err := CreateDIDTweet(*doc)
	assert.NoError(t, err)
	assert.NotEmpty(t, tweet)
}

func TestCreateDIDTweet(t *testing.T) {
	d := didTwit{username: "test"}
	doc, err := d.CreateDID()
	assert.NoError(t, err)
	assert.NotEmpty(t, doc)

	tweet, err := CreateDIDTweet(*doc)
	assert.NoError(t, err)
	assert.NotEmpty(t, tweet)
}
package pkg

import (
	"testing"

	"github.com/btcsuite/btcutil/base58"
	"github.com/stretchr/testify/assert"

	"github.com/did-twitter/did-twitter-cli/internal/lib"
)

var (
	priv = "ttVFGrTDz922rCVTF9DFh1UkGZco1miUbvwkLmaK59Qa6bKAKavau6xK7eVHqAgrttyUR5vxjR913UKfJgzZXvZ"
)

func TestRecoverDIDTweet(t *testing.T) {
	doc, err := lib.RecoverDIDDocument("didtwitt3r", base58.Decode(priv))
	assert.NoError(t, err)
	assert.NotEmpty(t, doc)

	tweet, err := CreateDIDTweet(*doc)
	assert.NoError(t, err)
	println(*tweet)
}

func TestCreateDIDTweet(t *testing.T) {
	d := didTwitter{username: "test"}
	doc, err := d.CreateDID()
	assert.NoError(t, err)
	assert.NotEmpty(t, doc)

	println(base58.Encode(d.privKey))

	tweet, err := CreateDIDTweet(*doc)
	assert.NoError(t, err)
	assert.NotEmpty(t, tweet)

	println(*tweet)
	println(len(*tweet))
}
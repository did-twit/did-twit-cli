package api

import (
	"crypto/ed25519"
	"testing"

	"github.com/stretchr/testify/assert"

	tweetlib "github.com/did-twit/did-twit-cli/internal/lib/tweet"
)

func TestCreateDIDTweet(t *testing.T) {
	d := didTwit{}

	tweet, doc, privKey, err := d.CreateDIDTweet("test")
	assert.NoError(t, err)
	assert.NotEmpty(t, tweet)
	assert.NotEmpty(t, doc)
	assert.NotEmpty(t, privKey)
}

func TestViewDIDTweet(t *testing.T) {
	d := didTwit{}

	createTweet, doc, privKey, err := d.CreateDIDTweet("test")
	assert.NoError(t, err)
	assert.NotEmpty(t, createTweet)
	assert.NotEmpty(t, doc)
	assert.NotEmpty(t, privKey)

	recoveredDoc, err := d.ViewDIDTweetDID(*createTweet)
	assert.NoError(t, err)

	assert.Equal(t, doc, recoveredDoc)
}

func TestGenerateTweet(t *testing.T) {
	d := didTwit{}
	_, doc, privKey, err := d.CreateDIDTweet("test")
	assert.NoError(t, err)

	tweetText := "test tweet"
	tweet, err := d.GenerateTweet(doc.VerificationMethods[0].ID, tweetText, privKey)
	assert.NoError(t, err)

	reconstructTweet, err := tweetlib.ReconstructTweet(*tweet)
	assert.NoError(t, err)

	pubKey := privKey.Public().(ed25519.PublicKey)
	err = tweetlib.VerifyTweet(*reconstructTweet, pubKey)
	assert.NoError(t, err)
	
	assert.Equal(t, tweetText, reconstructTweet.Tweet)
}

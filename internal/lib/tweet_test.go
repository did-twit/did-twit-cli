package lib

import (
	"crypto/ed25519"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	priv = "ttVFGrTDz922rCVTF9DFh1UkGZco1miUbvwkLmaK59Qa6bKAKavau6xK7eVHqAgrttyUR5vxjR913UKfJgzZXvZ"
)

func TestTweet(t *testing.T) {
	// Generate DID Doc
	doc, pk, err := GenerateSignedDIDDocument("didtwitt3r")
	assert.NoError(t, err)

	tweet, err := SignTweet(pk, doc.VerificationMethods[0].ID, "welcome to did:twitter")
	assert.NoError(t, err)

	err = VerifyTweet(pk.Public().(ed25519.PublicKey), *tweet)
	assert.NoError(t, err)

	tweetString, err := TweetToPost(*tweet)
	assert.NoError(t, err)

	println(*tweetString)
	println(len(*tweetString))
}

func TweetOfSizeN(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

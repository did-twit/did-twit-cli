package tweet

import (
	"crypto/ed25519"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/did-twitter/did-twitter-cli/internal/lib/did"
)

func TestTweet(t *testing.T) {
	// Generate DID Doc
	doc, pk, err := did.GenerateSignedDIDDocument("didtwitt3r")
	assert.NoError(t, err)

	tweet, err := SignTweet(pk, doc.VerificationMethods[0].ID, "welcome to did:twit")
	assert.NoError(t, err)

	err = VerifyTweet(pk.Public().(ed25519.PublicKey), *tweet)
	assert.NoError(t, err)

	tweetString, err := GenerateTweet(*tweet)
	assert.NoError(t, err)
	assert.NotEmpty(t, tweetString)
}

// Utility to return a tweet of size n characters
func TweetOfSizeN(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
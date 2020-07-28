package tweet

import (
	"crypto/ed25519"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/did-twit/did-twit-cli/internal/lib/did"
)

func TestSignGenerateAndVerifyTweet(t *testing.T) {
	// Generate DID Doc
	doc, pk, err := did.GenerateSignedDIDDocument("didtwitt3r")
	assert.NoError(t, err)

	tweet, err := SignTweet(pk, doc.VerificationMethods[0].ID, "welcome to api:twit")
	assert.NoError(t, err)

	err = VerifyTweet(*tweet, pk.Public().(ed25519.PublicKey))
	assert.NoError(t, err)

	tweetString, err := GenerateTweet(*tweet)
	assert.NoError(t, err)
	assert.NotEmpty(t, tweetString)
}

func TestReconstructTweet(t *testing.T) {
	// Generate DID Doc
	doc, pk, err := did.GenerateSignedDIDDocument("didtwitt3r")
	assert.NoError(t, err)

	tweet, err := SignTweet(pk, doc.VerificationMethods[0].ID, "welcome to api:twit")
	assert.NoError(t, err)

	err = VerifyTweet(*tweet, pk.Public().(ed25519.PublicKey))
	assert.NoError(t, err)

	tweetString, err := GenerateTweet(*tweet)
	assert.NoError(t, err)
	assert.NotEmpty(t, tweetString)

	// Reconstruct
	reconstructed, err := ReconstructTweet(*tweetString)
	assert.NoError(t, err)
	assert.Equal(t, tweet, reconstructed)
}

// Utility to return a tweet of size n characters
func tweetOfSizeN(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

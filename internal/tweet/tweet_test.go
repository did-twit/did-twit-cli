package tweet

import (
	"crypto/ed25519"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/did-twit/did-twit-cli/internal/did"
)

func TestSignGenerateAndVerifyTweet(t *testing.T) {
	// Generate DID Doc
	id, privKey, err := did.CreateDID("didtwitt3r")
	assert.NoError(t, err)

	tweet, err := SignTweet(privKey, *id, "hello world")
	assert.NoError(t, err)

	err = VerifyTweet(*tweet, privKey.Public().(ed25519.PublicKey))
	assert.NoError(t, err)

	tweetString, err := GenerateTweetText(*tweet)
	assert.NoError(t, err)
	assert.NotEmpty(t, tweetString)

	fmt.Printf("Tweet is: %s\n", *tweetString)
	fmt.Printf("Tweet length: %d\n", len(*tweetString))
}

func TestReconstructTweet(t *testing.T) {
	// Generate DID
	id, privKey, err := did.CreateDID("didtwit")
	assert.NoError(t, err)

	t.Run("happy path", func(t *testing.T) {
		tweet, err := SignTweet(privKey, *id, "welcome to did:twit")
		assert.NoError(t, err)

		err = VerifyTweet(*tweet, privKey.Public().(ed25519.PublicKey))
		assert.NoError(t, err)

		tweetString, err := GenerateTweetText(*tweet)
		assert.NoError(t, err)
		assert.NotEmpty(t, tweetString)

		// Reconstruct
		reconstructed, err := ReconstructTweet(*tweetString)
		assert.NoError(t, err)
		assert.Equal(t, tweet, reconstructed)
	})

	t.Run("multi dots path", func(t *testing.T) {
		tweet, err := SignTweet(privKey, *id, "welcome.to.did:twit")
		assert.NoError(t, err)

		err = VerifyTweet(*tweet, privKey.Public().(ed25519.PublicKey))
		assert.NoError(t, err)

		tweetString, err := GenerateTweetText(*tweet)
		assert.NoError(t, err)
		assert.NotEmpty(t, tweetString)

		// Reconstruct
		reconstructed, err := ReconstructTweet(*tweetString)
		assert.NoError(t, err)
		assert.Equal(t, tweet, reconstructed)
	})
}

// Tweets grow linearly in size
// 1 character = 438 chars = 2 tweets
// 240 characters = 677 chars = 3 tweets
func TestTweetSize(t *testing.T) {
	// Generate DID
	id, privKey, err := did.CreateDID("didtwitt3r")
	assert.NoError(t, err)

	for i := 1; i <= 240; i++ {
		tweet := tweetOfSizeN(i)
		signedTweet, err := SignTweet(privKey, *id, tweet)
		assert.NoError(t, err)

		tweetString, err := GenerateTweetText(*signedTweet)
		assert.NoError(t, err)

		fmt.Printf("Tweet of <%d>char size<%d>\n", i, len(*tweetString))
	}
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

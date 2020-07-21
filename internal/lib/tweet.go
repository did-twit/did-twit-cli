package lib

import (
	"crypto/ed25519"
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

func SignTweet(privKey ed25519.PrivateKey, verificationMethod, tweet string) (*Tweet, error) {
	t := Tweet{Tweet: tweet}
	tBytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	proof, err := GenerateProof(tBytes, privKey, verificationMethod)
	if err != nil {
		return nil, err
	}
	t.Proof = *proof
	return &t, nil
}

func VerifyTweet(key ed25519.PublicKey, tweet Tweet) error {
	bytes, err := json.Marshal(tweet)
	if err != nil {
		return err
	}
	return VerifyProof(bytes, key, tweet.Proof)
}

func TweetToPost(tweet Tweet) (*string, error) {
	bytes, err := json.Marshal(tweet.Proof)
	if err != nil {
		return nil, err
	}
	t := fmt.Sprintf("%s.%s", tweet.Tweet, base58.Encode(bytes))
	return &t, nil
}
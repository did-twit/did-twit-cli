package tweet

import (
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/btcsuite/btcutil/base58"

	"github.com/did-twit/did-twit-cli/pkg"
	"github.com/did-twit/did-twit-cli/pkg/crypto"
)

func SignTweet(privKey ed25519.PrivateKey, didTwit, tweet string) (*pkg.Tweet, error) {
	t := pkg.Tweet{Tweet: tweet}
	tBytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	proof, err := crypto.GenerateProof(tBytes, privKey, didTwit)
	if err != nil {
		return nil, err
	}
	t.Proof = *proof
	return &t, nil
}

func VerifyTweet(tweet pkg.Tweet, pubKey ed25519.PublicKey) error {
	bytes, err := json.Marshal(tweet)
	if err != nil {
		return err
	}
	return crypto.VerifyProof(bytes, pubKey, tweet.Proof)
}

func GenerateTweetText(tweet pkg.Tweet) (*string, error) {
	proofBytes, err := json.Marshal(tweet.Proof)
	if err != nil {
		return nil, err
	}
	t := fmt.Sprintf("%s.%s", tweet.Tweet, base58.Encode(proofBytes))
	return &t, nil
}

// ReconstructDIDDocument given a did:twit tweet, re-construct the DID Document
func ReconstructTweet(tweet string) (*pkg.Tweet, error) {
	split := strings.Split(tweet, ".")
	p := split[len(split)-1]
	t := strings.Join(split[:len(split)-1], ".")
	if len(split) < 2 {
		return nil, errors.New("malformed tweet")
	}
	proofBytes := base58.Decode(p)
	var proof pkg.Proof
	if err := json.Unmarshal(proofBytes, &proof); err != nil {
		return nil, err
	}
	return &pkg.Tweet{
		Tweet: t,
		Proof: proof,
	}, nil
}

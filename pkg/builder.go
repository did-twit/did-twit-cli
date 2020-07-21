package pkg

import (
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/base58"

	"github.com/did-twitter/did-twitter-cli/internal/lib"
)

func CreateDIDTweet(doc lib.SignedDIDDoc) (*string, error) {
	bytes, err := json.Marshal(doc)
	if err != nil {
		return nil, err
	}
	tweet := fmt.Sprintf("%s?create=%s", doc.ID, base58.Encode(bytes))
	return &tweet, nil
}
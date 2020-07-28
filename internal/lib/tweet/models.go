package tweet

import (
	"github.com/did-twitter/did-twitter-cli/internal/lib/did"
)

type Tweet struct {
	Tweet string    `json:"tweet"`
	Proof did.Proof `json:"proof"`
}

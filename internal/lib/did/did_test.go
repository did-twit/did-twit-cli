package did

import (
	"crypto/ed25519"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndExpandDID(t *testing.T) {
	// create
	did, privKey, err := CreateDID("god")
	assert.NoError(t, err)
	assert.NotEmpty(t, did)
	assert.NotEmpty(t, privKey)

	fmt.Printf("did is: %s\n", *did)
	fmt.Printf("length of did is: %d\n", len(*did))

	// expand
	pubKey, err := ExpandDID(*did)
	assert.NoError(t, err)
	assert.NotEmpty(t, pubKey)

	assert.Equal(t, privKey.Public().(ed25519.PublicKey), pubKey)
}

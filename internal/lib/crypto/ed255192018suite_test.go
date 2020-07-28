package crypto

import (
	"crypto/ed25519"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/did-twit/did-twit-cli/internal/lib"
)

func TestCanonicalize(t *testing.T) {
	doc := map[string]interface{}{
		"@context": map[string]interface{}{
			"ex": "http://example.org/vocab#",
		},
		"@id":   "http://example.org/test#example",
		"@type": "ex:Foo",
		"ex:embed": map[string]interface{}{
			"@type": "ex:Bar",
		},
	}
	bytes, err := json.Marshal(doc)
	assert.NoError(t, err)

	out, err := Canonicalize(bytes)
	assert.NoError(t, err)
	assert.NotEmpty(t, out)
}

func TestGenerateAndVerifyProof(t *testing.T) {
	// Generate
	input := []byte(`{"test":"data"}`)

	pubKey, privKey, err := ed25519.GenerateKey(nil)
	assert.NoError(t, err)

	verificationMethodRef := lib.KeyFragment("did:twit:test", lib.FirstKey)

	proof, err := GenerateProof(input, privKey, verificationMethodRef)
	assert.NoError(t, err)
	assert.NotEmpty(t, proof)
	assert.Equal(t, verificationMethodRef, proof.VerificationMethod)

	// Verify
	err = VerifyProof(input, pubKey, *proof)
	assert.NoError(t, err)
}

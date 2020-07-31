package crypto

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/did-twit/did-twit-cli/pkg/did"
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

	id, privKey, err := did.CreateDID("test")

	proof, err := GenerateProof(input, privKey, *id)
	assert.NoError(t, err)
	assert.NotEmpty(t, proof)
	assert.Equal(t, *id, proof.VerificationMethod)

	// Recover pub lib
	pubKey, err := did.ExpandDID(*id)
	assert.NoError(t, err)

	// Verify
	err = VerifyProof(input, pubKey, *proof)
	assert.NoError(t, err)
}

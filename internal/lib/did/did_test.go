package did

import (
	"crypto/ed25519"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDIDDocument(t *testing.T) {
	pubKey, privKey, err := ed25519.GenerateKey(nil)
	assert.NoError(t, err)

	// Generate
	doc := GenerateDIDDocumentWithKey("did:twit:test", pubKey)

	// Sign
	signed, err := SignDIDDocument(doc, privKey)
	assert.NoError(t, err)

	// Verify
	err = VerifyDIDDocument(*signed, pubKey)
	assert.NoError(t, err)
}

func TestGenerateSignedDIDDocument(t *testing.T) {
	// Generate
	doc, privKey, err := GenerateSignedDIDDocument("did:twit:test")
	assert.NoError(t, err)

	// Verify
	err = VerifyDIDDocument(*doc, privKey.Public().(ed25519.PublicKey))
	assert.NoError(t, err)
}

func TestDeactivateDIDDocument(t *testing.T) {
	doc, privKey, err := GenerateDIDDocument("did:twit:test")
	assert.NoError(t, err)

	// Make sure there's a key
	assert.NotEmpty(t, doc.VerificationMethods)

	// Create a proof for deactivation
	signedDoc, err := SignDIDDocument(*doc, privKey)
	assert.NoError(t, err)

	deactivated, err := DeactivateDIDDocument(signedDoc.DIDDoc, *signedDoc.Proof)
	assert.NoError(t, err)

	assert.NotEmpty(t, deactivated)
	assert.NotEmpty(t, deactivated.Updated)
	assert.Empty(t, deactivated.VerificationMethods)
	assert.Equal(t, doc.Created, deactivated.Created)
}
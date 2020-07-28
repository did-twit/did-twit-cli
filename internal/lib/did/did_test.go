package did

import (
	"crypto/ed25519"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDIDDocument(t *testing.T) {
	testUsername := "test"
	doc, privKey, err := GenerateDIDDocument(testUsername)
	assert.NoError(t, err)

	pubKey := privKey.Public().(ed25519.PublicKey)

	// Sign
	signed, err := SignDIDDocument(*doc, privKey)
	assert.NoError(t, err)

	// Verify
	err = VerifyDIDDocument(*signed, pubKey)
	assert.NoError(t, err)

	// Generate
	docWithKey := GenerateDIDDocumentWithKey(testUsername, pubKey)

	// Sign
	signed, err = SignDIDDocument(docWithKey, privKey)
	assert.NoError(t, err)

	// Verify
	err = VerifyDIDDocument(*signed, pubKey)
	assert.NoError(t, err)
}

func TestGenerateSignedDIDDocument(t *testing.T) {
	// Generate
	doc, privKey, err := GenerateSignedDIDDocument("test")
	assert.NoError(t, err)

	// Verify
	err = VerifyDIDDocument(*doc, privKey.Public().(ed25519.PublicKey))
	assert.NoError(t, err)
}

func TestRecoverDIDDocument(t *testing.T) {
	// Generate
	doc, privKey, err := GenerateSignedDIDDocument("test")
	assert.NoError(t, err)

	// Recover
	recovered, err := RecoverDIDDocument("test", privKey)
	assert.NoError(t, err)

	assert.Equal(t, doc.ID, recovered.ID)
	assert.Equal(t, doc.VerificationMethods[0].ID, recovered.VerificationMethods[0].ID)
	assert.Equal(t, doc.VerificationMethods[0].PublicKeyBase58, recovered.VerificationMethods[0].PublicKeyBase58)
	assert.Equal(t, doc.VerificationMethods[0].Controller, recovered.VerificationMethods[0].Controller)
	assert.Equal(t, doc.VerificationMethods[0].Type, recovered.VerificationMethods[0].Type)
	assert.Equal(t, doc.Authentication[0], recovered.Authentication[0])

}

func TestFindKeyAndVerifyDIDDocument(t *testing.T) {
	// Generate
	doc, _, err := GenerateSignedDIDDocument("test")
	assert.NoError(t, err)

	err = FindKeyAndVerifyDIDDocument(*doc)
	assert.NoError(t, err)
}

func TestDeactivateDIDDocument(t *testing.T) {
	doc, privKey, err := GenerateDIDDocument("test")
	assert.NoError(t, err)

	// Make sure there's a key
	assert.NotEmpty(t, doc.VerificationMethods)

	// Create a proof for deactivation
	signedDoc, err := SignDIDDocument(*doc, privKey)
	assert.NoError(t, err)

	deactivated, err := DeactivateDIDDocument(signedDoc.DIDDoc, privKey)
	assert.NoError(t, err)

	// Validate signature on deactivated with pub key
	err = VerifyDIDDocument(*deactivated, privKey.Public().(ed25519.PublicKey))
	assert.NoError(t, err)

	assert.NotEmpty(t, deactivated)
	assert.NotEmpty(t, deactivated.Updated)
	assert.Empty(t, deactivated.VerificationMethods)
	assert.Equal(t, doc.Created, deactivated.Created)
}

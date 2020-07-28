package did

import (
	"bytes"
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/btcsuite/btcutil/base58"

	"github.com/did-twit/did-twit-cli/internal/lib"
	"github.com/did-twit/did-twit-cli/internal/lib/crypto"
)

// GenerateDIDDocument generates a new DID document and key-pair for a provided lib name
func GenerateDIDDocument(name string) (*lib.DIDDoc, ed25519.PrivateKey, error) {
	pub, priv, err := lib.GenerateEd25519Key()
	if err != nil {
		return nil, nil, err
	}
	doc := GenerateDIDDocumentWithKey(name, pub)
	return &doc, priv, nil
}

// GenerateDIDDocumentWithKey generates a new, unsigned, DID document for the given username and public key
// Presently, this method does not support service endpoints or multiple keys
func GenerateDIDDocumentWithKey(username string, key ed25519.PublicKey) lib.DIDDoc {
	did := fmt.Sprintf("%s:%s", lib.DIDPrefix, username)
	didWithFragment := lib.KeyFragment(did, lib.FirstKey)
	b58PubKey := base58.Encode(key)
	verificationMethod := lib.VerificationMethod{
		ID:              didWithFragment,
		Type:            crypto.KeyType,
		Controller:      did,
		PublicKeyBase58: b58PubKey,
	}
	return lib.DIDDoc{
		ID:                  did,
		VerificationMethods: []lib.VerificationMethod{verificationMethod},
		Authentication:      []string{didWithFragment},
		Created:             time.Now().Format(time.RFC3339),
	}
}

// GenerateSignedDIDDocument creates a key and signs a new DID Document for a provided username
func GenerateSignedDIDDocument(username string) (*lib.SignedDIDDoc, ed25519.PrivateKey, error) {
	pubKey, privKey, err := lib.GenerateEd25519Key()
	if err != nil {
		return nil, nil, err
	}
	doc := GenerateDIDDocumentWithKey(username, pubKey)
	signedDoc, err := SignDIDDocument(doc, privKey)
	return signedDoc, privKey, err
}

func RecoverDIDDocument(username string, key ed25519.PrivateKey) (*lib.SignedDIDDoc, error) {
	pubKey := key.Public().(ed25519.PublicKey)
	doc := GenerateDIDDocumentWithKey(username, pubKey)
	return SignDIDDocument(doc, key)
}

// SignDIDDocument takes an unsigned DID Document and signs it with the given key, returning a new object that wraps
// the document and a proof. This method verifies that the signing key's public key is contained in the document.
func SignDIDDocument(doc lib.DIDDoc, key ed25519.PrivateKey) (*lib.SignedDIDDoc, error) {
	// Make sure the pub key is in the lib document
	pubKey := key.Public().(ed25519.PublicKey)
	verificationMethod, err := findMyKey(pubKey, doc.VerificationMethods)
	if err != nil {
		return nil, err
	}

	// Turn the doc into bytes for proof generation
	docBytes, err := json.Marshal(doc)
	if err != nil {
		return nil, err
	}

	// Get the proof
	proof, err := crypto.GenerateProof(docBytes, key, verificationMethod.ID)
	if err != nil {
		return nil, err
	}

	// Build response
	return &lib.SignedDIDDoc{
		DIDDoc: doc,
		Proof:  proof,
	}, nil
}

// VerifyDIDDocument takes a signed DID Document and verifies it using the Ed25519 2018 Linked Data Suite Verification
func VerifyDIDDocument(doc lib.SignedDIDDoc, key ed25519.PublicKey) error {
	docBytes, err := json.Marshal(doc.DIDDoc)
	if err != nil {
		return err
	}
	return crypto.VerifyProof(docBytes, key, *doc.Proof)
}

// FindKeyAndVerifyDIDDocument tries to verify the signature of the doc with the key in the verification method of the proof
func FindKeyAndVerifyDIDDocument(doc lib.SignedDIDDoc) error {
	method, err := findMyVerificationMethod(doc.Proof.VerificationMethod, doc.VerificationMethods)
	if err != nil {
		return err
	}

	pubKey := base58.Decode(method.PublicKeyBase58)
	if err := VerifyDIDDocument(doc, pubKey); err != nil {
		return err
	}

	return nil
}

// Takes in a (current) DID Document, and a signed updated DID Document and validates that
// the update.
//func AddKeyToDIDDoc(currentDoc DIDDoc, updatedDoc SignedDIDDoc) error {
//
//}

// Takes in a DID Document and a key in the doc that can be used to author its deactivation
func DeactivateDIDDocument(doc lib.DIDDoc, privKey ed25519.PrivateKey) (*lib.SignedDIDDoc, error) {
	// make sure the pub key is in the doc
	verificationMethod, err := findMyKey(privKey.Public().(ed25519.PublicKey), doc.VerificationMethods)
	if err != nil {
		return nil, err
	}

	deactivated := lib.DIDDoc{
		ID:      doc.ID,
		Created: doc.Created,
		Updated: time.Now().Format(time.RFC3339),
	}

	// Turn the doc into bytes for proof generation
	docBytes, err := json.Marshal(doc)
	if err != nil {
		return nil, err
	}

	// Get the proof
	proof, err := crypto.GenerateProof(docBytes, privKey, verificationMethod.ID)
	if err != nil {
		return nil, err
	}

	return &lib.SignedDIDDoc{
		DIDDoc: deactivated,
		Proof:  proof,
	}, nil
}

func findMyVerificationMethod(method string, methods []lib.VerificationMethod) (*lib.VerificationMethod, error) {
	for _, v := range methods {
		if method == v.ID {
			return &v, nil
		}
	}
	return nil, errors.New("unable to find matching verification method")
}

func findMyKey(key ed25519.PublicKey, methods []lib.VerificationMethod) (*lib.VerificationMethod, error) {
	for _, v := range methods {
		if bytes.Equal(base58.Decode(v.PublicKeyBase58), key) {
			return &v, nil
		}
	}
	return nil, errors.New("public key not found in verification methods")
}

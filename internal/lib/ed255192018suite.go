package lib

import (
	"bytes"
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"errors"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
	"github.com/piprate/json-gold/ld"
)

// Represents signing of a document using the Ed25519 2018 Signature suite https://w3c-ccg.github.io/lds-ed25519-2018/
// canonicalizationAlgorithm 	https://w3id.org/security#URDNA2015  [RDF-DATASET-NORMALIZATION]
// digestAlgorithm          	http://w3id.org/digests#sha512 	     [RFC6234]
// signatureAlgorithm       	http://w3id.org/security#ed25519 	 [ED25519]

const (
	KeyType       = "Ed25519VerificationKey2018"
	SignatureType = "Ed25519Signature2018"

	format    = "application/n-quads"
	algorithm = "URDNA2015"
)

var (
	processor = ld.NewJsonLdProcessor()
	byteDot   = []byte(".")
)

// GenerateProof takes in an unsigned document in byte array form, canonicalizes it, and appends a proof value
// with a nonce before signing the combination. The signature is added to the proof value, and the proof with
// signature is returned. This is in compliance with the Ed25519 2018 Linked Data Signature Suite.
func GenerateProof(input []byte, key ed25519.PrivateKey, verificationMethod string) (*Proof, error) {
	canonicalized, err := Canonicalize(input)
	if err != nil {
		return nil, err
	}

	// Create proof without signature value to be signed over
	proof := Proof{
		Type:               SignatureType,
		Created:            time.Now().Format(time.RFC3339),
		VerificationMethod: verificationMethod,
		Challenge:          uuid.New().String(),
	}

	// Append unsigned proof value to canonicalized document
	toSign, err := appendUnsignedProof(canonicalized, proof)
	if err != nil {
		return nil, err
	}

	// Do the signing
	signature, err := key.Sign(rand.Reader, toSign, crypto.Hash(0))
	if err != nil {
		return nil, err
	}

	// Add the base58 encoded signature value to the proof and return
	proof.SignatureValue = base58.Encode(signature)
	return &proof, nil
}

// VerifyProof takes in an unsigned document in byte array form, canonicalizes it, and appends the provided
// proof value without the signature value set. Then the proof without signature is appended to the input. The result
// is verified using the provided public key value.
func VerifyProof(input []byte, key ed25519.PublicKey, proof Proof) error {
	if proof.SignatureValue == "" {
		return errors.New("cannot verify proof without a signatureValue")
	}

	// canonicalize input to be safe
	canonicalized, err := Canonicalize(input)
	if err != nil {
		return err
	}

	// Copy proof and unset signature
	var withoutSignature Proof
	if err := Copy(&proof, &withoutSignature); err != nil {
		return err
	}
	withoutSignature.SignatureValue = ""

	// Add proof to input bytes
	toVerify, err := appendUnsignedProof(canonicalized, withoutSignature)
	if err != nil {
		return err
	}

	// Save decoded signature
	signature := base58.Decode(proof.SignatureValue)

	if verified := ed25519.Verify(key, toVerify, signature); !verified {
		return errors.New("could not verify signature")
	}
	return nil
}

func Canonicalize(input []byte) ([]byte, error) {
	options := ld.NewJsonLdOptions("")
	options.Format = format
	options.Algorithm = algorithm

	// convert to map[string]interface{} which is what the library expects
	out, err := toJSONStructure(input)
	if err != nil {
		return nil, err
	}

	normalized, err := processor.Normalize(out, options)
	if err != nil {
		return nil, err
	}
	return []byte(normalized.(string)), nil
}

func appendUnsignedProof(input []byte, proof Proof) ([]byte, error) {
	// Can't already have a signature
	if proof.SignatureValue != "" {
		return nil, errors.New("proof already has a signature value")
	}
	reader := bytes.NewBuffer(input)
	reader.Write(byteDot)

	proofBytes, err := json.Marshal(proof)
	if err != nil {
		return nil, err
	}
	reader.Write(proofBytes)
	return reader.Bytes(), nil
}

// takes JSON bytes and turns into a go representation of JSON
func toJSONStructure(input []byte) (interface{}, error) {
	var out map[string]interface{}
	err := json.Unmarshal(input, &out)
	return out, err
}

package lib

// SignedDIDDoc wraps a DID Doc with a Linked Data proof object
type SignedDIDDoc struct {
	DIDDoc
	Proof *Proof `json:"proof,omitempty"`
}

// DID structure as represented in the lib-core specification https://w3c.github.io/did-core/
// Up to date as of June 21, 2020

type DIDDoc struct {
	ID                  string               `json:"id"`
	VerificationMethods []VerificationMethod `json:"verificationMethod"`
	Authentication      []string             `json:"authentication,omitempty"`
	ServiceEndpoints    []ServiceEndpoint    `json:"service,omitempty"`
	Created             string               `json:"created"`
	Updated             string               `json:"updated,omitempty"`
}

type VerificationMethod struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	Controller      string `json:"controller"`
	PublicKeyBase58 string `json:"publicKeyBase58"`
}

type ServiceEndpoint struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}

type Proof struct {
	Type               string `json:"type"`
	Created            string `json:"created"`
	VerificationMethod string `json:"verificationMethod"`
	Challenge          string `json:"challenge,omitempty"`
	SignatureValue     string `json:"signatureValue,omitempty"`
}

type Tweet struct {
	Tweet string    `json:"tweet"`
	Proof Proof `json:"proof"`
}

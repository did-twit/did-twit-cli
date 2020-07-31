package pkg

type Tweet struct {
	Tweet string `json:"tweet"`
	Proof Proof  `json:"proof"`
}

type Proof struct {
	Type               string `json:"type"`
	Created            string `json:"created"`
	VerificationMethod string `json:"verificationMethod"`
	Challenge          string `json:"challenge,omitempty"`
	SignatureValue     string `json:"signatureValue,omitempty"`
}

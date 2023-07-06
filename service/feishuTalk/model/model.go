package model

type ChallengeEncrypt struct {
	Encrypt string `json:"encrypt"`
}

type Challenge struct {
	Challenge string `json:"challenge"`
	Token     string `json:"token"`
	Type      string `json:"type"`
}

package models

type ChallengeMessage struct {
	Type     string `json:"type"`
	From     string `json:"from"`
	To       string `json:"to"`
	Accepted *bool  `json:"accepted,omitempty"`
}

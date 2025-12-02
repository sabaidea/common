package mq

import "fmt"

type ConsumeMessage struct {
	Action       Action       `json:"action"`
	Service      Service      `json:"service"`
	ReceiverType ReceiverType `json:"receiver_type"`
	Receiver     Receiver     `json:"receiver"`
	Platform     Platform     `json:"platform"`
	MetaData     *Metadata    `json:"metadata,omitempty"`
	Payload      interface{}  `json:"payload"`
}

func (m ConsumeMessage) Validate() error {
	if !m.Action.IsValid() {
		return fmt.Errorf("invalid action: %s", m.Action)
	}
	if !m.Service.IsValid() {
		return fmt.Errorf("invalid service: %s", m.Service)
	}
	if !m.ReceiverType.IsValid() {
		return fmt.Errorf("invalid receiver_type: %s", m.ReceiverType)
	}
	if !m.Platform.IsValid() {
		return fmt.Errorf("invalid platform: %s", m.Platform)
	}
	return nil
}

package messages

import "fmt"

// ProduceMessage type of message that services must produce to mercury
type ProduceMessage struct {
	Action       Action       `json:"action"`
	Service      Service      `json:"service"`
	ReceiverType ReceiverType `json:"receiver_type"`
	Receiver     Receiver     `json:"receiver"`
	Platform     Platform     `json:"platform"`
	MetaData     *Metadata    `json:"metadata,omitempty"`
	Payload      interface{}  `json:"payload"`
}

func (p *ProduceMessage) Validate() error {
	if !p.Action.IsValid() {
		return fmt.Errorf("invalid action: %s", p.Action)
	}
	if !p.Service.IsValid() {
		return fmt.Errorf("invalid service: %s", p.Service)
	}
	if !p.ReceiverType.IsValid() {
		return fmt.Errorf("invalid receiver_type: %s", p.ReceiverType)
	}
	if !p.Platform.IsValid() {
		return fmt.Errorf("invalid platform: %s", p.Platform)
	}
	return nil
}

package messages

type ConsumeMessage struct {
	Metadata Metadata    `json:"metadata"`
	Payload  interface{} `json:"payload"`
}

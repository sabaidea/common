package mq

type ReceiverType string

const (
	RCIReceiver  ReceiverType = "rci"
	UserReceiver ReceiverType = "user"
	RoomReceiver ReceiverType = "room"
)

func (rt *ReceiverType) IsValid() bool {
	if rt == nil {
		return false
	}
	switch *rt {
	case RCIReceiver, UserReceiver, RoomReceiver:
		return true
	default:
		return false
	}
}

func (rt *ReceiverType) String() string {
	return string(*rt)
}

type Receiver string

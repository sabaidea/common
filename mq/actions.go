package mq

type Action string

const (
	SendMessageAction Action = "send_message"
	JoinRoomAction    Action = "join_room"
	LeaveRoomAction   Action = "leave_room"
	DeleteRoomAction  Action = "delete_room"
)

func (a *Action) IsValid() bool {
	if a == nil {
		return false
	}
	switch *a {
	case SendMessageAction, JoinRoomAction, LeaveRoomAction, DeleteRoomAction:
		return true
	default:
		return false
	}
}

func (a *Action) String() string {
	return string(*a)
}

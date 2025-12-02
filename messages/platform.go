package messages

type Platform string

const (
	AparatSport Platform = "aparatsport"
)

func (p *Platform) IsValid() bool {
	if p == nil {
		return false
	}
	switch *p {
	case AparatSport:
		return true
	default:
		return false
	}
}

package mq

type Service string

const (
	Mercury  Service = "mercury"
	Hermes   Service = "hermes"
	Kharazmi Service = "kharazmi"
	Odin     Service = "odin"
)

func (s *Service) String() string {
	return string(*s)
}

func (s *Service) IsValid() bool {
	if s == nil {
		return false
	}
	switch *s {
	case Mercury, Hermes, Kharazmi, Odin:
		return true
	default:
		return false
	}
}

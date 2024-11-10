package amadla

type IAmadla interface {
	VersionSupport() string
}

type SAmadla struct{}

// VersionSupport
// TODO: Need to comme up with some logic
func (s *SAmadla) VersionSupport() string {
	return "^0"
}

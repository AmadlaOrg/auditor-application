package support

type ISupport interface{}

type SSupport struct{}

func (s *SSupport) List() []string {
	return []string{
		"github.com/AmadlaOrg/EntityApplication",
	}
}

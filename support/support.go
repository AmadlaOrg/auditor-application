package support

type ISupport interface {
	List() []string
}

type SSupport struct{}

// List all the supported entities
func (s *SSupport) List() []string {
	return []string{
		"github.com/AmadlaOrg/EntityApplication",
	}
}

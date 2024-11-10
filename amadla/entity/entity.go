package entity

type IEntity interface {
	List() []string
}

type SEntity struct{}

// List all the supported entities
func (s *SEntity) List() []string {
	return []string{
		"github.com/AmadlaOrg/EntityApplication",
	}
}

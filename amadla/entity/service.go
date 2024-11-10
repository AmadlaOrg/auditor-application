package entity

// NewEntityService to set up the support service
func NewEntityService() IEntity {
	return &SEntity{}
}

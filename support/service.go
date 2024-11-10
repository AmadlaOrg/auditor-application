package support

// NewSupportService to set up the support service
func NewSupportService() ISupport {
	return &SSupport{}
}

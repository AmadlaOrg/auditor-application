package amadla

// NewAmadlaService to set up the support service
func NewAmadlaService() IAmadla {
	return &SAmadla{}
}

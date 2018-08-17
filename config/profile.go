package config

// Profile defines the current running app profile
type Profile int

const (
	// Dev profile
	Dev Profile = iota
	//Integration profile
	Integration
	//Production profile
	Production
)

// PropertyFile returns the name of the active profile configuration file
func (p *Profile) PropertyFile() string {
	names := [...]string{
		"application-dev.json",
		"application-int.json",
		"application.json",
	}

	if *p < Dev || *p > Production {
		logger.Error("Unknown profile was requested.")
		return "application.json"
	}

	return names[*p]
}

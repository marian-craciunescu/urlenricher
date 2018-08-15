package config

type Profile int

const (
	Dev Profile = iota
	Integration
	Production
)

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

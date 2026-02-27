package ports

// ProcessPort represents a parsed port and its associated process.
type ProcessPort struct {
	Port     string
	Protocol string
	PID      string
	Process  string
	Status   string
	User     string
}

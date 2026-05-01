package target

type TargetType string

const (
	TargetDocker TargetType = "docker"
	TargetPID    TargetType = "pid"
)

type Target struct {
	Type TargetType
	ID   string
	PID  int
}

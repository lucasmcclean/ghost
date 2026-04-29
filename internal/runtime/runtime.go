package runtime

import (
	"fmt"
	"strconv"
)

type Runtime interface {
	PIDFor(container string) (int, error)
}

func Resolve(id string) (int, error) {
	if pid, err := strconv.Atoi(id); err == nil {
		return pid, nil
	}

	if pid, err := (&Docker{}).PIDFor(id); err == nil {
		return pid, nil
	}

	return 0, fmt.Errorf("cannot resolve %q to a PID", id)
}

package runtime

import (
	"os/exec"
	"strconv"
	"strings"
)

type Docker struct{}

func (d *Docker) PIDFor(container string) (int, error) {
	cmd := exec.Command("docker", "inspect", "-f", "{{.State.Pid}}", container)
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	pidStr := strings.TrimSpace(string(out))
	return strconv.Atoi(pidStr)
}

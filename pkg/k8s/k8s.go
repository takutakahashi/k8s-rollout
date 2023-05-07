package k8s

import (
	"os/exec"

	"github.com/pkg/errors"
)

func Rollout(resource, namespace, name string) (string, error) {
	out, err := exec.Command("kubectl", "rollout", "restart", resource, name, "-n", namespace).CombinedOutput()
	if err != nil {
		return "", errors.Wrap(err, string(out))
	}
	return string(out), nil
}

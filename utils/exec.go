/**
 * @Author:      leafney
 * @Date:        2022-12-23 20:29
 * @Project:     rpi-monitor
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package utils

import (
	"os/exec"
	"strings"
)

// RunCommand execute shell command
func RunCommand(str string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", str)
	outBytes, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	// Clean the output and remove special characters
	outStr := strings.TrimSpace(string(outBytes))
	return outStr, nil
}

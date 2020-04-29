package cmdtool

import (
	"bufio"
	"errors"
	"os/exec"
	"runtime"
)

func Cmd(cmdString string) (result string, err error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd")
	case "linux":
		cmd = exec.Command("/bin/bash", "-c", cmdString)
	default:
		err = errors.New("OS only support LINUX and WINDOWS")
		return
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}

	if err = cmd.Start(); err != nil {
		return
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result += scanner.Text() + "\n"
	}

	if err = cmd.Wait(); err != nil {
		return
	}

	return
}

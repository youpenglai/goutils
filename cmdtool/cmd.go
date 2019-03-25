package cmdtool

import (
	"bytes"
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
		cmd = exec.Command("/bin/bash")
	default:
		return "", errors.New("OS only support LINUX and WINDOWS")
	}

	var in, out bytes.Buffer
	cmd.Stdin, cmd.Stdout = &in, &out
	_, err = in.WriteString(cmdString)
	if err != nil {
		return
	}

	err = cmd.Start()
	if err != nil {
		return
	}

	err = cmd.Wait()
	if err != nil {
		return
	}

	result = out.String()
	return
}

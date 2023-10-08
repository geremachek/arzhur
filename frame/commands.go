package frame

import (
	"io"
	"os/exec"
	"strings"
)

// run an external command

func runExternal(command string, input string, pipe bool) (string, error) {
	arguments := strings.Split(strings.TrimSpace(command), " ")
	cmd := exec.Command(arguments[0], arguments[1:]...)

	// if pipe is true, send input to the command's stdin

	if pipe {
		if stdin, err := cmd.StdinPipe(); err == nil {
			go func() {
				defer stdin.Close()
				io.WriteString(stdin, input)
			}()
		} else {
			return "", err
		}
	}

	if out, err := cmd.Output(); err == nil {
		return string(out), nil
	} else {
		return "", err
	}
}
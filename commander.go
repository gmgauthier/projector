package main

import (
	"os/exec"
	"strings"
)

func execute(cmdstr string) (string, error) {
	cmdargs := strings.Split(cmdstr, " ")         // string arrayified
	cmd := cmdargs[0]                             // command
	cmdargs = append(cmdargs[:0], cmdargs[1:]...) // argument array sans cmd
	out, err := exec.Command(cmd, cmdargs...).CombinedOutput()
	return string(out[:]), err
}

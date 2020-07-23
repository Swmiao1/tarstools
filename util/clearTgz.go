package util

import (
	"os/exec"
)

func ClearTgz() {
	clear := exec.Command("cmd")
	infile, _ := clear.StdinPipe()
	_, _ = infile.Write([]byte("del *.tgz\nexit\n"))
}

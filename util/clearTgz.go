package util

import (
	"fmt"
	"os/exec"
)

func ClearTgz() {
	fmt.Println("正在清理Tgz")
	clear := exec.Command("cmd")
	infile, _ := clear.StdinPipe()
	_, _ = infile.Write([]byte("del *.tgz\nexit\n"))
	_ = clear.Wait()
}

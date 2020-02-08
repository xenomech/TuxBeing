package main

import (
	"os/exec"
	"strings"
	"time"
)

func exe_cmd(cmd string){
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]
	exec.Command(head,parts...).Output()
	}

func main() {
		// counter:=1
		var command = "xinput disable 17"
		exe_cmd(command)
		time.Sleep(120*time.Second)
		var command1 = "xinput enable 17"
		exe_cmd(command1)
}

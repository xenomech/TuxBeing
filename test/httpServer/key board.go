package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func exe_cmd(cmd string)(string) {
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]
	out, err := exec.Command(head,parts...).Output()
	if err != nil {
		fmt.Println(cmd)
	  fmt.Printf("Error:%s", err)
	  return "err"
	}
	}

func main() {
		flag:=1
		var command = "xdotool disable _________"
		exe_cmd(command)
	for{
		if(flag==120){
				var command = "xdotool enable _________"
				exe_cmd(command)
			}
	}
}

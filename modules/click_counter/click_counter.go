package main

import (
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	// "os"
	"os/exec"
	// "strings"
	// "time"
)

func runcmd(cmd string, status int) string {
	out, err := exec.Command("bash", cmd).Output()
	if err != nil {
		fmt.Println(cmd)
		fmt.Printf("Error:%s", err)
		return "err"
	}
	return string(out)
}
func main() {

	// xinput --list | grep -i -m 1 'mouse' | grep -o 'id=[0-9]\+' | grep -o '[0-9]\+'
	var mouseID = runcmd("xinput --list | grep -i -m 1 'mouse' | grep -o 'id=[0-9]\+' | grep -o '[0-9]\+' ", 0)
	fmt.Println(mouseID)
}

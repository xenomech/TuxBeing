package main

import(
	"fmt"
	"os/exec"
	"strings"
)

func get_process_name(arg string)(string){
	var command = "ps"
	var argument = "-p "+arg+" -o comm="
	parts := strings.Fields(argument)
	parts = parts[0:len(parts)]
	out,err:= exec.Command(command,parts...).Output()
	if err != nil{
		fmt.Printf("Error:%s", err)
	}
	return string(out)
}
func exe_cmd(cmd string)(string) {
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]
	out, err := exec.Command(head,parts...).Output()
	if err != nil {
	  fmt.Printf("Error:%s", err)
	}
	arg := string(out)
	arg = arg[:len(arg) - 1]
	process_name := get_process_name(arg)
	return process_name
}

func main(){
	var command = "xdotool getwindowfocus   getwindowpid"
	// var previous_command = exe_cmd(command)
	for{
		process_name := exe_cmd(command)
		fmt.Printf(process_name)
	}
}
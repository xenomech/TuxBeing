package main

import(
	"fmt"
	"encoding/json"
	"os/exec"
	"strings"
	"time"
	"io/ioutil"
)
type Process struct{
	Name string `json:name"`
	Time float64 `json:"time"`
}
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
		fmt.Println(cmd)
	  fmt.Printf("Error:%s", err)
	}
	arg := string(out)
	arg = arg[:len(arg) - 1]
	process_name := get_process_name(arg)
	return process_name
}

func main(){
	// car current_details = getFromJson()
	var  current_details = make([]Process,0)
	var command = "xdotool getwindowfocus   getwindowpid"
	var previous_command = exe_cmd(command)
	var time_command = time.Now()
	var process_name = exe_cmd(command)
	current_details = append(current_details,Process{
		Name: process_name,
		Time: 0,
	})
	// fmt.Println(current_details)
	for{
		flag :=1
		process_name = exe_cmd(command)
		if process_name != previous_command{
			current_time := time.Now()
			time_used := current_time.Sub(time_command).Hours()
			// current_details = addToJson(current_details,process_name,time_used)
			for _,data := range current_details{
				if data.Name == process_name && flag == 1{
					cur_time := data.Time
					cur_time += time_used
					data.Time = cur_time
					flag = 0
				}
			}
			if flag ==1{
				current_details = append(current_details,Process{
					Name:process_name,
					Time:time_used,
				})
			}
			time_command = current_time
			previous_command = process_name
			json_data,err := json.Marshal(current_details)
			if err != nil{
				fmt.Println("error_jsonMarshal:%s",err)
			}
			ioutil.WriteFile("output.json", json_data,0644)
		}
	}
}

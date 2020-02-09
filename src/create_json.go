package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Process struct {
	Name string  `json:"name"`
	Time float64 `json:"time"`
}

func get_process_name(arg string) string {
	var command = "ps"
	var argument = "-p " + arg + " -o comm="
	parts := strings.Fields(argument)
	parts = parts[0:len(parts)]
	out, err := exec.Command(command, parts...).Output()
	if err != nil {
		fmt.Printf("Error:%s", err)
	}
	return string(out)
}
func exe_cmd(cmd string, status int) string {
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]
	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Println(cmd)
		fmt.Printf("Error:%s", err)
		return "err"
	}
	if status ==1{
		return " "
	}
	arg := string(out)
	arg = arg[:len(arg)-1]
	process_name := get_process_name(arg)
	process_name = process_name[:len(process_name) -1]
	// if status ==0{
	return process_name
}
func file_exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {
	c := make(chan int)
	d := make(chan int)
	go Server(c)
	go ui(d)
	var current_details []Process
	var command = "xdotool getwindowfocus   getwindowpid"
	var process_name = exe_cmd(command, 0)
	if file_exists("output.json") {
		data, err := ioutil.ReadFile("output.json")
		if err != nil {
			fmt.Println("Error:", err)
		}
		err = json.Unmarshal(data, &current_details)
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else {
		current_details = make([]Process, 0)
		current_details = append(current_details, Process{
			Name: process_name,
			Time: 0,
		})
	}
	var previous_command = exe_cmd(command, 0)
	var time_command = time.Now()

	for {
		flag := 1
		process_name = exe_cmd(command, 0)
		if process_name == "err" {
			break
		}
		if process_name != previous_command {
			current_time := time.Now()
			time_used := current_time.Sub(time_command).Hours()
			for index, data := range current_details {
				if data.Name == process_name && flag == 1 {
					cur_time := data.Time
					cur_time += time_used
					current_details[index] = current_details[len(current_details)-1]
					current_details = current_details[:len(current_details)-1]
					current_details = append(current_details, Process{
						Name: process_name,
						Time: cur_time,
					})
					flag = 0
				}
			}
			if flag == 1 {
				current_details = append(current_details, Process{
					Name: process_name,
					Time: time_used,
				})
			}
			time_command = current_time
			previous_command = process_name
			json_data, err := json.Marshal(current_details)
			if err != nil {
				fmt.Println("error_jsonMarshal:%s", err)
			}
			ioutil.WriteFile("output.json", json_data, 0644)
		}
	}
	<-c
	<-d
}

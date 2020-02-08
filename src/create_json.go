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
	Name string `json:"name"`
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
func update_items(){

}
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
	arg := string(out)
	arg = arg[:len(arg) - 1]
	process_name := get_process_name(arg)
	process_name = process_name[:len(process_name) -1]
	return process_name
}
func lock_screen(){
	fmt.Printf("in lock screen: %d", time.Now().Unix())
	time.Sleep(30 * time.Second)
	fmt.Printf("After sleep: %d", time.Now().Unix())
	lock_screen()
}

func main(){
	// go lock_screen()
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
		if process_name == "err"{
			break
		}
		if process_name != previous_command{
			current_time := time.Now()
			time_used := current_time.Sub(time_command).Hours()
			// test_array := current_details
			// current_details = addToJson(current_details,process_name,time_used)
			for index,data := range current_details{
				if data.Name == process_name && flag == 1{
					cur_time := data.Time
					cur_time += time_used
					// data.Time = cur_time
					current_details[index] = current_details[len(current_details)-1]
					current_details = current_details[:len(current_details)-1]
					current_details = append(current_details,Process{
						Name: process_name,
						Time: cur_time,
					})
					flag = 0
				}
			}
			fmt.Println("After forloop")
			fmt.Println(current_details)
			if flag ==1{
				current_details = append(current_details,Process{
					Name:process_name,
					Time:time_used,
				})
			}
			fmt.Println(current_details)
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

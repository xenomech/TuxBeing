package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Process1 struct {
	Process1 []Process `json:"process"`
}

type Process struct {
	Name string  `json:"name"`
	Time float64 `json:"time"`
}

func main() {
	fmt.Println("Hello World")
	jsonFile, err := os.Open("test.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var process1 Process1
	json.Unmarshal(byteValue, &process1)
	for i := 0; i < len(process1.Process1); i++ {
		fmt.Println(process1.Process1[i].Name)
		fmt.Println(process1.Process1[i].Time)
	}

}

//jsonFile, err := os.Open("users.json")
//if err != nil {
//    fmt.Println(err)
//js}

package main
import (
    "fmt"
    "log"
    "net/http"
    "os"
    "encoding/json"
    "io/ioutil"
)

// type Process struct{
// 	Name string `json:"name"`
// 	Time float64 `json:"time"`
// }

func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    switch r.Method {

    case "GET":

        jsonFile,_ := os.Open("output.json")
         user_data, err  := ioutil.ReadAll(jsonFile)
         var data []Process
         json.Unmarshal(user_data, &data)
         if err !=nil {
             fmt.Println("error:%s", err)
         }
         w.Header().Set("Content-Type", "application/json; charset=utf-8")
         w.Header().Set("Access-Control-Allow-Origin", "*")
         json.NewEncoder(w).Encode(data)

    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}
func enable_keyboard(w http.ResponseWriter, r *http.Request){
    fmt.Println("Enabling keyboard")
    if r.URL.Path != "/quit/" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    switch r.Method{
    case "POST":
        var command_enable_keyboard = "xinput enable 17"
        _ =exe_cmd(command_enable_keyboard,1)
    default:
        http.Error(w, "404 not found.", http.StatusNotFound)
    }
}

func disable_keyboard(w http.ResponseWriter, r *http.Request){
    fmt.Println("Disabling keyboard")
    if r.URL.Path != "/disable/" {
        http.Error(w,"Unauthorized", http.StatusUnauthorized)
        return
    }
    switch r.Method{
    case "POST":
        var command_disable_keyboard = "xinput disable 17"
        _ = exe_cmd(command_disable_keyboard,1)
    default:
        http.Error(w,"404 not found.",http.StatusNotFound)
    }
}

func Server(c chan int) {
    http.HandleFunc("/", hello)
    http.HandleFunc("/quit/",enable_keyboard)
    http.HandleFunc("/disable/",disable_keyboard)
    fmt.Printf(" Server running at 8080...\n")

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
    c <- 1
}

package main
import (
    "fmt"
    "log"
    "net/http"
    "os"
    "encoding/json"
    "io/ioutil"
)

type Process struct{
	Name string `json:"name"`
	Time float64 `json:"time"`
}

func hello(w http.ResponseWriter, r *http.Request) {
    
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    switch r.Method {

    case "GET":     

        jsonFile,_ := os.Open("../../src/output.json")
         user_data, err  := ioutil.ReadAll(jsonFile) 
         var data []Process
         json.Unmarshal(user_data, &data)
         if err !=nil {
             fmt.Println("error:%s", err)
         }
         
         w.Header().Set("Content-Type", "application/json; charset=utf-8")
         w.Header().Set("Access-Control-Allow-Origin", "*")
         json.NewEncoder(w).Encode(data)
    case "POST":
        // Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
           return
        }
        fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
        name := r.FormValue("name")
        address := r.FormValue("address")
        fmt.Fprintf(w, "Name = %s\n", name)
        fmt.Fprintf(w, "Address = %s\n", address)

    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}

func main() {
    http.HandleFunc("/", hello)
    fmt.Printf("Starting server for testing HTTP POST...\n")
    
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

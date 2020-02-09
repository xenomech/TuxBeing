//multidev firebase access

package main
import (
	"context"
	"fmt"
	"log"
	"google.golang.org/api/iterator"
  	firebase "firebase.google.com/go"
  	"google.golang.org/api/option"
	"encoding/json"
	"os"
	"io/ioutil"
)
type Process struct{
	Name string `json:"name"`
	Time float64 `json:"time"`
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
// Use a service account
var current_details []Process

ctx := context.Background()
sa := option.WithCredentialsFile("./serviceAccount.json")
app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
  		log.Fatalln(err)
	}
client, err := app.Firestore(ctx)
if err != nil {
  log.Fatalln(err)
}

defer client.Close()

// [END fs_initialize]


if file_exists("output.json"){
		data, err := ioutil.ReadFile("output.json")
		if err != nil{
			fmt.Println("Error:",err)
		}
		err = json.Unmarshal(data, &current_details)
		if err != nil{
			fmt.Println("Error:",err)
		}
// [START fs_add_data_1]
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first":  "Alan",
		"middle": "Mathison",
		"last":   "Turing",
		"born":   1912,})
		if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	// [END fs_add_data_1]




	// [START fs_get_all_users]
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
	// [END fs_get_all_users]
}}
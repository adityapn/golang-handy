package main

import ("fmt"
		"encoding/json"
		"io"
		"net/http")

type App struct {
    Id string `json:"id"`
    Title string `json:"title"`
}

func main() {

	data := []byte(`
	    {
	        "id": "k34rAT4",
	        "title": "My Awesome App"
	    }
	`)

	// Converting input json to struct
	var app App
	err := json.Unmarshal(data, &app)	
	if err != nil{
		fmt.Println("Error")
		fmt.Sprintln("can not serialize json,Error ",err)
		return
	}

	// Struct to json
	var newApp App
	newApp.Id = "12"
	newApp.Title = "twelle"
	fmt.Println("new response")
	response, error := json.Marshal(newApp)
	if error != nil{
		fmt.Println("error ",error)
		return
	}
	fmt.Println(response)
}




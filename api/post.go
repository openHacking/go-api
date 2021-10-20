package main

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

// content represents data about a record content
type Post struct {
	ID    int `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	Lang  string `json:"lang"`
}

/*
get json data from local json file

reference : https://tutorialedge.net/golang/parsing-json-with-golang/#reading-and-parsing-a-json-file

*/
func getPostsJson() []Post {
	// Open our jsonFile
	jsonFile, err := os.Open(getCurrentAbPathByCaller()+"/posts.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array
	byteValue,_ := ioutil.ReadAll(jsonFile)

	// we initialize our posts array
	var posts []Post

	// we unmarshal our byteArray which contains our jsonFile's content into 'posts' which we defined above
	json.Unmarshal(byteValue,&posts)

	return posts
}


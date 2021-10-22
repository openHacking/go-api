package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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
func getPostsJsonByRemote() []Post {
	// Open our jsonFile
	// jsonFile, err := os.Open(getCurrentAbPathByCaller()+"/posts.json")
	// jsonFile, err := os.Open("posts.json")
	url := "https://cdn.jsdelivr.net/gh/openHacking/static-files/cache/posts.json"

	spaceClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	
	if getErr != nil {
		log.Fatal(getErr)
	}

	log.Printf("after getPostsJsonByRemote", time.Now())
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// we initialize our posts array
	var posts []Post

	// we unmarshal our body which contains our jsonFile's content into 'posts' which we defined above
	jsonErr := json.Unmarshal(body, &posts)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// save file
	savePostsJson(posts)

	return posts
}
func getPostsJsonByLocal() []Post {
	// Open our jsonFile
	// jsonFile, err := os.Open(getCurrentAbPathByCaller()+"/posts.json")
	jsonFile, err := os.Open("posts.json")

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

func savePostsJson(posts []Post)  {

	jsonData, err := json.Marshal(posts)

	if err != nil {
	   fmt.Println(err)
	}

	// write to JSON file

	jsonFile, err := os.Create("./posts.json")

	if err != nil {
	   fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())
}
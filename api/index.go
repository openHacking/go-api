/*
  export index module

  FTS reference: https://github.com/akrylysov/simplefts
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// index is an inverted index. It maps tokens to post IDs.
type index map[string][]int

// add adds posts to the index.
func (idx index) add(posts []Post) {
	for _, post := range posts {
		for _, token := range analyze(post.Title) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == post.ID {
				// Don't add same ID twice.
				continue
			}
			idx[token] = append(ids, post.ID)
		}
	}
}

// intersection returns the set intersection between a and b.
// a and b have to be sorted in ascending order and contain no duplicates.
func intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

// search queries the index for the given text.
func (idx index) search(text string) []int {
	var r []int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			// Token doesn't exist.
			return nil
		}
	}
	return r
}

func indexed(docs []Post) index  {
	idx := make(index)
	idx.add(docs)

	// save file
	saveIndexedJson(idx)
	fmt.Println("indexed success")
	return idx
}

func saveIndexedJson(idx index)  {

         jsonData, err := json.Marshal(idx)

         if err != nil {
			fmt.Println(err)
         }

         // write to JSON file

         jsonFile, err := os.Create("./indexed.json")

         if err != nil {
			fmt.Println(err)
         }
         defer jsonFile.Close()

         jsonFile.Write(jsonData)
         jsonFile.Close()
         fmt.Println("JSON data written to ", jsonFile.Name())
}

func getIndexedJson() index  {
	// Open our jsonFile
	jsonFile, err := os.Open("indexed.json")
	// jsonFile, err := os.Open(getCurrentAbPathByCaller()+"/indexed.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array
	byteValue,_ := ioutil.ReadAll(jsonFile)

	// we initialize our posts array
	var idx index

	// we unmarshal our byteArray which contains our jsonFile's content into 'posts' which we defined above
	json.Unmarshal(byteValue,&idx)

	return idx
}

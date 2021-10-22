package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func storeJsonFile() {

	log.Printf("before getPostsJsonByRemote", time.Now())
	posts := getPostsJsonByRemote()

	log.Printf("before indexed", time.Now())
	indexed(posts)

}
func searchText(c *gin.Context) {
	q := c.Query("q")

	// start := time.Now()

	posts := getPostsJsonByLocal()

	idx := getIndexedJson()

	matchedIDs := idx.search(q)

	res := make([]Post, 0, len(matchedIDs))
	for _, id := range matchedIDs {
		post := posts[id]
		res = append(res, post)
	}

	// log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))
	// fmt.Println(q)
	c.IndentedJSON(http.StatusOK, res)
}

func main() {

	storeJsonFile()

	router := gin.Default()

	router.GET("/search", searchText)
	router.Run("0.0.0.0:8081")
}

////////////////////////////

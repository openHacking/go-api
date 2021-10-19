package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Lang  string `json:"lang"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "There-isn't-anything-to-compare", Title: "The Master Submits a PR to Main Branch and appears:There Isn’t Anything To Compare", Lang: "en"}, {ID: "cannot-find-module-scss", Title: "Cannot Find Module 'XXX.scss' or Its Corresponding Type Declarations", Lang: "en"}, {ID: "canvas-basic-api", Title: "Canvas Basic API Usage Tutorial", Lang: "en"}, {ID: "customized-range-slider", Title: "Customized Range Slider", Lang: "en"}, {ID: "format-date-time-add-zero", Title: "JS Regular Format Date and Time Automatically Fill 0", Lang: "en"}, {ID: "full-screen-pop-up-window-drag-and-drop-by-vanilla-js", Title: "Full-Screen Pop-Up Window Drag and Drop by Vanilla JavaScript", Lang: "en"}, {ID: "github-cannot-be-accessed", Title: "GitHub Clone Is Very Slow, GitHub Cannot Be Uploaded, github.io Cannot Be Accessed", Lang: "en"}, {ID: "gnvm-switch-node-version", Title: "Gnvm Switch Node Version", Lang: "en"}, {ID: "how-to-clear-all-github-commit-records", Title: "How To Clear All GitHub Commit Records?", Lang: "en"}, {ID: "how-to-clear-wechat-cache", Title: "How To Clear WeChat Cache", Lang: "en"}, {ID: "jquery-apply-is-not-a-function", Title: "jQuery Apply Is Not a Function", Lang: "en"}, {ID: "leetcode-binary-tree-postorder-traversal", Title: "LeetCode Notes: Binary Tree Postorder Traversal", Lang: "en"}, {ID: "leetcode-can-place-flowers", Title: "LeetCode Notes: Can Place Flowers", Lang: "en"}, {ID: "leetcode-combine-two-tables", Title: "LeetCode Notes: Combine Two Tables", Lang: "en"}, {ID: "leetcode-contains-duplicate-ii", Title: "LeetCode Notes: Contains Duplicate II", Lang: "en"}, {ID: "leetcode-counting-bits", Title: "LeetCode Notes: Counting Bits", Lang: "en"}, {ID: "leetcode-delete-the-node-of-the-linked-list", Title: "LeetCode Notes: Delete the Node of the Linked List", Lang: "en"}, {ID: "leetcode-excel-sheet-column-number", Title: "LeetCode Notes: Excel Sheet Column Number", Lang: "en"}, {ID: "leetcode-excel-sheet-column-title", Title: "LeetCode Notes: Excel Sheet Column Title", Lang: "en"}, {ID: "leetcode-find-kth-node-from-the-end-linked-list", Title: "LeetCode Notes: Find K’th Node From the End of a Linked List", Lang: "en"}, {ID: "leetcode-hamming-distance", Title: "LeetCode Notes: Hamming Distance", Lang: "en"}, {ID: "leetcode-happy-number", Title: "LeetCode Notes: Happy Number", Lang: "en"}, {ID: "leetcode-implement-queue-using-stacks", Title: "LeetCode Notes: Implement Queue using Stacks", Lang: "en"}, {ID: "leetcode-implement-stack-using-queues", Title: "LeetCode Notes: Implement Stack using Queues", Lang: "en"}, {ID: "leetcode-intersection-of-two-arrays", Title: "LeetCode Notes: Intersection of Two Arrays", Lang: "en"}, {ID: "leetcode-invert-binary-tree", Title: "LeetCode Notes: Invert Binary Tree", Lang: "en"}, {ID: "leetcode-isomorphic-strings", Title: "LeetCode Notes: Isomorphic Strings", Lang: "en"}, {ID: "leetcode-longest-palindrome", Title: "LeetCode Notes: Longest Palindrome", Lang: "en"}, {ID: "leetcode-majority-element", Title: "LeetCode Notes: Majority Element", Lang: "en"}, {ID: "leetcode-middle-of-the-linked-list", Title: "LeetCode Notes: Middle of the Linked List", Lang: "en"}, {ID: "leetcode-min-stack", Title: "LeetCode Notes: Min Stack", Lang: "en"}, {ID: "leetcode-number-of-1-bits", Title: "LeetCode Notes: Number of 1 Bits", Lang: "en"}, {ID: "leetcode-numbers-missing-from-0-to-n-1", Title: "LeetCode Notes: Numbers Missing From 0 to N-1", Lang: "en"}, {ID: "leetcode-palindrome-linked-list", Title: "LeetCode Notes: Palindrome Linked List", Lang: "en"}, {ID: "leetcode-power-of-two", Title: "LeetCode Notes: Power of Two", Lang: "en"}, {ID: "leetcode-print-linked-list-from-end-to-beginning", Title: "LeetCode Notes: Print linked list from end to beginning", Lang: "en"}, {ID: "leetcode-reverse-bits", Title: "LeetCode Notes: Reverse Bits", Lang: "en"}, {ID: "leetcode-reverse-linked-list", Title: "LeetCode Notes: Reverse Linked List", Lang: "en"}, {ID: "leetcode-reverse-string", Title: "LeetCode Notes: Reverse String", Lang: "en"}, {ID: "leetcode-sequence-of-consecutive-positive-numbers-whose-sum-is-s", Title: "LeetCode Notes: Sequence of Consecutive Positive Numbers Whose Sum Is S", Lang: "en"}, {ID: "leetcode-single-number", Title: "LeetCode Notes: Single Number", Lang: "en"}, {ID: "leetcode-spiral-matrix", Title: "LeetCode Notes: Spiral Matrix", Lang: "en"}, {ID: "leetcode-summary-ranges", Title: "LeetCode Notes: Summary Ranges", Lang: "en"}, {ID: "leetcode-valid-palindrome", Title: "LeetCode Notes: Valid Palindrome", Lang: "en"}, {ID: "npm-install-error-guide", Title: "Npm Install Reports an Error and Freezes, Npm Installation Guide", Lang: "en"}, {ID: "python-selenium-life-restart", Title: "Python + Selenium Automated Test Life Restart Simulator", Lang: "en"}, {ID: "recommend-vuepress-template", Title: "Recommend a Vuepress Template, One-Click To Quickly Build a Document Station", Lang: "en"}, {ID: "tree-menu-node-path", Title: "Get the Tree Menu Node Path", Lang: "en"},
}

func getAlbumByID(c *gin.Context) {

	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func searchText(c *gin.Context) {
	q := c.Query("q")
	
	
}
func main() {

	// store and tranform  all content
	// search.StoreContent()

	router := gin.Default()

	router.GET("/search", searchText)
	router.Run("localhost:8080")
}

////////////////////////////

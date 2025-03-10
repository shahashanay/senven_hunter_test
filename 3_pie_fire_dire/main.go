package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

func beefSummaryHandler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	resp, err := client.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		http.Error(w, "Failed to fetch Bacon Ipsum", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	re := regexp.MustCompile(`[^\w\s]`)
	cleanText := re.ReplaceAllString(string(body), "")

	cleanText = strings.ToLower(cleanText)

	words := strings.Fields(cleanText)
	wordCount := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, word := range words {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			mu.Lock()
			wordCount[word]++
			mu.Unlock()
		}(word)
	}

	wg.Wait()
	response := map[string]interface{}{
		"beef": wordCount,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/beef/summary", beefSummaryHandler)

	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
